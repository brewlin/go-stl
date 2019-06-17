# SEGMENT- TREE
>线段树

## struct
> 线段数实际是 以数组等方式来存储二叉树，因为数组等索引访问速度为o（1）
```go
type SGTree struct {
	tree   []interface{}
	data   []interface{}
	merger func(a, b interface{}) interface{}
}

```

## 提供的方法
### @初始化线段树
> 需要传递merger 回调函数
```go
func merger(a, b interface{}) bool {
	return a.(int) + b.(int)
}
//初始化
arr := []int{3,43,23,43,3,43,4}
sgtree := sg.NewSGTree(arr,merger)

```
### @Query()
> 范围查询 例如查询 该范围 2-4 位范围的和
```go
//tree => [3,41,4,35,54,6,23,35,4,4,2]
sgtree.Query(2,4) => 4+ 35 + 54 = 93
```