package double

import "fmt"

//Node node
type node struct {
	key   interface{}
	value interface{}
	next  *node
	prev  *node
}

//NewNode create one node
func newNode(key, value interface{}) *node {
	var nd node
	nd.value = value
	nd.key = key
	nd.next = nil
	nd.prev = nil
	return &nd
}

type List struct {
	head  *node
	tail  *node
	size  int
	equal func(a, b interface{}) bool
}

func (l List) GetSize() int {
	return l.size
}

func NewList(equal func(a, b interface{}) bool) *List {
	var list List
	list.head = nil
	list.size = 0
	list.equal = equal
	return &list
}
// Insert push on node to list tail
func (l *List) Insert(key,value interface{}){
	nd := newNode(key,value)
	if l.head == nil || l.tail == nil {
		l.head = nd
		l.tail = nd
	}else{
		l.tail.next = nd
		nd.prev = l.tail
		l.tail = nd
	}
	l.size ++

}
//InsertToHead  push node to list head
func (l *List) InsertToHead(key, value interface{}) {
	node := newNode(key, value)
	if l.head == nil || l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		node.prev = nil
		l.head.prev = node
		l.head = node
	}
	l.size++
}
//Print print all node
func (l List) Print() {
	temp := l.head
	for temp.next != nil {
		fmt.Printf("%d", temp.next.key)
		temp = temp.next
	}
	fmt.Println()
}
//Find 根据key  找到对应的链表node 并返回value
//@return interface{}
func (l List) Find(key interface{}) interface{} {
	if l.head == nil {
		return nil
	}
	temp := l.head
	for temp != nil && temp.next != nil && !l.equal(temp.key, key) {
		temp = temp.next
	}
	if l.equal(temp.key, key) {
		return temp.value
	}
	return nil
}
//DeleteTail  删除尾节点
func (l *List) DeleteTail() *node {
	if l.size <= 0 || l.head == nil || l.tail == nil {
		return nil
	}
	tail := l.tail
	l.tail = tail.prev
	l.tail.next = nil
	tail.prev = nil
	l.size--
	return tail

}
//Delete 根据key 进行删除
func (l *List) Delete(key interface{}) {
	if l.size <= 0 || l.head == nil || l.tail == nil {
		return
	}
	if l.equal(l.head.key, key) && l.equal(l.tail.key, key) {
		l.head = nil
		l.tail = nil
	}
	temp := l.head
	for temp.next != nil && !l.equal(temp.next.key, key) {
		temp = temp.next
	}
	if temp.next == nil {
		return
	}
	if temp.next.key == key {
		delNode := temp.next
		temp.next = delNode.next
		delNode.next.prev = temp
	}
	l.size--
}

//UpdateToHead 将key 节点更新为head 节点
func (l *List) UpdateToHead(key interface{}) {
	node := l.find(key)
	if node == nil || l.head == nil || l.tail == nil || node.prev == nil{
		return
	}
	if node.prev != nil && node.next != nil {
		node.prev.next = node.next
		node.next.prev = node.prev
	}else if node.next == nil {
		node.prev.next = nil
		node.prev = nil
	}
	node.next = l.head
	l.head.prev = node
	l.head = node
}
func (l List) find(key interface{}) *node {
	if l.head == nil {
		return nil
	}
	temp := l.head
	for temp != nil &&temp.next != nil && temp.key != key {
		temp = temp.next
	}
	if l.equal(temp.key, key) {
		return temp
	}
	return nil
}

//Update 更新key 对应的value
func (l *List) Update(key, value interface{}) {
	if l.head == nil {
		return
	}
	temp := l.head
	for temp != nil && temp.next != nil && !l.equal(temp.key, key) {
		temp = temp.next
	}
	if l.equal(temp.key, key) {
		temp.value = value
	}
}

//PopTail 删除尾节点  返回删除节点的key
func (l *List) PopTail() interface{} {
	if l.head == nil || l.tail == nil {
		return nil
	}
	rmNode := l.tail
	if l.head == l.tail{
		l.head = nil
		l.tail = nil
		l.size--
		return rmNode.value
	}
	l.tail = rmNode.prev
	rmNode.prev.next = nil
	rmNode.prev = nil
	l.size--
	return rmNode.key
}
//PopTail 删除尾节点  返回删除节点的key
func (l *List) PopHead() interface{} {
	if l.head == nil || l.tail == nil {
		return nil
	}
	rmNode := l.head
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		l.size--
		return rmNode.value
	}
	l.head = rmNode.next
	rmNode.next = nil
	l.head.prev = nil
	l.size--
	return rmNode.key
}
