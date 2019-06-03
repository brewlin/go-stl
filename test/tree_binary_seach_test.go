package test

import (
	"testing"

	"github.com/brewlin/go-stl/tree/bs"
)

func BenchmarkBSTree(bt *testing.B) {
	arr := RandIntArr(bt.N, 0, bt.N)
	bt.ReportAllocs()
	bt.ResetTimer()
	b := bs.NewBSTree()
	for _, v := range arr {
		// for i := 0; i < bt.N; i++ {
		// b.Add(i)
		b.Add(v)
	}
	// }
	for _, v := range arr {
		// for i := 0; i < bt.N; i++ {
		b.Remove(v)
	}
	// fmt.Println("binary-serach-tree", time.Now().Sub(t1))
	// b.PreOrder()
	// b.LevelOrder()
}
