package lru

//LRUCaches struct
type LRUCaches struct {
	cacheCapcity int          //容量
	hashmap      *cacheEntrys //缓存的hash表
	lruListHead  *cacheEntrys //缓存双向链表表头
	lruListTail  *cacheEntrys //缓存双向链表表尾
	lruListSize  int          //节点个数
}

type cacheEntrys struct {
	key          string
	data         string
	hashListPrev *cacheEntrys //hash链表的前一个元素
	hashListNext *cacheEntrys //hash链表的后一个元素
	lruListPrev
}
