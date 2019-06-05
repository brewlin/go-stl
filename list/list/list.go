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

func NewList(equal func(a, b interface{})bool) *List {
	var list List
	list.head = nil
	list.size = 0
	list.equal = equal
	return &list
}
func (l *List) Insert(key, value interface{}) {
	node := NewNode(key, value)
	if l.head == nil {
		l.head = node
		l.size++
		return
	}
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
func (l *List) Find(key interface{}) interface{} {
	if l.head == nil {
		return nil
	}
	temp := l.head
	for temp.next != nil && !l.equal(temp.key, key) {
		temp = temp.next
	}
	if l.equal(temp.key, key) {
		return temp.value
	}
	return nil
}
func (l List) Update(key, value interface{}) {
	temp := l.head
	for temp.next != nil && !l.equal(temp.key, key) {
		temp = temp.next
	}
	if l.equal(temp.key, key) {
		temp.value = value
	}
}
func (l *List) Delete(key interface{}) {
	if l.head == nil {
		return
	}
	temp := l.head
	if l.equal(l.head.key,key){
		l.head = temp.next
		return
	}
	for temp.next != nil && !l.equal(temp.next.key, key) {
		temp = temp.next
	}
	if temp == nil || temp.next == nil {
		return
	}
	if l.equal(temp.next.key, key) {
		temp.next = temp.next.next
	}

}
