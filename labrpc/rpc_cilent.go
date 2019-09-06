package labrpc

import (
	"bytes"
	"encoding/gob"
	"log"
	"reflect"
)

type ClientEnd struct {
	endname interface{} // this end-point's name
	ch      chan reqMsg // copy of Network.endCh
}

// send an RPC, wait for the reply.
// the return value indicates success; false means the
// server couldn't be contacted.
func (e *ClientEnd) Call(svcMeth string, args interface{}, reply interface{}) bool {
	req := reqMsg{}
	req.endname = e.endname
	req.svcMeth = svcMeth
	req.argsType = reflect.TypeOf(args)
	req.replyCh = make(chan replyMsg)

	qb := new(bytes.Buffer)
	qe := gob.NewEncoder(qb)
	qe.Encode(args)
	req.args = qb.Bytes()

	e.ch <- req

	rep := <-req.replyCh
	if rep.ok {
		rb := bytes.NewBuffer(rep.reply)
		rd := gob.NewDecoder(rb)
		if err := rd.Decode(reply); err != nil {
			log.Fatalf("ClientEnd.Call(): decode reply: %v\n", err)
		}
		return true
	} else {
		return false
	}
}
