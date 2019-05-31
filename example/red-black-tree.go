package example

import (
	"fmt"
	"time"

	"github.com/brewlin/go-stl/tree/rb"
)

func Less(a, b interface{}) bool {
	return a.(int) < b.(int)
}
func More(a, b interface{}) bool {
	return a.(int) > b.(int)
}
func Equal(a, b interface{}) bool {
	return a.(int) == b.(int)
}
func RBTree(n int) {
	// arr := RandIntArr(n, 0, n)
	b := rb.NewRBTree(Less, More, Equal)
	t1 := time.Now()
	// for _, v := range arr {
	for i := 0; i < n; i++ {
		// b.Add(i)
		b.Add(i, i)
	}
	// for _, v := range arr {
	for i := 0; i < n; i++ {
		b.Remove(i)
	}
	// b.Prints()
	// fmt.Println(b.GetSize())
	fmt.Println("red-black tree", time.Now().Sub(t1))
	// b.PreOrder()
	// b.LevelOrder()
}
