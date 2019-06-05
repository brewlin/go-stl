package test

import (
	"testing"

	hashlist "github.com/brewlin/go-stl/hash"
)

func TestHashInsert(t *testing.T) {
	list := hashlist.NewHashList(10)
	list.Insert("sdf", "xiaodo")
}
func TestHashGet(t *testing.T) {
	list := hashlist.NewHashList(10)
	list.Insert("sdf", "xiaodo")
	res := list.Get("sdf")
	if res != "xiaodo" {
		t.Error("HASH list 查询失败")
	}
}
func TestHashSet(t *testing.T){
	list := hashlist.NewHashList(10)
	list.Insert("sdf", "xiaodo")
	list.Set("sdf","111")
	res := list.Get("sdf")
	if res != "111" {
		t.Error("Hash list Set 更新失败")
	}
}
// test hash remove func
func TestHashRemove(t *testing.T){
	list := hashlist.NewHashList(10)
	for i:= 1 ; i < 10 ;i ++ {
		list.Insert(string(i), string(i))
	}
	list.Remove("3")
	res := list.Get("3")
	if res != nil {
		t.Error("Hash list Set 更新失败")
	}
}



