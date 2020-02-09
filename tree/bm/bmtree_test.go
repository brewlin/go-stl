package bm_test

import (
	"testing"

	"github.com/brewlin/go-stl/tree/bm"
)

func TestBMTreeInsert(t *testing.T) {
	newTree := bm.NewBMTree(4)
	newTree.Insert(40)
	newTree.Insert(20)
	newTree.Insert(39)
	newTree.Insert(41)
	newTree.Insert(42)
	newTree.Insert(46)
	newTree.Insert(43)
	newTree.Insert(44)
	if !newTree.Search(44) {
		t.Error("insert error")
	}
}

func TestBMTreeDelete(t *testing.T) {
	newTree := bm.NewBMTree(4)
	newTree.Insert(46)
	newTree.Insert(43)
	newTree.Insert(44)

	newTree.Delete(44)
	newTree.Delete(43)
	if newTree.Search(44) {
		t.Error("delete error")
	}
}
func BenchmarkBMTree(bt *testing.B) {
	b := bm.NewBMTree(10)
	bt.ReportAllocs()
	bt.ResetTimer()
	for i := 1; i < bt.N; i++ {
		b.Insert(i)
	}
	for i := 1; i < bt.N; i++ {
		b.Delete(i)
	}
}
