package test

import (
	"fmt"
	"testing"

	"github.com/brewlin/go-stl/tree/sg"
)

func merger(a, b interface{}) interface{} {
	return a.(int) + b.(int)
}
func TestSGTBuild(t *testing.T) {
	arr := []interface{}{1, 3}
	s := sg.NewSGTree(arr, merger)
	if s == nil {
		t.Error("segment tree build falid")
	}

}

func TestSGTQuery(t *testing.T) {
	arr := []interface{}{1, 3, 5}
	s := sg.NewSGTree(arr, merger)
	re, _ := s.Query(1, 2)
	fmt.Println(re)
	if re.(int) != 8 {
		t.Error("segment tree query falid")
	}

}
