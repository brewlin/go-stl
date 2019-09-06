package labrpc

//
// channel-based RPC, for 824 labs.
// allows tests to disconnect RPC connections.
//
// we will use the original labrpc.go to test your code for grading.
// so, while you can modify this code to help you debug, please
// test against the original before submitting.
//
// adapted from Go net/rpc/server.go.
//
// sends gob-encoded values to ensure that RPCs
// don't include references to program objects.
//
// net := MakeNetwork() -- holds network, clients, servers.
// end := net.MakeEnd(endname) -- create a client end-point, to talk to one server.
// net.AddServer(servername, server) -- adds a named server to network.
// net.DeleteServer(servername) -- eliminate the named server.
// net.Connect(endname, servername) -- connect a client to a server.
// net.Enable(endname, enabled) -- enable/disable a client.
// net.Reliable(bool) -- false means drop/delay messages
//
// end.Call("Raft.AppendEntries", &args, &reply) -- send an RPC, wait for reply.
// the "Raft" is the name of the server struct to be called.
// the "AppendEntries" is the name of the method to be called.
// Call() returns true to indicate that the server executed the request
// and the reply is valid.
// Call() returns false if the network lost the request or reply
// or the server is down.
// It is OK to have multiple Call()s in progress at the same time on the
// same ClientEnd.
// Concurrent calls to Call() may be delivered to the server out of order,
// since the network may re-order messages.
// Call() is guaranteed to return (perhaps after a delay) *except* if the
// handler function on the server side does not return. That is, there
// is no need to implement your own timeouts around Call().
// the server RPC handler function must declare its args and reply arguments
// as pointers, so that their types exactly match the types of the arguments
// to Call().
//
// srv := MakeServer()
// srv.AddService(svc) -- a server can have multiple services, e.g. Raft and k/v
//   pass srv to net.AddServer()
//
// svc := MakeService(receiverObject) -- obj's methods will handle RPCs
//   much like Go's rpcs.Register()
//   pass svc to srv.AddService()
//

import "reflect"
import "sync"
import "log"

import "math/rand"
import "time"

type reqMsg struct {
	endname  interface{} // name of sending ClientEnd
	svcMeth  string      // e.g. "Raft.AppendEntries"
	argsType reflect.Type
	args     []byte
	replyCh  chan replyMsg
}

type replyMsg struct {
	ok    bool
	reply []byte
}

type Network struct {
	mu             sync.Mutex
	reliable       bool
	longDelays     bool                        // pause a long time on send on disabled connection
	longReordering bool                        // sometimes delay replies a long time
	ends           map[interface{}]*ClientEnd  // ends, by name
	enabled        map[interface{}]bool        // by end name
	servers        map[interface{}]*Server     // servers, by name
	connections    map[interface{}]interface{} // endname -> servername
	endCh          chan reqMsg
}

func MakeNetwork() *Network {
	rn := &Network{}
	rn.reliable = true
	rn.ends = map[interface{}]*ClientEnd{}
	rn.enabled = map[interface{}]bool{}
	rn.servers = map[interface{}]*Server{}
	rn.connections = map[interface{}](interface{}){}
	rn.endCh = make(chan reqMsg)

	// single goroutine to handle all ClientEnd.Call()s
	//endch 监听rpc 请求 并分发进行rpc请求
	go func() {
		//处理rpc请求事件,遍历通道
		for xreq := range rn.endCh {
			go rn.ProcessReq(xreq)
		}
	}()

	return rn
}

func (rn *Network) Reliable(yes bool) {
	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.reliable = yes
}

func (rn *Network) LongReordering(yes bool) {
	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.longReordering = yes
}

func (rn *Network) LongDelays(yes bool) {
	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.longDelays = yes
}

//ReadEndnameInfo 获取rpc客户端基本信息
func (rn *Network) ReadEndnameInfo(endname interface{}) (enabled bool,
	servername interface{}, server *Server, reliable bool, longreordering bool,
) {
	rn.mu.Lock()
	defer rn.mu.Unlock()
	//rpc客户端是否关闭
	enabled = rn.enabled[endname]
	//获取servername
	servername = rn.connections[endname]
	if servername != nil {
		server = rn.servers[servername]
	}
	//可用
	reliable = rn.reliable
	longreordering = rn.longReordering
	return
}

//IsServerDead is
func (rn *Network) IsServerDead(endname interface{}, servername interface{}, server *Server) bool {
	rn.mu.Lock()
	defer rn.mu.Unlock()

	if rn.enabled[endname] == false || rn.servers[servername] != server {
		return true
	}
	return false
}

