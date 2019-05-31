#REB-BLACK TREE
>红黑树

## struct
>结构体 的 key 是一个接口 需要提供 `Less() Mor() Equal()`等可以比较的方法，在红黑树get 和 set 的时候会进行判断
```go
type Node struct {
	key   component.Key
	value interface{}
	left  *Node
	right *Node
	//default RED
	color bool
}
```

## 提供的方法
### @key 接口实现
>自定义的key需要实现 接口的三个可比较方法 `Less() More() Equal()`
>初始化 new 一个 红黑树
```go
func Less(a, b interface{}) bool {
	return a.(int) < b.(int)
}
func More(a, b interface{}) bool {
	return a.(int) > b.(int)
}
func Equal(a, b interface{}) bool {
	return a.(int) == b.(int)
}
//初始化
rbtree := rb.NewRBTree(less,more,equal)
```
### @Add()
```go
//add 添加节点
for i:= 0 ; i < 10 ; i++{
    rbtree.Add(i,i)
}
```
### @Remove()
```go
//删除节点
for i:= 0 ; i < 10 ; i++{
    rbtree.Remove(i)
}
```
### @Contains 是否存在
```go
rbtree.Contains(key)//bool
```
### @IsEmpty()bool
### @GetSize()int
### @Set(key,value)bool 更新value