package mr

import (
	"fmt"
	"strconv"
)

const debugEnable = true

func debug(format string, a ...interface{}) (n int, err error) {
	if debugEnable {
		n, err = fmt.Printf(format, a...)
	}
	return
}

type jobPhase string

const (
	mapPhase    jobPhase = "Map"
	reducePhase          = "Reduce"
)

//KeyValue key=>value
type KeyValue struct {
	Key   string
	Value string
}

func reduceName(jobName string, mapTask int, reduceTask int) string {
	return "mrtmp." + jobName + "-" + strconv.Itoa(mapTask) + "-" + strconv.Itoa(reduceTask)
}

func mergeName(jobName string, reduceTask int) string {
	return "mrtmp." + jobName + "-res-" + strconv.Itoa(reduceTask)
}
