package test

import (
	"fmt"
	"time"
	"testing"

	"github.com/brewlin/go-stl/queue/list"
)

func BenchmarkQueuelist(b *testing.B) {
	t1 := time.Now()
	q := list.NewQueue()
	for i := 0; i < b.N; i++ {
		q.Push(i)
	}
	fmt.Println("queue-list ", time.Now().Sub(t1))
}
