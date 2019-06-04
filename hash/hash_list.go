package hash

import (
	"github.com/brewlin/go-stl/list/list"
)

type HashList struct {
	capacity int
	b        uint8 //bit
	buckets  []*list.List
}

func NewHashList(capacity int) *HashList {
	var hash HashList
	hash.b = uint8(capacity)
	hash.capacity = capacity << (hash.b - 1)
	bs := make([]*list.List, hash.capacity)
	for i := 0; i < hash.capacity; i++ {
		bs[i] = list.NewList()
	}
	hash.buckets = bs
	return &hash
}

func equal(a, b interface{}) bool {
	return a.(string) == b.(string)
}
func (h HashList) hashCode(key string) int {
	hashcode := Strhash(key)
	m := uintptr(1)<<h.b - 1
	return int(hashcode & m)
}

func (h HashList) Get(key string) interface{} {
	list := h.buckets[h.hashCode(key)]
	return list.Find(key)
}
func (h HashList) Set(key string, value interface{}) {
	list := h.buckets[h.hashCode(key)]
	list.Update(key, value)
}
func (h *HashList) Insert(key string, value interface{}) {
	list := h.buckets[h.hashCode(key)]
	list.Insert(key, value)
}
func (h *HashList) Remove(key string) {
	list := h.buckets[h.hashCode(key)]
	list.Delete(key)
}
