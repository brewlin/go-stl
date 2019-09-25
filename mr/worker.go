package mr

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"sync"
)

//Worker 提供两个rpc接口，dotask，shutdown
type Worker struct {
	sync.Mutex

	name   string
	Map    func(string, string) []KeyValue
	Reduce func(string, []string) string
	nRPC   int
	nTasks int
	l      net.Listener
}

//DoTask 提供接口 由master调度执行worker任务
func (wk *Worker) DoTask(arg *DoTaskArgs, _ *struct{}) error {
	fmt.Println("given  task  on file (nios:)\n", wk.name, arg.Phase, arg.TaskNumber, arg.File, arg.NumOtherPhase)
	switch arg.Phase {
	case mapPhase:
		doMap(arg.JobName, arg.TaskNumber, arg.File, arg.NumOtherPhase, wk.Map)
	case reducePhase:
		doReduce(arg.JobName, arg.TaskNumber, arg.NumOtherPhase, wk.Reduce)
	}
	fmt.Printf("%s: %v task #%d done\n", wk.name, arg.Phase, arg.TaskNumber)
	return nil

}

// rpc通知master当前worker已准备好工作
func (wk *Worker) register(master string) {
	args := new(RegisterArgs)
	args.Worker = wk.name
	ok := call(master, "Master.Register", args, new(struct{}))
	if ok == false {
		fmt.Printf("Register :rpc %s register error\n", master)
	}
}

//Shutdown 提供关闭rpc 由master调用
func (wk *Worker) Shutdown(_ *struct{}, re *ShutdownReply) error {
	debug("shutdown %s\n", wk.name)
	wk.Lock()
	defer wk.Unlock()
	re.Ntasks = wk.nTasks
	wk.nRPC = 1
	wk.nTasks--
	return nil
}

//RunWorker run the worker
func RunWorker(master string, me string, mapF func(string, string) []KeyValue, reduceF func(string, []string) string, nRPC int) {
	debug("run worker %s\n", me)
	wk := new(Worker)
	wk.name = me
	wk.Map = mapF
	wk.Reduce = reduceF
	wk.nRPC = nRPC
	rpcs := rpc.NewServer()
	rpcs.Register(wk)
	os.Remove(me)
	l, e := net.Listen("unix", me)
	if e != nil {
		log.Fatal("run worker:worker", me, " error:", e)
	}
	wk.l = l
	wk.register(master)

	for {
		wk.Lock()
		if wk.nRPC == 0 {
			wk.Unlock()
			break
		}
		wk.Unlock()
		conn, err := wk.l.Accept()
		if err == nil {
			wk.Lock()
			wk.nRPC--
			wk.Unlock()
			go rpcs.ServeConn(conn)
			wk.Lock()
			wk.nTasks++
			wk.Unlock()
		} else {
			break
		}
	}
	wk.l.Close()
	debug("run worker %s exist\n", me)

}
