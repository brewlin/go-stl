package trie

import (
	"github.com/brewlin/go-stl/tree/rb"
)

type node struct {
	isWord bool
	next   *rb.RBTree
}

func less(a, b interface{}) bool {
	return a.(int32) < b.(int32)
}
func more(a, b interface{}) bool {
	return a.(int32) > b.(int32)
}
func equal(a, b interface{}) bool {
	return a.(int32) == b.(int32)
}
func newNode() *node {
	return &node{
		isWord: false,
		next:   rb.NewRBTree(less, more, equal),
	}
}

//Trie the datastruct trie base rbtree
type Trie struct {
	root *node
	size int
}

func NewTrie() *Trie {
	return &Trie{
		root: newNode(),
		size: 0,
	}
}

//Add add one string
func (t *Trie) Add(word string) {
	cur := t.root
	for _, v := range word {
		if cur.next.Get(v) == nil {
			cur.next.Add(v, newNode())
		}
		cur = cur.next.Get(v).(*node)
	}
	if !cur.isWord {
		cur.isWord = true
		t.size++
	}
}

//Contains check exist the word
func (t Trie) Contains(word string) bool {
	cur := t.root
	for _, v := range word {
		if cur.next.Get(v) == nil {
			return false
		}
		cur = cur.next.Get(v).(*node)
	}
	return cur.isWord
}

//IsPrefix pre search
func (t Trie) IsPrefix(pre string) bool {
	cur := t.root
	for _, v := range pre {
		if cur.next.Get(v) == nil {
			return false
		}
		cur = cur.next.Get(v).(*node)
	}
	return true
}

//Search check exist
func (t Trie) Search(word string) bool {
	return t.match(t.root, word, 0)
}
func (t Trie) match(nd *node, word string, i int) bool {
	if i == len(word) {
		return nd.isWord
	}
	c := int32(word[i])

	if string(c) != "." {
		if nd.next.Get(c) == nil {
			return false
		}
		return t.match(nd.next.Get(c).(*node), word, i+1)
	} else {
		for _, v := range nd.next.KeySet() {
			if t.match(nd.next.Get(v).(*node), word, i+1) {
				return true
			}
			return false
		}

	}
	return false
}
