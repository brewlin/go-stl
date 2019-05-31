package example

import (
	"math/rand"
	"time"
)

func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}
func RandInt64Arr(num int) []int64 {
	arr := make([]int64, num)
	for i := 0; i < num; i++ {
		arr[i] = RandInt64(100, 10000)
	}
	return arr
}
func RandIntArr(num, min, max int) []int {
	arr := make([]int, num)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < num; i++ {
		arr[i] = r.Intn(max - min)
	}
	return arr
}
func RandInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max)
}
