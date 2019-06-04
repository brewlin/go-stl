package test

import (
	"hash/fnv"
	"testing"

	stlhash "github.com/brewlin/go-stl/hash"
)

func hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func BenchmarkHashOffical(bt *testing.B) {
	r := "helloworld hellowrld"
	bt.ReportAllocs()
	bt.ResetTimer()
	for i := 0; i < bt.N; i++ {
		hash(r)
	}
}

func BenchmarkHashStl(bt *testing.B) {
	r := "helloworld hellowrld"
	bt.ReportAllocs()
	bt.ResetTimer()
	for i := 0; i < bt.N; i++ {
		stlhash.Strhash(r)
	}
}
