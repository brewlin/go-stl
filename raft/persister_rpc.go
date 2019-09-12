package raft

import (
	"bytes"
	"encoding/gob"
)

type InstallSnapshotArgs struct {
	Term     int
	LeaderId int
	//快照中最后的日志条目对应的索引
	LastIncludedIndex int
	//快照中最后的日志条目对应的任期
	LastIncludedTerm int
	//快照原始数据
	Data []byte
}

type InstallSnapshotReply struct {
	//用于leader更新自己
	Term int
}

/**
 * server 节点根据快照更新自己的日志
 * 处理 leader节点发送的快照 rpc (可能当前节点的快照太落后了)
 */
func (rf *Raft) InstallSnapshot(args InstallSnapshotArgs, reply *InstallSnapshotReply) {
	rf.mu.Lock()
	defer rf.mu.Unlock()
	//当前任期比该leader节点大
	if args.Term < rf.currentTerm {
		reply.Term = rf.currentTerm
		return
	}
	//收到心跳,重置定时器
	rf.chanHeartbeat <- true
	rf.state = FLLOWER
	rf.currentTerm = rf.currentTerm
	//持久化保存leader发送的快照
	rf.persister.SaveSnapshot(args.Data)

	rf.log = truncateLog(args.LastIncludedIndex, args.LastIncludedTerm, rf.log)

	msg := ApplyMsg{UseSnapshot: true, Snapshot: args.Data}

	rf.lastApplied = args.LastIncludedIndex
	rf.commitIndex = args.LastIncludedIndex

	rf.persist()

	rf.chanApply <- msg
}

func (rf *Raft) sendInstallSnapshot(server int, args InstallSnapshotArgs, reply *InstallSnapshotReply) bool {
	ok := rf.peers[server].Call("Raft.InstallSnapshot", args, reply)
	if ok {
		if reply.Term > rf.currentTerm {
			rf.currentTerm = reply.Term
			rf.state = FLLOWER
			rf.votedFor = -1
			return ok
		}

		rf.nextIndex[server] = args.LastIncludedIndex + 1
		rf.matchIndex[server] = args.LastIncludedIndex
	}
	return ok
}
func (rf *Raft) readSnapshot(data []byte) {

	rf.readPersist(rf.persister.ReadRaftState())

	if len(data) == 0 {
		return
	}

	r := bytes.NewBuffer(data)
	d := gob.NewDecoder(r)

	var LastIncludedIndex int
	var LastIncludedTerm int

	d.Decode(&LastIncludedIndex)
	d.Decode(&LastIncludedTerm)

	rf.commitIndex = LastIncludedIndex
	rf.lastApplied = LastIncludedIndex

	rf.log = truncateLog(LastIncludedIndex, LastIncludedTerm, rf.log)

	msg := ApplyMsg{UseSnapshot: true, Snapshot: data}

	go func() {
		rf.chanApply <- msg
	}()
}

//
// restore previously persisted state.
//
func (rf *Raft) readPersist(data []byte) {
	// Your code here.
	// Example:
	r := bytes.NewBuffer(data)
	d := gob.NewDecoder(r)
	d.Decode(&rf.currentTerm)
	d.Decode(&rf.votedFor)
	d.Decode(&rf.log)
}

func (rf *Raft) GetPerisistSize() int {
	return rf.persister.RaftStateSize()
}

func (rf *Raft) StartSnapshot(snapshot []byte, index int) {

	rf.mu.Lock()
	defer rf.mu.Unlock()

	baseIndex := rf.log[0].LogIndex
	lastIndex := rf.getLastIndex()

	if index <= baseIndex || index > lastIndex {
		// in case having installed a snapshot from leader before snapshotting
		// second condition is a hack
		return
	}

	var newLogEntries []LogEntry

	newLogEntries = append(newLogEntries, LogEntry{LogIndex: index, LogTerm: rf.log[index-baseIndex].LogTerm})

	for i := index + 1; i <= lastIndex; i++ {
		newLogEntries = append(newLogEntries, rf.log[i-baseIndex])
	}

	rf.log = newLogEntries

	rf.persist()

	w := new(bytes.Buffer)
	e := gob.NewEncoder(w)
	e.Encode(newLogEntries[0].LogIndex)
	e.Encode(newLogEntries[0].LogTerm)

	data := w.Bytes()
	data = append(data, snapshot...)
	rf.persister.SaveSnapshot(data)
}
