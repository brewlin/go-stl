package list

import "fmt"

//Node node
type Node struct {
	key   int
	value interface{}
	next  *Node
}

//NewNode create one node
func NewNode(key int, value interface{}) *Node {
	var node Node
	node.value = value
	node.key = key
	node.next = nil
	return &node
}

type List struct {
	head *Node
	size int
}

func (l List) GetSize() int {
	return l.size
}

func NewList() *List {
	var list List
	list.head = NewNode(-1, nil)
	list.size = 0
	return &list
}
func (l *List) Insert(key int, value interface{}) {
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
func (l List) Find(key int) interface{} {
	temp := l.head
	for temp.next != nil && temp.next.key != key {
		temp = temp.next
	}
	if temp.next.key == key {
		return temp.next.value
	}
	return nil
}
func (l *List) Delete(key int) {
	temp := l.head
	for temp.next != nil && temp.next.key != key {
		temp = temp.next
	}
	if temp.next == nil {
		return
	}
	if temp.next.key == key {
		temp.next = temp.next.next
	}

}
