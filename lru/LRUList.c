#include "LRUCacheImpl.h"
#include "stdio.h"

/**
 * 从双向链表中删除指定节点
 */
static void removeFromList(LRUCacheS *cache,cacheEntrys *entry){
    //null
    if(cache->lruListSize <= 0)
        return;
    if(entry == cache->lruListHead && entry == cache->lruListTail){
        //链表中仅剩当前节点
        cache->lruListHead = cache->lruListTail = NULL;
    }else if(entry == cache->lruListHead){//删除头节点
        cache->lruListHead = entry->lruListNext;
        cache->lruListHead->lruListPrev = NULL;//将prev节点去除
    }else if(entry == cache -> lruListTail){//删除尾节点
        cache->lruListTail = entry->hashListPrev;
        cache->lruListTail->lruListNext = NULL;
        entry->lruListPrev = NULL;
    }else{
        entry->lruListPrev->lruListNext = entry->lruListNext;
        entry->lruListNext->lruListPrev = entry->lruListPrev;
        entry->lruListPrev = NULL;
        entry->lruListNext = NULL;
    }
    cache->lruListSize --;
}
/**
 * 将节点插入链表表头
 */
static cacheEntrys* insertToListHead(LRUCacheS *cache,cacheEntrys *entry){
    cacheEntrys *removeEntry = NULL;
    //超过容量就删除尾节点， 淘汰过期数据
    if (++cache->lruListSize > cache->cacheCapcity){
        removeEntry = cache->lruListTail;
        removeFromList(cache,cache->lruListTail);
    }
    //如果当前链表为空
    if(cache->lruListHead == NULL && cache->lruListTail == NULL){
        cache->lruListHead = cache->lruListTail = entry;
    }else{
        //插入表头
        entry->lruListNext = cache->lruListHead;
        entry->lruListPrev = NULL;
        cache->lruListHead->lruListPrev = entry;
        cache->lruListHead = entry;

    }
    
}
/**
 * 释放整个链表
 */
static void freeList(LRUCacheS *cache){
    if( 0 == cache->lruListSize);return;
    cacheEntrys *entry = cache->lruListHead;
    //遍历删除
    while(entry){
        cacheEntrys *temp = entry->lruListNext;
        freeCacheEntry(entry);
        entry = temp;
    }
    cache->lruListSize = 0;
}

/**
 * 辅助接口，保证最近访问的节点总是位于链表表头
 */
static void updateLRUList(LRUCacheS *cache,cacheEntrys *entry){
    removeFromList(cache,entry);
    insertToListHead(cache,entry)
}