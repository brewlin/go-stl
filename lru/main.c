#include<stdio.h>
#include<stdlib.h>
#include"LRUCache.h" 

/**
 * 错误处理宏
 */
#define HANDLE_ERROR(msg) \
        do{ fprintf(stderr,"% fail.\n",msg);exit(-1);}while(0)

/**
 * 封装缓存数据储存接口，此处我们让data同事充当key的角色
 */
#define LRUCACHE_PUTDATA(cache,data) \
do { \
    if (0 != LRUCacheSet(cache,data,data)) \
        fprintf(stderr,"put (%c,%c) to cache fail.\n",data,data); \
    else \
        fprintf(stdout,"put (%c,%c) to cache success.\n",data,data); \
}while(0)

/**
 * 封装缓存数据储存接口
 */
#define LRUCACHE_GETDATA(cache,key) \
do{\
    char data = LRUCacheGet(cache,key); \
    if ('\0' == data ) \
        fprintf(stderr,"get data (key:%c) from cache fail.\n",key); \
    else if (key == data) \
        fprintf(stdout,"got (%c,%c) from cache\n",key,data); \
}while(0)

/**
 * 测试用例
 */
void testcase1(void){
    fprintf(stdout, "=========================\n");
    fprintf(stdout, "In testcase1....\n");
    fprintf(stdout, "=========================\n");
    void *lruCache;
    if (0 != LRUCacheCreate(5, &lruCache)) 
        HANDLE_ERROR("LRUCacheCreate");
    /*ABC!*/
    LRUCACHE_PUTDATA(lruCache, 'A');
    LRUCACHE_GETDATA(lruCache, 'A');
    LRUCACHE_PUTDATA(lruCache, 'B');
    LRUCACHE_GETDATA(lruCache, 'B');
    LRUCACHE_PUTDATA(lruCache, 'C');
    LRUCACHE_GETDATA(lruCache, 'C');
    LRUCachePrint(lruCache);/*CBA*/

    /*DEAF!*/
    LRUCACHE_PUTDATA(lruCache, 'D');
    LRUCACHE_GETDATA(lruCache, 'D');
    LRUCACHE_PUTDATA(lruCache, 'E');
    LRUCACHE_GETDATA(lruCache, 'E');
    LRUCACHE_PUTDATA(lruCache, 'A');
    LRUCACHE_GETDATA(lruCache, 'A');
    LRUCACHE_PUTDATA(lruCache, 'F');
    LRUCACHE_GETDATA(lruCache, 'F');
    LRUCachePrint(lruCache); /*FAEDC*/

    /*B!*/
    LRUCACHE_PUTDATA(lruCache, 'F');
    LRUCACHE_GETDATA(lruCache, 'F');
    LRUCachePrint(lruCache); /*FAEDC*/

    if (0 != LRUCacheDestroy(lruCache))
        HANDLE_ERROR("LRUCacheDestroy");
    fprintf(stdout, "\n\ntestcase1 finished\n");
    fprintf(stdout, "=========================\n\n");
}
/**
 * 测试用例2
 */
void testcase2(void){
    fprintf(stdout, "=========================\n");
    fprintf(stdout, "In testcase2....\n");
    fprintf(stdout, "=========================\n");
    void *lruCache;
    if (0 != LRUCacheCreate(3, &lruCache)) 
        HANDLE_ERROR("LRUCacheCreate");

    /*WXWYZ!*/
    LRUCACHE_PUTDATA(lruCache, 'W');
    LRUCACHE_PUTDATA(lruCache, 'X');
    LRUCACHE_PUTDATA(lruCache, 'W');
    LRUCACHE_PUTDATA(lruCache, 'Y');
    LRUCACHE_PUTDATA(lruCache, 'Z');
    LRUCachePrint(lruCache);/*ZYW*/

    LRUCACHE_GETDATA(lruCache, 'Z');
    LRUCACHE_GETDATA(lruCache, 'Y');
    LRUCACHE_GETDATA(lruCache, 'W');
    LRUCACHE_GETDATA(lruCache, 'X');
    LRUCACHE_GETDATA(lruCache, 'W');
    LRUCachePrint(lruCache);/*WYZ*/

    /*YZWYX!*/
    LRUCACHE_PUTDATA(lruCache, 'Y');
    LRUCACHE_PUTDATA(lruCache, 'Z');
    LRUCACHE_PUTDATA(lruCache, 'W');
    LRUCACHE_PUTDATA(lruCache, 'Y');
    LRUCACHE_PUTDATA(lruCache, 'X');
    LRUCachePrint(lruCache); /*XYW*/


    LRUCACHE_GETDATA(lruCache, 'X');
    LRUCACHE_GETDATA(lruCache, 'Y');
    LRUCACHE_GETDATA(lruCache, 'W');
    LRUCACHE_GETDATA(lruCache, 'Z');
    LRUCACHE_GETDATA(lruCache, 'Y');
    LRUCachePrint(lruCache); /*WYX*/

    /*XYXY!*/
    LRUCACHE_PUTDATA(lruCache, 'X');
    LRUCACHE_PUTDATA(lruCache, 'Y');
    LRUCACHE_PUTDATA(lruCache, 'X');
    LRUCACHE_PUTDATA(lruCache, 'Y');
    LRUCachePrint(lruCache);/*YX*/

    LRUCACHE_GETDATA(lruCache, 'Y');
    LRUCACHE_GETDATA(lruCache, 'X');
    LRUCACHE_GETDATA(lruCache, 'Y');
    LRUCACHE_GETDATA(lruCache, 'X');
    LRUCachePrint(lruCache); /*XY*/

    if (0 != LRUCacheDestroy(lruCache))
        HANDLE_ERROR("LRUCacheDestroy");
    fprintf(stdout, "\n\ntestcase2 finished\n");
    fprintf(stdout, "=========================\n\n");    
}
void testcase3(void)
{
    fprintf(stdout, "=========================\n");
    fprintf(stdout, "In testcase3....\n");
    fprintf(stdout, "=========================\n");

    void *lruCache;
    if (0 != LRUCacheCreate(5, &lruCache)) 
        HANDLE_ERROR("LRUCacheCreate");

    /*EIEIO!*/
    LRUCACHE_PUTDATA(lruCache, 'E');
    LRUCACHE_PUTDATA(lruCache, 'I');
    LRUCACHE_PUTDATA(lruCache, 'E');
    LRUCACHE_PUTDATA(lruCache, 'I');
    LRUCACHE_PUTDATA(lruCache, 'O');
    LRUCachePrint(lruCache);/*OIE*/


     LRUCACHE_GETDATA(lruCache, 'A');
    LRUCACHE_GETDATA(lruCache, 'I');
    LRUCACHE_GETDATA(lruCache, 'B');
    LRUCACHE_GETDATA(lruCache, 'O');
    LRUCACHE_GETDATA(lruCache, 'C');
    LRUCACHE_PUTDATA(lruCache, 'E');
    LRUCachePrint(lruCache); /*EOI*/

    if (0 != LRUCacheDestroy(lruCache))
        HANDLE_ERROR("LRUCacheDestroy");
    fprintf(stdout, "\n\ntestcase3 finished\n");
    fprintf(stdout, "=========================\n\n");
}

int main(){
    testcase1();
    testcase2();
    testcase3();
    return 0;
}
