package bm

import "fmt"

//BMTree b(minus)- tree
type BMTree struct {
	root *BMNode
	m    int
}

func NewBMTree(m int) *BMTree {
	b := BMTree{}
	b.m = m
	b.root = NewBMNode(m)
	return &b
}
func (b *BMTree) Search(v int) bool {
	isexist, _, _ := b.search(v)
	if isexist {
		return true
	}
	return false
}
func (b *BMTree) search(v int) (bool, int, *BMNode) {
	n := NewBMNode(b.m)
	var i int

	t := b.root
	if t == nil {
		return false, 0, nil
	}
	for t != nil {
		for i = t.size; i > 0 && v <= t.key[i]; i-- {
			if v == t.key[i] {
				return true, i, t
			}
		}
		if t.child[i] == nil {
			n = t
		}
		t = t.child[i]
	}
	return false, i, n
}

func (b *BMTree) Insert(v int) *BMNode {
	var i int
	ok, _, node := b.search(v)
	if !ok {
		node.key[0] = v
		for i = node.size; i > 0 && v < node.key[i]; i-- {
			node.key[i+1] = node.key[i]
		}
		node.key[i+1] = v
		node.size++
		if node.size < b.m {
			return b.root
		} else {
			parent := node.split(b.m)
			for parent.parent != nil {
				parent = parent.parent
			}
			b.root = parent
			return b.root
		}
	}
	return b.root
}
func (b *BMTree) Delete(v int) *BMTree {
	ok, i, node := b.search(v)
	if ok {
		b.root = node.deleteNode(v, i, b.m)
	}
	return b
}

func (b *BMTree) String() string {
	q := []*BMNode{}
	q = append(q, b.root)
	i := 0
	for i < len(q) {
		c := q[i]
		i += 1
		fmt.Print("[")
		for k := 1; k <= c.size; k++ {
			fmt.Printf(" %d ", c.key[k])
		}
		fmt.Print("]\n")
		for k := 0; k <= c.size; k++ {
			if c.child[k] != nil {
				q = append(q, c.child[k])
			}
		}
	}
	return ""
}
