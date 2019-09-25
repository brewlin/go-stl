package mr

import (
	"encoding/json"
	"hash/fnv"
	"io/ioutil"
	"log"
	"os"
)

func doMap(jobName string, mapTaskNumber int, inFile string, nReduce int, mapF func(file string, contents string) []KeyValue) {
	//step1 读取文件
	contents, err := ioutil.ReadFile(inFile)
	if err != nil {
		log.Fatal("map read file error", err)
	}
	//step2 调用用户的map回调函数获取keyvalue值
	kvs := mapF(inFile, string(contents))
	//step3 写入临时文件
	var tmpFiles = make([]*os.File, nReduce)
	var encoders = make([]*json.Encoder, nReduce)

	for i := 0; i < nReduce; i++ {
		tmpFileName := reduceName(jobName, mapTaskNumber, i)
		tmpFiles[i], err = os.Create(tmpFileName)
		if err != nil {
			log.Fatal(err)
		}
		defer tmpFiles[i].Close()
		encoders[i] = json.NewEncoder(tmpFiles[i])
	}

	for _, kv := range kvs {
		hashKey := int(ihash(kv.Key)) % nReduce
		err := encoders[hashKey].Encode(&kv)
		if err != nil {
			log.Fatal("do map encoders", err)
		}
	}

}

func ihash(s string) uint32 {
	h := fnv.New32a()
	h.Write(([]byte(s)))
	return h.Sum32()
}
