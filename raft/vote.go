package raft

type RequestVoteArgs struct {
	//任期
	Term int
	//全局唯一候选者ID
	CandidateId  int
	LastLogTerm  int
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
