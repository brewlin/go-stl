# BINARY-SEARCHER-TREE
>二分搜索树

## struct
> 基础版的实现，用于展示，因为在多种情况下效率比较差，没有`自平衡`，会`退化`成链表等,建议使用红黑树，或者hash表等，
```go
type Node struct {
	e     int
	left  *Node
	right *Node
}
//BSTree struct
type BSTree struct {
	root *Node
	size int
}
```

## 提供的方法
### @初始化`
```go
//初始化
import "github.com/brewlin/go-stl/tree/bs"
rbtree := bs.NewRBTree(less,more,equal)
```
### @Add()
```go
//add 添加节点
for i:= 0 ; i < 10 ; i++{
    bstree.Add(i)
}
```
### @Remove()
```go
//删除节点
for i:= 0 ; i < 10 ; i++{
    bstree.Remove(i)
}
```
### @Contains 是否存在
```go
bstree.Contains(key)//bool
```
### @IsEmpty()bool
### @GetSize()int
### @Set(key,value)bool 更新value