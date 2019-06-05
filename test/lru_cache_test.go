package test

import (
	"testing"
	"github.com/brewlin/go-stl/containers/lru"
)
func TestLRUSet(t *testing.T){
	cache := lru.NewLRUCache(10)
	if cache == nil {
		t.Error("lru缓存创建失败")
	}
}
func TestLRUGet(t *testing.T) {
	cache := lru.NewLRUCache(10)
	cache.Set("s", "sdf")
	value := cache.Get("s")
	if value != "sdf" {
		t.Error("cache 缓存 查询失败")
	}
}
//检查是否能过淘汰过期数据
func TestLRUAuto_DelteTailData(t *testing.T){
	cache := lru.NewLRUCache(2)
	cache.Set("1", "1")
	cache.Set("2", "2")
	cache.Set("3","3")//del 1
	res := cache.Get("1")
	if res != nil {
		t.Error("cache 没有淘汰掉过期数据")
	}

}
