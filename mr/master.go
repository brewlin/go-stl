package mr

import (
	"fmt"
	"net"
	"sync"
)

//Master master
type Master struct {
	sync.Mutex

	address   string
	registerC chan string
	doneC     chan bool
	workers   []string

	//当前执行的job
	jobName string
	//输入文件列表
	files []string
	//reduce 分区数
	nReduce int

	shutdownC chan struct{}
	//net 监听
	l     net.Listener
	stats []int
}

//Register register
func (mr *Master) Register(args *RegisterArgs, _ *struct{}) error {
	mr.Lock()
	defer mr.Unlock()

	debug("Register : worker %s\n", args.Worker)
	mr.workers = append(mr.workers, args.Worker)
	go func() {
		mr.registerC <- args.Worker
	}()
	return nil
}

func newMaster(master string) (mr *Master) {
	mr = new(Master)
	mr.address = master
	mr.shutdownC = make(chan struct{})
	mr.registerC = make(chan string)
	mr.doneC = make(chan bool)
	return
}

//Sequential 顺序执行map和reduce任务
func Sequential(jobName string, files []string, nreduce int, mapF func(string, string) []KeyValue, reduceF func(string, []string) string) (mr *Master) {
	mr = newMaster("master")
	go mr.run(jobName, files, nreduce, func(phase jobPhase) {
		switch phase {
		case mapPhase:
			for i, f := range mr.files {
				doMap(mr.jobName, i, f, mr.nReduce, mapF)
			}
		case reducePhase:
			for i := 0; i < mr.nReduce; i++ {
				doReduce(mr.jobName, i, len(mr.files), reduceF)
			}

		}

	}, func() {
		mr.stats = []int{len(files) + nreduce}
	})
	return
}

//Distributed 分布式调度 通讯基于rpc
func Distributed(jobName string, files []string, nreduce int, master string) (mr *Master) {
	mr = newMaster(master)
	mr.startRPCServer()
	go mr.run(jobName, files, nreduce, mr.schedule, func() {
		mr.stats = mr.killWorkers()
		mr.stopRPCServer()
	})
	return
}
func (mr *Master) run(jobName string, files []string, nreduce int, schedule func(phase jobPhase), finish func()) {
	mr.jobName = jobName
	mr.files = files
	mr.nReduce = nreduce
	fmt.Printf("%s:Starting Map/Reduce task %s\n", mr.address, mr.jobName)
	schedule(mapPhase)
	schedule(reducePhase)
	finish()
	mr.merge()
	fmt.Printf("%s: Map/Reduce task completed\n", mr.address)
	mr.doneC <- true
}

//给所有worker发送shutdownrpc关闭worker
func (mr *Master) killWorkers() []int {
	mr.Lock()
	defer mr.Unlock()
	ntasks := make([]int, 0, len(mr.workers))
	for _, w := range mr.workers {
		debug("Master : shutdown worker %s\n", w)
		var reply ShutdownReply
		ok := call(w, "Worker.Shutdown", new(struct{}), &reply)
		if ok == false {
			fmt.Println("master : rpc shutdown error")
		} else {
			ntasks = append(ntasks, reply.Ntasks)
		}
	}
	return ntasks

}

//Wait wait
func (mr *Master) Wait() {
	<-mr.doneC
}
