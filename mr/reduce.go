package mr

import (
	"encoding/json"
	"log"
	"os"
	"sort"
)

func doReduce(jobName string, reduceTaskNumber int, nMap int, reduceF func(key string, values []string) string) {
	//step1 迭代合并map
	kvs := make(map[string][]string)
	for i := 0; i < nMap; i++ {
		fileName := reduceName(jobName, i, reduceTaskNumber)
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal("doReduce1: ", err)
		}
		dec := json.NewDecoder(file)
		for {
			var kv KeyValue
			err = dec.Decode(&kv)
			if err != nil {
				break
			}
			_, ok := kvs[kv.Key]
			if !ok {
				kvs[kv.Key] = []string{}
				kvs[kv.Key] = append(kvs[kv.Key], kv.Value)
			}
			file.Close()
		}
	}
	var keys []string
	for k := range kvs {
		keys = append(keys, k)
	}

	//step2 根据keys 排序
	sort.Strings(keys)
	//step3 创建结果写入file
	p := mergeName(jobName, reduceTaskNumber)
	file, err := os.Create(p)
	if err != nil {
		log.Fatal("doReduce2: create ", err)
	}
	enc := json.NewEncoder(file)
	//step4 调用用户回调函数
	for _, k := range keys {
		res := reduceF(k, kvs[k])
		enc.Encode(KeyValue{k, res})
	}
	file.Close()
}
