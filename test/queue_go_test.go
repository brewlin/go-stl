package test

import (
	"fmt"
	"time"
	"testing"
)

func BenchmarkQueuego(b *testing.B) {
	t1 := time.Now()
	list := make([]int, 1)
	for i := 0; i < b.N; i++ {
		list = append(list, i)
	}
	for i := 0; i < b.N; i++ {
		list = list[2:]
	}
	fmt.Println("go slice", time.Now().Sub(t1))

}
