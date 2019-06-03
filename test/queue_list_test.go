package test

import (
	"fmt"
	"time"

	"github.com/brewlin/go-stl/queue/list"
)

func Queue_list(n int) {
	t1 := time.Now()
	q := list.NewQueue()
	for i := 0; i < n; i++ {
		q.Push(i)
	}
	fmt.Println("queue-list ", time.Now().Sub(t1))
}
