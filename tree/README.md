# Tree 树结构 相关的树结构实现

## @bs 平衡二分搜索树
> binary-search-tree 二分搜索树[binary-search-tree](./bs)
```go
import "github.com/brewlin/go-stl/tree/bs"
bst := bs.NewBSTree()
```
## @rb 红黑树
> red-black-tree 红黑树[red-black-tree](./rb)
```go
import "github.com/brewlin/go-stl/tree/rb"
rb := rb.NewRBTree(less,more,equal)
```
## @sg 线段树
>  线段树[segment-tree](./sg)
```go
import "github.com/brewlin/go-stl/tree/sg"
rb := rb.NewRBTree(less,more,equal)
```
## @bm b-树
> b树&b-树[b-minus-tree](./bm)
```go
import(
    "github.com/brewlin/go-stl/tree/bm"
)
b := bm.NewBMTree(4)//传入树的阶乘，表示几阶b树
```