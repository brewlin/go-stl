package example

import (
	"fmt"
	"stl/queue/list"
	"time"
)

func List() {
	num := 10000000
	queue := list.NewQueue()
	arr := RandIntArr(num, 1, num)

	t1 := time.Now()
	for i := 0; i < num; i++ {
		queue.Push(arr[i])
	}
	for i := 0; i < num; i++ {
		queue.Pop()
	}
	fmt.Println(time.Now().Sub(t1))
}
