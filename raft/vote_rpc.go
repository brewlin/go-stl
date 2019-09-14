package raft

//
// example RequestVote RPC handler.
//
func (rf *Raft) RequestVote(args RequestVoteArgs, reply *RequestVoteReply) {
	// Your code here.
	rf.mu.Lock()
	defer rf.mu.Unlock()
	defer rf.persist()
	reply.VoteGranted = false
	//????????????
	if args.Term < rf.currentTerm {
		reply.Term = rf.currentTerm
		return
	}
	//If RPC request or response contains term T > currentTerm: set currentTerm = T, convert to follower
	//??????? ??????
	if args.Term > rf.currentTerm {
		rf.currentTerm = args.Term
		rf.state = FLLOWER
		//?????
		rf.votedFor = -1
	}
	reply.Term = rf.currentTerm
	//????????
	term := rf.getLastTerm()
	//????????
	index := rf.getLastIndex()
	uptoDate := false

	//If votedFor is null or candidateId, and candidate’s log is at least as up-to-date as receiver’s log, grant vote
	//Raft determines which of two logs is more up-to-date by comparing the index and term of the last entries in the logs.
	// If the logs have last entries with different terms,then the log with the later term is more up-to-date.
	// If the logs end with the same term, then whichever log is longer is more up-to-date
	//???????????????????
	if args.LastLogTerm > term {
		uptoDate = true
	}
	//?????????????????????????????
	if args.LastLogTerm == term && args.LastLogIndex >= index {
		// at least up to date
		uptoDate = true
	}
	//?????????????? ?????
	if (rf.votedFor == -1 || rf.votedFor == args.CandidateId) && uptoDate {
		//rpc?? true
		rf.chanGrantVote <- true
		rf.state = FLLOWER
		reply.VoteGranted = true
		//?????? id
		rf.votedFor = args.CandidateId
	}
}
