package test

import (
	"testing"

	"github.com/brewlin/go-stl/containers/lru"
)

func TestLRUSet(t *testing.T) {
	cache := lru.NewLRUCache(10)
	if cache == nil {
		t.Error("lru缓存创建失败")
	}
	cache.Set("s", "sdf")
	value := cache.Get("s")
	if value != "sdf" {
		t.Error("cache 缓存 查询失败")
	}
}
