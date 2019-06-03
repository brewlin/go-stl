package test

import (
	"fmt"
	"time"
)

func Go_queue(n int) {
	t1 := time.Now()
	list := make([]int, 1)
	for i := 0; i < n; i++ {
		list = append(list, i)
	}
	for i := 0; i < n; i++ {
		list = list[2:]
	}
	fmt.Println("go slice", time.Now().Sub(t1))

}
