#ifndef __LRUCACHE_H__

#define __LRUCACHE_H__

/**
 * 创建一个缓存单元
 */
static cacheEntrys* newCacheEntry(char key,char data);
/**
 * 释放一个缓存单元
 */
static void freeCacheEntry(cacheEntrys* entry);

//create cache
int LRUCacheCreate(int capacity,void **lruCache);

//delete cache
int LRUCacheDestroy(void *lruCache);

//set data
int LRUCacheSet(void *lruCache,char key,char data);

//get data
char LRUCacheGet(void *lruCache,char key);

//print
void LRUCachePrint(void *lruCache);


#endif