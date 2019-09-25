package mr

import "fmt"

func (mr *Master) schedule(phase jobPhase) {
	var ntasks int
	var nios int
	switch phase {
	case mapPhase:
		ntasks = len(mr.files)
		nios = mr.nReduce

	case reducePhase:
		ntasks = mr.nReduce
		nios = len(mr.files)
	}
	fmt.Println("schedule ", ntasks, phase, nios)

	done := make(chan bool)
	for i := 0; i < ntasks; i++ {
		go func(number int) {
			args := DoTaskArgs{mr.jobName, mr.files[number], phase, number, nios}
			var worker string
			reply := new(struct{})
			ok := false
			for ok != true {
				worker = <-mr.registerC
				ok = call(worker, "Worker.DoTask", args, reply)
			}
			done <- true
			mr.registerC <- worker
		}(i)
	}
	for i := 0; i < ntasks; i++ {
		//阻塞等待所有worker完成任务
		<-done
	}
	fmt.Println("schedule done", phase)

}
