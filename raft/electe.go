package raft

import (
	"fmt"
	"log"
	"time"
)

// Raft protocol
type Raft struct {
	state         int
	voteFor       string
	me            string
	currentLeader int
	currentTerm   int
	peers         []int
}

func (rf *Raft) election() {
	var result bool
	for {
		//超时时间是raft协议推荐的 150-300 随机生成
		timeout := randomRange(150, 300)
		printTime()
		log.Printf("candidate=%d wait election timeout=%d\n", rf.me, timeout)
		rf.setMessageTime(milliseconds())
		for rf.lastMessageTime+timeout > milliseconds() {
			select {
			case <-time.After(time.Duration(timeout) * time.Millisecond):
				printTime()
				log.Printf("candidate=%d,lastMessageTime=%d,timeout=%d,plus=%d,now=%d\n", rf.me, rf.lastMessageTime, milliseconds)
				if rf.lastMessageTime+timeout <= milliseconds() {
					break
				} else {
					rf.setMessageTime(milliseconds())
					timeout = randomRange(150, 300)
					continue
				}
			}
		}
		printTime()
		fmt.Printf("candidate=%d timeouted\n", rf.me)
		//election till success
		result = false
		for !result {
			//真正开始选举
			result = rf.election_one_round()
		}

	}
}

//成为候选人 增加term，然后选举自己
func (rf *Raft) becomeCandidate() {
	rf.state = 1
	rf.setTerm(rf.currentTerm + 1)
	rf.voteFor = rf.me
	rf.currentLeader = -1
}

func (rf *Raft) election_one_round() {
	for i := 0; i < len(rf.peers); i++ {
		if i != rf.me {
			var args RequestVoteArgs
			server := i
			args.Term = rf.currentTerm
			args.CandidateId = rf.me
			var reply RequestVoteReply
			printTime()
			log.Printf("candidate=%d send request vote to server=%d\n", rf.me, i)
			go rf.sendRequestVoteAndTrigger(server, args, &reply, rpcTimeout)
		}
	}
	done = 0
	triggerHeartbeat = false
	for i:=0;i <len(rf.peers) - 1; i++ {
		printTime()
		log.Printf("candidate=%d waiting for select for i=%d\n",rf.me,i)
	}
}
func printTime() {

	log.Println(time())
}
