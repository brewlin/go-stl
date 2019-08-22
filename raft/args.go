package raft

type RequestVoteArgs struct {
	Term        int
	CandidateId string
}

type RequestVoteReply struct {
}
