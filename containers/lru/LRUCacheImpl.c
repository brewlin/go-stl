#include "LRUCacheImpl.h"
#include "LRUCache.h"
#include <stdio.h>



/**
 * 创建一个缓存单元
 */
static cacheEntrys* newCacheEntry(char key,char data){
    cacheEntrys* entry = (cacheEntrys *)malloc(sizeof(sizeof(*entry)));
    if(entry == NULL){
        perror("malloc new entry point is error");
        return NULL;
    }
    memset(entry,0,sizeof(*entry));
    entry->key = key;
    entry->data = data;
    return entry;
}
/**
 * 释放一个缓存单元
 */
static void freeCacheEntry(cacheEntrys* entry){
    if(NULL == entry)return;
    free(entry);
}

/**
 * 创建一个LRU缓存
 * 1.创建主结构
 * 2.创建hashmap内存
 */
int LRUCacheCreate(int capacity,void **lruCache){
    LRUCacheS* cache = (LRUCacheS *)malloc(sizeof(*cache));
    if(cache == NULL){
        perror("malloc error");
        return -1;
    }
    memset(cache,0,sizeof(*cache));
    cache->cacheCapcity = capacity;
    cache->hashMap = malloc(sizeof(cacheEntrys) * capacity);
    if( cache->hashMap == NULL){
        free(cache);
        perror("hashmap mallock faild");
        return -1;
    }
    memset(cache->hashMap,0,sizeof(cacheEntrys) * capacity);
    *lruCache = cache;
    return 0;
}

/**
 * 释放一个LRU缓存
 */
int LRUCacheDestroy(void *lruCache){
    LRUCacheS *cache = (LRUCacheS *)lruCache;
    if(cache == NULL)return 0;
    free(cache);
    return 0;
}

/**
 * 缓存存取接口
 * 将数据放到LRU缓存中
 */
int LRUCacheSet(void *lruCache,char key , char data){
    LRUCacheS *cache = (LRUCacheS *)lruCache;

    //从hash表中查看是否存在该表中
    cacheEntrys *entry = getValueFromHashMap(cache,key);
    if(entry != NULL){
        //该数据存在，更新数据到表头
        entry->data = data;
        updateLRUList(cache,entry);
    }else{
        //新建节点
        entry = newCacheEntry(cache,data);

        //将缓存插入链表表头
        cacheEntrys *removeEntry = insertToListHead(cache,entry);
        if(NULL != removeEntry){
            //如果缓存满了
            removeEntryFromHashMap(cache,removeEntry);
            freeCacheEntry(removeEntry);
        }
        //新建节点插入hash表
        insertEntryToHashMap(cache,entry);
    }
    return 0;
}

/**
 * 从缓存中获取数据
 */
char LRUCacheGet(void *lruCache,char key){
    LRUCacheS *cache = (LRUCacheS *)lruCache; 
    cacheEntrys *entry =  getValueFromHashMap(cacheEntrys,key);
    if( NULL != entry){
        updateLRUList(cache,entry);
        return entry->data;
    }else{
        return '\0';
    }
}

/**
 * 打印
 */
void LRUCachePrint(void *lruCache){
    LRUCacheS *cache = (LRUCacheS *)lruCache;

    if(NULL == cache || 0 == cache->lruListSize)return;
    fprintf(stdout,"\n>>>>>>>>>>>\n");
    fprintf(stdout,"cache (key data):\n");
    cacheEntrys *entry = cache->lruListHead;
    while(entry){
        fprintf(stdout,"(%c,%c) ",entry->key,entry->data);
        entry = entry->lruListNext;
    }
    fprintf(stdout,"\n<<<<<<<<<<<<<<<<<<<<<<<\n\n");
}


