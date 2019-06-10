package test

import (
	"testing"

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
func TestRBTreeKeySet(t *testing.T) {
	b := rb.NewRBTree(Less, More, Equal)
	b.Add(1, 1)
	b.Add(2, 2)
	r := b.KeySet()
	if r[0] != 2 {
		t.Error("keyset error")
	}

}
func BenchmarkRBTree(bt *testing.B) {
	// arr := RandIntArr(bt.N, 0, bt.N)
	b := rb.NewRBTree(Less, More, Equal)
	bt.ReportAllocs()
	bt.ResetTimer()
	// for _, v := range arr {
	for i := 0; i < bt.N; i++ {
		b.Add(i, i)
		// b.Add(v, v)
	}
	// for _, v := range arr {
	for i := 0; i < bt.N; i++ {
		b.Remove(i)
	}
	// b.Prints()
	// fmt.Println(b.GetSize())
	// b.PreOrder()
	// b.LevelOrder()
}
