#include "LRUCacheImpl.h"
#include "LRUCache.h"
#include <stdio.h>
/**
 * hash key  
 * 在java里 有 hashcode()
 */
static int hashKey(LRUCacheS *cache, char key){
    //必须保证capcity 容量为 2 n方 可以进行位运算
    return (int)key & (cache->cacheCapcity - 1);
}

/**
 * 查找 value
 */
static cacheEntrys *getValueFromHashMap(LRUCacheS *cache,int key){
    //先获取存放在哪个桶里
    cacheEntrys *entry = cache->hashMap[hashKey(cache,key)];

    //遍历查询找到准确项
    while(entry){
        if(entry->key == key){
            break;
        }
        entry = entry->hashListNext;
    }
    return entry;
}

/**
 * 插入缓存单元
 */
static void insertEntryToHashMap(LRUCacheS *cache,cacheEntrys *entry)
{
    int hashcode = hashKey(cache,entry->key);
    cacheEntrys *n = cache->hashMap[hashcode];
    
    if(n != NULL){
        entry->hashListNext = n;
        n->hashListPrev = entry;
    }
    cache->hashMap[hashcode] = entry;
}

/**
 * 从hash表中删除节点
 */
static void removeEntryFromHashMap(LRUCacheS *cache ,cacheEntrys *entry){
    if(NULL == entry || NULL == cache || NULL == cache->hashMap)return;

    int hashcode = hashKey(cache,entry->key);
    cacheEntrys *n = cache->hashMap[hashKey(cache,entry->key)];

    //遍历
    while(n){
        if(n->key == entry->key){
            if(n->hashListPrev){
                n->hashListPrev->hashListNext = n->hashListNext;
            }else{
                cache->hashMap[hashcode] = n->hashListNext;
            }
            if(n->hashListNext)
                n->hashListNext->hashListPrev = n->hashListPrev;
            return
        }
        n = n->hashListNext;
    }
}




