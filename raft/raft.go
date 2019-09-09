package raft

//
// this is an outline of the API that raft must expose to
// the service (or tester). see comments below for
// each of these functions for more details.
//
// rf = Make(...)
//   create a new Raft server.
// rf.Start(command interface{}) (index, term, isleader)
//   start agreement on a new log entry
// rf.GetState() (term, isLeader)
//   ask a Raft for its current term, and whether it thinks it is leader
// ApplyMsg
//   each time a new entry is committed to the log, each Raft peer
//   should send an ApplyMsg to the service (or tester)
//   in the same server.
//

import (
	"log"
	//"fmt"

	"math/rand"
	"sync"
	"time"

	labrpc "github.com/brewlin/go-stl/pkg/rpc"
)

const (
	LEADER = iota
	CANDIDATE
	FLLOWER

	HBINTERVAL = 50 * time.Millisecond // 50ms
)

//
// as each Raft peer becomes aware that successive log entries are
// committed, the peer should send an ApplyMsg to the service (or
// tester) on the same server, via the applyCh passed to Make().
//
type ApplyMsg struct {
	Index       int
	Command     interface{}
	UseSnapshot bool   // ignore for lab2; only used in lab3
	Snapshot    []byte // ignore for lab2; only used in lab3
}

//
// the service or tester wants to create a Raft server. the ports
// of all the Raft servers (including this one) are in peers[]. this
// server's port is peers[me]. all the servers' peers[] arrays
// have the same order. persister is a place for this server to
// save its persistent state, and also initially holds the most
// recent saved state, if any. applyCh is a channel on which the
// tester or service expects Raft to send ApplyMsg messages.
// Make() must return quickly, so it should start goroutines
// for any long-running work.
// 每个server节点都保存着相同的peers结构体，当前节点也保存在数组perrs[me] 中，persister表示持久化当前节点最近的历史状态
//
func Make(peers []*labrpc.ClientEnd, me int,
	persister *Persister, applyCh chan ApplyMsg) *Raft {
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

	// 事件循环
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

//
// A Go object implementing a single Raft peer.
//
type Raft struct {
	mu        sync.Mutex
	peers     []*labrpc.ClientEnd
	persister *Persister
	me        int // index into peers[]

	// Your data here.
	// Look at the paper's Figure 2 for a description of what
	// state a Raft server must maintain.

	//channel
	state         int
	voteCount     int
	chanCommit    chan bool
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
	//对于每一个服务器，记录需要发给它的下一个日志条目的索引（初始化为领导人上一条日志条目的索引值 + 1）
	nextIndex []int
	//对于每一个服务器，记录已经复制到该服务器的日志的最高索引值(从0递增)
	matchIndex []int
}

// return currentTerm and whether this server
// believes it is the leader.
func (rf *Raft) GetState() (int, bool) {
	return rf.currentTerm, rf.state == LEADER
}

// 获取最新的一个记录索引
func (rf *Raft) getLastIndex() int {
	return rf.log[len(rf.log)-1].LogIndex
}

// 查询最新的任期 term
func (rf *Raft) getLastTerm() int {
	return rf.log[len(rf.log)-1].LogTerm
}

//判断是否是leader节点
func (rf *Raft) IsLeader() bool {
	return rf.state == LEADER
}

//Start d
// the service using Raft (e.g. a k/v server) wants to start
// agreement on the next command to be appended to Raft's log. if this
// server isn't the leader, returns false. otherwise start the
// agreement and return immediately. there is no guarantee that this
// command will ever be committed to the Raft log, since the leader
// may fail or lose an election.
//
// the first return value is the index that the command will appear at
// if it's ever committed. the second return value is the current
// term. the third return value is true if this server believes it is
// the leader.
//
func (rf *Raft) Start(command interface{}) (int, int, bool) {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	index := -1
	term := rf.currentTerm
	isLeader := rf.state == LEADER
	if isLeader {
		index = rf.getLastIndex() + 1
		//fmt.Printf("raft:%d start\n",rf.me)
		rf.log = append(rf.log, LogEntry{LogTerm: term, LogCommand: command, LogIndex: index}) // append new entry from client
		rf.persist()
	}
	return index, term, isLeader
}

//Kill kill
// the tester calls Kill() when a Raft instance won't
// be needed again. you are not required to do anything
// in Kill(), but it might be convenient to (for example)
// turn off debug output from this instance.
//
func (rf *Raft) Kill() {
	// Your code here, if desired.
}
