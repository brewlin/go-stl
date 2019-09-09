package raft

//
// example RequestVote RPC arguments structure.
//
type RequestVoteArgs struct {
	//任期
	Term int
	//全局唯一候选者ID
	CandidateId int
	//候选人最新日志条目对应的任期号
	LastLogTerm int
	//候选人最新日志条目的索引值
	LastLogIndex int
}

//
// example RequestVote RPC reply structure.
//
type RequestVoteReply struct {
	//表示确认的server的任期 如果candidate的term小于它。则更新
	Term int
	//表明该follower是否给自己投票
	VoteGranted bool
}

//
// example code to send a RequestVote RPC to a server.
// server is the index of the target server in rf.peers[].
// expects RPC arguments in args.
// fills in *reply with RPC reply, so caller should
// pass &reply.
// the types of the args and reply passed to Call() must be
// the same as the types of the arguments declared in the
// handler function (including whether they are pointers).
//
// returns true if labrpc says the RPC was delivered.
//
// if you're having trouble getting RPC to work, check that you've
// capitalized all field names in structs passed over RPC, and
// that the caller passes the address of the reply struct with &, not
// the struct itself.
// 给server 节点发送投票参与选举leader
//
func (rf *Raft) sendRequestVote(server int, args RequestVoteArgs, reply *RequestVoteReply) bool {
	//调用投票rpc 像 server 节点发送rpc请求
	ok := rf.peers[server].Call("Raft.RequestVote", args, reply)
	rf.mu.Lock()
	defer rf.mu.Unlock()
	if ok {
		term := rf.currentTerm
		// 如果当前以不是候选者身份，则退出选举为主
		if rf.state != CANDIDATE {
			return ok
		}
		// 当前任期发生了改变(说明在此期间 有其他的投票返回发现任期比自己新，则应该结束当前投票)
		if args.Term != term {
			return ok
		}
		// 如果回复的任期term 比自己还大，说明失去该投票
		if reply.Term > term {
			//更新自己的任期为最新的对方的任期
			rf.currentTerm = reply.Term
			//更新自己的状态为fllower
			rf.state = FLLOWER
			rf.votedFor = -1
			//触发持久化
			rf.persist()
		}
		//说明对方投了自己一票
		if reply.VoteGranted {
			//加上一票
			rf.voteCount++
			//如果当前为候选者 且 投票数 大于 一半的节点
			if rf.state == CANDIDATE && rf.voteCount > len(rf.peers)/2 {
				rf.state = FLLOWER
				//切换为
				rf.chanLeader <- true
			}
		}
	}
	return ok
}

/**
 * candidate 候选者进行广播投票选举
 */
func (rf *Raft) broadcastRequestVote() {
	var args RequestVoteArgs
	rf.mu.Lock()
	args.Term = rf.currentTerm
	args.CandidateId = rf.me
	args.LastLogTerm = rf.getLastTerm()
	args.LastLogIndex = rf.getLastIndex()
	rf.mu.Unlock()

	for i := range rf.peers {
		if i != rf.me && rf.state == CANDIDATE {
			//并行的推送给所有节点进行选举leader
			go func(i int) {
				var reply RequestVoteReply
				rf.sendRequestVote(i, args, &reply)
			}(i)
		}
	}
}
