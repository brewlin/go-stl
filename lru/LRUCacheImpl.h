#ifndef __LRUCACHEIMPLE_H__
#define __LRUCACHEIMPLE_H__
/**
 * 定义LRU缓存数据结构
 */
typedef struct {
    int cacheCapcity;//容量
    cacheEntrys **hashMap;//缓存的hash表
    cacheEntrys *lruListHead;//缓存双向链表表头
    cacheEntrys *lruListTail;//缓存双向链表表尾
    int lruListSize;//节点个数
}LRUCacheS;

/**
 * 缓存区中的缓存单元
 */
typedef struct {
    char key; //数据key
    char data;//数据data

    cacheEntrys *hashListPrev;//指向hash链表的前一个元素
    cacheEntrys *hashListNext;//指向hash链表的后一个元素
    cacheEntrys *lruListPrev;//指向双向链表的前一个元素
    cacheEntrys *lruListNext;//指向双向链表的后一个元素
    
}cacheEntrys;

/**
 * 从双向链表中删除指定节点
 */
static void removeFromList(LRUCacheS *cache,cacheEntrys *entry);
/**
 * 将节点插入链表表头
 */
static cacheEntrys* insertToListHead(LRUCacheS *cache,cacheEntrys *entry);
/**
 * 释放整个链表
 */
static void freeList(LRUCacheS *cache);
/**
 * 辅助接口，保证最近访问的节点总是位于链表表头
 */
static void updateLRUList(LRUCacheS *cache,cacheEntrys *entry);
static void removeFromList(LRUCacheS *cache,cacheEntrys *entry);



#endif