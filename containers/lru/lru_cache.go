package lru

import (
	"errors"

	"github.com/brewlin/go-stl/hash"
	"github.com/brewlin/go-stl/list/double"
)

//LRUCaches struct
type LRUCaches struct {
	cacheCapcity int            //容量
	hashmap      *hash.HashList //缓存的hash表
	lruList      *double.List
}

func equal(a, b interface{}) bool {
	return a.(string) == b.(string)
}

//NewLruCache
func NewLRUCache(capcity int) *LRUCaches {
	var cache LRUCaches
	cache.cacheCapcity = capcity
	cache.hashmap = hash.NewHashList(5)
	cache.lruList = double.NewList(equal)
	return &cache
}

//Set 插入缓存
func (l *LRUCaches) Set(key, data string) error {
	entry := l.hashmap.Get(key)
	//缓存已存在
	if entry != nil {
		//更新value
		l.lruList.Update(key, data)
		//更新到链表表头
		l.lruList.UpdateToHead(key)
	} else { //新建节点
		//缓存满了
		if l.lruList.GetSize() >= l.cacheCapcity {
			//remove tail
			rmKey := l.lruList.PopTail()
			if rmKey == nil {
				return errors.New("缓存异常 无法删除链表尾节点")
			}
			l.hashmap.Remove(rmKey.(string))
		}
		l.lruList.InsertToHead(key, data)
		l.hashmap.Insert(key, data)
	}
	return nil
}
func (l *LRUCaches) Get(key string) interface{} {
	entry := l.hashmap.Get(key)
	if entry == nil {
		return nil
	}
	l.lruList.UpdateToHead(key)
	return entry

}
