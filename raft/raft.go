package raft

import (
	"sync"
	"time"
	"log"
)

const (
	//LEADER 主
	LEADER = iota
	CANDIDATE
	FLLOWER

	HBINTERVAL = 50 * time.Millisecond // 50ms
)

func newRaft(peers []*rpc.ClientEnd,me int,persister *Persister,applyCh chan ApplyMsg)*Raft{
	rf := &Raft{}
	rf.peers = peers
	rf.persister = persister
	rf.me = me

	// Your initialization code here.
	rf.state = FLLOWER
	rf.votedFor = -1
	rf.log = append(rf.log, LogEntry{LogTerm: 0})
	rf.currentTerm = 0
	rf.chanCommit = make(chan bool, 100)
	rf.chanHeartbeat = make(chan bool, 100)
	rf.chanGrantVote = make(chan bool, 100)
	rf.chanLeader = make(chan bool, 100)
	rf.chanApply = applyCh

	// initialize from state persisted before a crash
	rf.readPersist(persister.ReadRaftState())
	rf.readSnapshot(persister.ReadSnapshot())

	//投票选举
	go func() {
		//进行事件循环
		for {
			switch rf.state {
			//普通server节点
			case FLLOWER:
				select {
				case <-rf.chanHeartbeat:
				case <-rf.chanGrantVote:
				//随机超时时间内未收到leader的心跳包，到期置为候选者身份参与leader选举
				case <-time.After(time.Duration(rand.Int63()%333+550) * time.Millisecond):
					rf.state = CANDIDATE
				}
			//leader节点 进行rpc广播给所有server节点
			case LEADER:
				log.Printf("Leader:%v %v\n", rf.me, "boatcastAppendEntries	")
				rf.broadcastAppendEntries()
				time.Sleep(HBINTERVAL)
			//候选者节点
			case CANDIDATE:
				//保证原子操作
				rf.mu.Lock()
				//To begin an election, a follower increments its current term and transitions to candidate state
				//开始选举，增加当前任期
				rf.currentTerm++
				//It then votes for itself and issues RequestVote RPCs in parallel to each of the other servers in the cluster.
				//投票给自己
				rf.votedFor = rf.me
				//投票+1 自己投的票
				rf.voteCount = 1
				//持久化当前raft状态
				rf.persist()
				rf.mu.Unlock()
				// 进行投票选举leader
				// 1.当前赢得leader选举
				// 2.当前选举失败，其他节点选举为leader节点
				// 3.该选举期间没有任何节点选举成功
				go rf.broadcastRequestVote()
				select {
				//增加超时时间
				case <-time.After(time.Duration(rand.Int63()%300+510) * time.Millisecond):
				//收到心跳 表明其他的 server选举为leader节点
				case <-rf.chanHeartbeat:
					rf.state = FLLOWER
					log.Printf("CANDIDATE %v reveive chanHeartbeat\n", rf.me)
				//成功选举为leader 节点
				case <-rf.chanLeader:
					rf.mu.Lock()
					//切换为leader身份
					rf.state = LEADER
					log.Printf("%v is Leader\n", rf.me)
					rf.nextIndex = make([]int, len(rf.peers))
					rf.matchIndex = make([]int, len(rf.peers))
					for i := range rf.peers {
						//The leader maintains a nextIndex for each follower, which is the index of the next log entry the leader will send to that follower.
						// When a leader first comes to power, it initializes all nextIndex values to the index just after the last one in its log
						rf.nextIndex[i] = rf.getLastIndex() + 1
						rf.matchIndex[i] = 0
					}
					rf.mu.Unlock()
					//rf.boatcastAppendEntries()
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case <-rf.chanCommit:
				rf.mu.Lock()
				commitIndex := rf.commitIndex
				baseIndex := rf.log[0].LogIndex
				for i := rf.lastApplied + 1; i <= commitIndex; i++ {
					msg := ApplyMsg{Index: i, Command: rf.log[i-baseIndex].LogCommand}
					applyCh <- msg
					//fmt.Printf("me:%d %v\n",rf.me,msg)
					rf.lastApplied = i
				}
				rf.mu.Unlock()
			}
		}
	}()
	return rf
}

//Raft raft
type Raft struct {
	mu sync.Mutex
	peers []*rpc.ClientEnd
	persister *Persister
	me int
	//当前身份
	state int 

	voteCount int

	chanCommit chan bool
	chanHeartbeat chan bool
	chanGrantVote chan bool
	chanLeader    chan bool
	chanApply     chan ApplyMsg

	//persistent state on all server
	currentTerm int
	votedFor    int
	log         []LogEntry

	//volatile state on all servers
	commitIndex int
	lastApplied int

	//volatile state on leader
	nextIndex  []int
	matchIndex []int
}