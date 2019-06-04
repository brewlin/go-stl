package list

import "fmt"

//Node node
type Node struct {
	key   interface{}
	value interface{}
	next  *Node
}

//NewNode create one node
func NewNode(key, value interface{}) *Node {
	var node Node
	node.value = value
	node.key = key
	node.next = nil
	return &node
}

type List struct {
	head  *Node
	size  int
	equal func(a, b interface{}) bool
}

func (l List) GetSize() int {
	return l.size
}

func NewList() *List {
	var list List
	list.head = nil
	list.size = 0
	return &list
}
func (l *List) Insert(key, value interface{}) {
	node := NewNode(key, value)
	temp := l.head
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = node
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
func (l List) Update(key, value interface{}) {
	temp := l.head
	for temp.next != nil && !l.equal(temp.next.key, key) {
		temp = temp.next
	}
	if l.equal(temp.next.key, key) {
		temp.next.value = value
	}
}
func (l *List) Delete(key interface{}) {
	temp := l.head
	for temp.next != nil && !l.equal(temp.next.key, key) {
		temp = temp.next
	}
	if temp.next == nil {
		return
	}
	if l.equal(temp.next.key, key) {
		temp.next = temp.next.next
	}

}
