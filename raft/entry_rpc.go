package raft

//AppendEntries append
//leader节点进行日志同步  raft强制要求 leader节点 覆盖到其他server节点来保持一致性
func (rf *Raft) AppendEntries(args AppendEntriesArgs, reply *AppendEntriesReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	defer rf.persist()

	reply.Success = false
	//如果该leader任期 小于当前server节点的任期，则返回当前节点最新的索引
	if args.Term < rf.currentTerm {
		reply.Term = rf.currentTerm
		reply.NextIndex = rf.getLastIndex() + 1
		return
	}
	//收到leader 的心跳 重置定时器恢复正常
	rf.chanHeartbeat <- true

	//If RPC request or response contains term T > currentTerm: set currentTerm = T, convert to follower
	//更新当前任期号为最新，并重置身份
	if args.Term > rf.currentTerm {
		rf.currentTerm = args.Term
		rf.state = FLLOWER
		//重置投票权
		rf.votedFor = -1
	}
	reply.Term = args.Term

	if args.PrevLogIndex > rf.getLastIndex() {
		reply.NextIndex = rf.getLastIndex() + 1
		return
	}
	//获取第一个日志索引
	baseIndex := rf.log[0].LogIndex

	// If a follower’s log is inconsistent with the leader’s, the AppendEntries consis- tency check will fail in the next AppendEntries RPC.
	// Af- ter a rejection, the leader decrements nextIndex and retries the AppendEntries RPC
	//Eventually nextIndex will reach a point where the leader and follower logs match
	//which removes any conflicting entries in the follower’s log and appends entries from the leader’s log (if any).
	//检查日志是否同步，并检查出不同的节点索引值
	if args.PrevLogIndex > baseIndex {
		term := rf.log[args.PrevLogIndex-baseIndex].LogTerm
		if args.PrevLogTerm != term {
			//找出双方日志节点不一致的索引值 => 然后追随者FOLLOWER 会删除该索引之后的条目
			for i := args.PrevLogIndex - 1; i >= baseIndex; i-- {
				if rf.log[i-baseIndex].LogTerm != term {
					//返回当前节点与 该leader节点不一致的位置
					reply.NextIndex = i + 1
					break
				}
			}
			return
		}
	}
	if args.PrevLogIndex < baseIndex {

	} else {
		//追加该索引值后的不同日志条目，保持和领导人的一致性
		//Append any new entries not already in the log
		rf.log = rf.log[:args.PrevLogIndex+1-baseIndex]
		rf.log = append(rf.log, args.Entries...)
		//回复成功，这样就表示当前和leader节点的日志保持一致了
		reply.Success = true
		reply.NextIndex = rf.getLastIndex() + 1
	}
	//If leaderCommit > commitIndex, set commitIndex =min(leaderCommit, index of last new entry)
	if args.LeaderCommit > rf.commitIndex {
		last := rf.getLastIndex()
		if args.LeaderCommit > last {
			rf.commitIndex = last
		} else {
			rf.commitIndex = args.LeaderCommit
		}
		//进行提交
		rf.chanCommit <- true
	}
	return
}
