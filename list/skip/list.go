package skip

import (
	"fmt"
	"math/rand"
)

//Node is a node
type Node struct {
	value interface{}
	key   int
	level int
	next  []*Node
}

func (n Node) GetValue() interface{} {
	return n.value
}

//NewNode new node
func NewNode(key int, value interface{}, level int) *Node {
	var nd Node
	nd.next = make([]*Node, level)
	nd.value = value
	nd.level = level
	nd.key = key
	return &nd
}

func NewSkipList() *SkipList {
	var skip SkipList
	skip.maxLevel = 32
	skip.head = NewNode(-1, nil, 16)
	skip.size = 0
	skip.levelCount = 1
	return &skip
}

//SkipList is a struct skip_list
type SkipList struct {
	maxLevel   int   //16
	head       *Node //newnode(-1,16)
	size       int   //0
	levelCount int   //1
}

//Find find key
func (s SkipList) Find(key int) *Node {
	temp := s.head
	for i := s.levelCount - 1; i >= 0; i-- {
		for temp.next[i] != nil && temp.next[i].key < key {
			temp = temp.next[i]
		}
	}
	//check if exist
	if temp.next[0] != nil && temp.next[0].key == key {
		return temp.next[0]
	}
	return nil
}
func (s *SkipList) Insert(key int, value interface{}) {
	//get the coin level
	level := s.getlevel()
	node := NewNode(key, value, level)
	//update
	update := make([]*Node, level)

	temp := s.head
	for i := level - 1; i >= 0; i-- {
		for temp.next[i] != nil && temp.next[i].key < key {
			temp = temp.next[i]
		}
		update[i] = temp
	}
	for i := 0; i < level; i++ {
		node.next[i] = update[i].next[i]
		update[i].next[i] = node
	}
	//check need update level
	if level > s.levelCount {
		s.levelCount = level
	}
	s.size++
}

//delete
func (s *SkipList) Delete(key int) {
	update := make([]*Node, s.levelCount)

	temp := s.head
	for i := s.levelCount - 1; i >= 0; i-- {
		for temp.next[i] != nil && temp.next[i].key < key {
			temp = temp.next[i]
		}
		update[i] = temp
	}

	if temp.next[0] != nil && temp.next[0].key == key {
		s.size--
		for i := s.levelCount - 1; i >= 0; i-- {
			if update[i].next[i] != nil && update[i].next[i].key == key {
				update[i].next[i] = update[i].next[i].next[i]
			}
		}
	}
}

//print all
func (s SkipList) Print() {
	temp := s.head
	for temp.next[0] != nil {
		fmt.Println(temp.next[0].key)
		temp = temp.next[0]
	}
}

// get level by rand
func (s SkipList) getlevel() int {
	level := 1
	for rand.Float64() < 0.25 && level < s.maxLevel {
		level++
	}

	return level
}
