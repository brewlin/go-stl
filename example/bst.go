package example

import (
	"fmt"
	"stl/tree/bst"
	"time"
)

func Bst() {
	n := 1000000
	arr := RandIntArr(n, 10, n)
	b := bst.NewBst()
	t1 := time.Now()
	for _, v := range arr {
		b.Add(v)
	}
	for _, v := range arr {
		b.Remove(v)
	}
	fmt.Println(b)
	fmt.Println(time.Now().Sub(t1))
	// b.PreOrder()
	// b.LevelOrder()
}
