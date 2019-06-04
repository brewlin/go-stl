package double

import "fmt"

//Node node
type Node struct {
	key   interface{}
	value interface{}
	next  *Node
	prev  *Node
}

//NewNode create one node
func NewNode(key, value interface{}) *Node {
	var node Node
	node.value = value
	node.key = key
	node.next = nil
	node.prev = nil
	return &node
}

type List struct {
	head  *Node
	tail  *Node
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
func (l *List) InsertToHead(key, value interface{}) {
	node := NewNode(key, value)
	if l.head == nil && l.tail == nil {
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
func (l List) Print() {
	temp := l.head
	for temp.next != nil {
		fmt.Printf("%d", temp.next.key)
		temp = temp.next
	}
	fmt.Println()
}
func (l List) Find(key interface{}) interface{} {
	temp := l.head
	for temp.next != nil && !l.equal(temp.next.key, key) {
		temp = temp.next
	}
	if l.equal(temp.next.key, key) {
		return temp.next.value
	}
	return nil
}
func (l *List) DeleteTail() *Node {
	if l.size <= 0 || l.head == nil || l.tail == nil {
		return nil
	}
	tail := l.tail
	l.tail = tail.prev
	tail.prev.next = nil
	tail.prev = nil
	l.size--
	return tail

}
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
	if node == nil {
		return
	}
	//如果节点为 尾节点
	if node == l.tail {
		l.tail = node.prev
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	node.prev = nil
	node.next = l.head
	l.head.prev = node
	l.head = node
}
func (l List) find(key interface{}) *Node {
	temp := l.head
	for temp.next != nil && temp.next.key != key {
		temp = temp.next
	}
	if l.equal(temp.next.key, key) {
		return temp.next
	}
	return nil
}

//Update 更新key 对应的value
func (l *List) Update(key, value interface{}) {
	temp := l.head
	for temp.next != nil && !l.equal(temp.next.key, key) {
		temp = temp.next
	}
	if l.equal(temp.next.key, key) {
		temp.next.value = value
	}
}

//PopTail 删除尾节点  返回删除节点的key
func (l *List) PopTail() interface{} {
	if l.head == nil || l.tail == nil {
		return nil
	}
	rmNode := l.tail
	l.tail = rmNode.prev
	rmNode.prev.next = nil
	rmNode.prev = nil
	l.size--
	return rmNode.key
}
