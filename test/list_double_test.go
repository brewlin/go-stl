package test

import (
	"testing"

	"github.com/brewlin/go-stl/list/double"
)
func equal(a,b interface{})bool{
	return a.(int) == b.(int)
}
//测试 双向链表的插入 数据
func TestListDoubleInsert(t *testing.T) {
	list := double.NewList(equal)
	list.Insert(1,1)
}
//测试 双向链表的插入 数据 到头节点
func TestListDoubleInsertToHead(t *testing.T) {
	list := double.NewList(equal)
	list.InsertToHead(1,1)
}

//测试 双向链表的查找数据
func TestListDoubleFind(t *testing.T) {
	list := double.NewList(equal)
	list.Insert(3,3)
	list.Insert(2,2)
	list.InsertToHead(5,5)
	res := list.Find(3)
	if res != 3 {
		t.Error("double list  查找失败")
	}
}
//测试 双向链表的删除
func TestListDoubleDelete(t *testing.T) {
	list := double.NewList(equal)
	list.Insert(3,3)
	list.Insert(2,2)
	list.InsertToHead(5,5)
	list.Delete(3)
	res := list.Find(3)
	if res != nil {
		t.Error("double list  根据key 删除node失败")
	}
}
//测试 双向链表的删除 尾节点
func TestListDoubleDeleteTail(t *testing.T) {
	list := double.NewList(equal)
	list.Insert(3,3)
	list.Insert(2,2)
	list.InsertToHead(5,5)
	list.DeleteTail()
	res := list.Find(2)
	if res != nil {
		t.Error("double list  根据key 删除node失败")
	}
}
func TestListDoubleUpdate(t *testing.T) {
	list := double.NewList(equal)
	list.Insert(3,3)
	list.Update(4,4)
	list.Update(3,33)
	res := list.Find(3)
	if res != 33 {
		t.Error("doublie list 双向链表更新失败")
	}
}
func TestListDoublePopHead(t *testing.T){
	list := double.NewList(equal)
	list.Insert(3,3)
	list.Insert(4,4)
	list.InsertToHead(5,5)
	res := list.PopHead()
	if res != 5 {
		t.Error("double list 删除双向链表 头结点失败")
	}
}

func TestListDoubleUpdateToTail(t *testing.T) {
	list := double.NewList(equal)
	list.Insert(3,3)
	list.UpdateToHead(4)
	res := list.Find(4)
	if res != nil {
		t.Error("doublie list 双向链表更新节点到 head 节点失败, 不应该将空的节点 不存在的节点更新到head去")
	}
	list.UpdateToHead(3)
	res = list.PopHead()
	if res != 3 {
		t.Error("doublie list 双向链表更新节点到 head 节点失败")
	}
}