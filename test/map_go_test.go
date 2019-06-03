package test

import (
	"testing"
)

func BenchmarkGOMap(b *testing.B) {

	// t1 := time.Now()

	arr := RandIntArr(b.N, 0, b.N)
	b.ReportAllocs()
	b.ResetTimer()
	m := make(map[int]interface{})
	// for i := 0; i < n; i++ {
	for _, v := range arr {
		m[v] = v
	}
	// for i := 0; i < n; i++ {
	for _, v := range arr {
		delete(m, v)
	}
	// fmt.Println("go map", time.Now().Sub(t1))
}
