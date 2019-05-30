package example

import "math/rand"

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
	for i := 0; i < num; i++ {
		arr[i] = RandInt(min, max)
	}
	return arr
}
func RandInt(min, max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min

}
