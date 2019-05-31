package example

import (
	"fmt"
	"time"

	"github.com/brewlin/go-stl/tree/bs"
)

func BSTree(n int) {
	// arr := RandIntArr(n, 0, n)
	b := bs.NewBSTree()
	t1 := time.Now()
	// for _, v := range arr {
	for i := 0; i < n; i++ {
		b.Add(i)
		// b.Add(v)
	}
	// }
	// for _, v := range arr {
	for i := 0; i < n; i++ {
		b.Remove(i)
	}
	fmt.Println("binary-serach-tree", time.Now().Sub(t1))
	// b.PreOrder()
	// b.LevelOrder()
}
