package mr

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
)

//Shutdown 关闭rpc服务
func (mr *Master) Shutdown(_, _ *struct{}) error {
	debug("shutdown : server")
	close(mr.shutdownC)
	mr.l.Close()
	return nil
}

//启动rpcserver 提供rpc服务调用
func (mr *Master) startRPCServer() {
	rpcs := rpc.NewServer()
	//注册相关rpc 接口
	rpcs.Register(mr)
	os.Remove(mr.address)
	//基于unix域套接字通讯，类似本地进程IPC
	l, e := net.Listen("unix", mr.address)
	if e != nil {
		log.Fatal("regiter rpc server", mr.address, " error :", e)
	}
	mr.l = l

	go func() {
	loop:
		for {
			select {
			case <-mr.shutdownC:
				break loop
			default:
			}
			conn, err := mr.l.Accept()
			if err == nil {
				//处理rpc请求
				go func() {
					rpcs.ServeConn(conn)
					conn.Close()
				}()
			} else {
				debug("register rpc accept error %s", err)
				break
			}
		}
		debug("register server done")
	}()
}

func (mr *Master) stopRPCServer() {
	var reply ShutdownReply
	ok := call(mr.address, "Master.Shutdown", new(struct{}), &reply)
	if ok == false {
		fmt.Printf("clean up rpc %s error \n", mr.address)
	}
	debug("clean done\n")

}
