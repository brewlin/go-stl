package test

import (
	"testing"

	hashlist "github.com/brewlin/go-stl/hash"
)

func TestHashAdd(t *testing.T) {
	list := hashlist.NewHashList(10)
	list.Insert("sdf", "xiaodo")
	res := list.Get("sdf")
	if res != "sdf" {
		t.Error("查找key 失败 hash list表 insert 失败")
	}

}
