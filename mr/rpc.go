package mr

import (
	"fmt"
	"net/rpc"
)

//DoTaskArgs dotask
type DoTaskArgs struct {
	JobName    string
	File       string
	Phase      jobPhase
	TaskNumber int

	NumOtherPhase int
}

//ShutdownReply shutdownrelpy
type ShutdownReply struct {
	Ntasks int
}

//RegisterArgs register
type RegisterArgs struct {
	Worker string
}

func call(srv string, rpcname string, args interface{}, reply interface{}) bool {
	c, errx := rpc.Dial("unix", srv)
	if errx != nil {
		return false
	}
	defer c.Close()

	err := c.Call(rpcname, args, reply)
	if err == nil {
		return true
	}
	fmt.Println(err)
	return false

}