//ProcessReq rpc请求执行
func (rn *Network) ProcessReq(req reqMsg) {
	enabled, servername, server, reliable, longreordering := rn.ReadEndnameInfo(req.endname)

	//检查rpc client 是否合法
	if enabled && servername != nil && server != nil {
		//是否可靠
		if reliable == false {
			// 随机延时后再进行提交
			ms := (rand.Int() % 27)
			time.Sleep(time.Duration(ms) * time.Millisecond)
		}
		//丢掉请求，返回像超时那样
		if reliable == false && (rand.Int()%1000) < 100 {
			// drop the request, return as if timeout
			req.replyCh <- replyMsg{false, nil}
			return
		}

		// execute the request (call the RPC handler).
		// in a separate thread so that we can periodically check
		// if the server has been killed and the RPC should get a
		// failure reply.
		ech := make(chan replyMsg)

		go func() {
			r := server.dispatch(req)
			ech <- r
		}()

		// wait for handler to return,
		// but stop waiting if DeleteServer() has been called,
		// and return an error.
		var reply replyMsg
		replyOK := false
		serverDead := false
		for replyOK == false && serverDead == false {
			select {
			//监听rpc 请求的响应
			case reply = <-ech:
				replyOK = true
			//如果超时
			case <-time.After(100 * time.Millisecond):
				serverDead = rn.IsServerDead(req.endname, servername, server)
			}
		}

		// do not reply if DeleteServer() has been called, i.e.
		// the server has been killed. this is needed to avoid
		// situation in which a client gets a positive reply
		// to an Append, but the server persisted the update
		// into the old Persister. config.go is careful to call
		// DeleteServer() before superseding the Persister.
		serverDead = rn.IsServerDead(req.endname, servername, server)
		//返回server端被kill调了
		if replyOK == false || serverDead == true {
			// server was killed while we were waiting; return error.
			req.replyCh <- replyMsg{false, nil}
			//返回 超时的错误
		} else if reliable == false && (rand.Int()%1000) < 100 {
			// drop the reply, return as if timeout
			req.replyCh <- replyMsg{false, nil}
			//延时返回这条消息
		} else if longreordering == true && rand.Intn(900) < 600 {
			// delay the response for a while
			ms := 200 + rand.Intn(1+rand.Intn(2000))
			time.Sleep(time.Duration(ms) * time.Millisecond)
			req.replyCh <- reply
		} else { //正常情况下 返回rpc response
			req.replyCh <- reply
		}
	} else {
		// simulate no reply and eventual timeout.
		ms := 0
		if rn.longDelays {
			// let Raft tests check that leader doesn't send
			// RPCs synchronously.
			ms = (rand.Int() % 7000)
		} else {
			// many kv tests require the client to try each
			// server in fairly rapid succession.
			ms = (rand.Int() % 100)
		}
		time.Sleep(time.Duration(ms) * time.Millisecond)
		req.replyCh <- replyMsg{false, nil}
	}

}

// create a client end-point.
// start the thread that listens and delivers.
// 创建一个客户端，并启动线程 去监听和分发
func (rn *Network) MakeEnd(endname interface{}) *ClientEnd {
	rn.mu.Lock()
	defer rn.mu.Unlock()

	if _, ok := rn.ends[endname]; ok {
		log.Fatalf("MakeEnd: %v already exists\n", endname)
	}

	e := &ClientEnd{}
	e.endname = endname
	e.ch = rn.endCh
	rn.ends[endname] = e
	rn.enabled[endname] = false
	rn.connections[endname] = nil

	return e
}

func (rn *Network) AddServer(servername interface{}, rs *Server) {
	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.servers[servername] = rs
}

func (rn *Network) DeleteServer(servername interface{}) {
	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.servers[servername] = nil
}

// connect a ClientEnd to a server.
// a ClientEnd can only be connected once in its lifetime.
// 连接远程server 并在该生命周期中只会连接一次
func (rn *Network) Connect(endname interface{}, servername interface{}) {
	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.connections[endname] = servername
}

// enable/disable a ClientEnd.
// 激活 |注销 一个rpc客户端
func (rn *Network) Enable(endname interface{}, enabled bool) {
	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.enabled[endname] = enabled
}

// get a server's count of incoming RPCs.
func (rn *Network) GetCount(servername interface{}) int {
	rn.mu.Lock()
	defer rn.mu.Unlock()

	svr := rn.servers[servername]
	return svr.GetCount()
}
