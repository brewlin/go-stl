package raft

// AppendEntriesArgs rpc心跳包，leader心跳广播所有节点
type AppendEntriesArgs struct {
	//该term为leader的term
	Term int
	//全局leader唯一id
	LeaderId     int
	PrevLogTerm  int
	PrevLogIndex int
	Entries      []LogEntry
	LeaderCommit int
}

//AppendEntriesReply 心跳rpc的回复信息
type AppendEntriesReply struct {
	//收到所有server节点的rpc回复term
	Term int
	//是否认可当前leader身份 继续为主节点
	Success   bool
	NextIndex int
}
