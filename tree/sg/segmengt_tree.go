package sg

import (
	"errors"
)

type SGTree struct {
	tree   []interface{}
	data   []interface{}
	merger func(a, b interface{}) interface{}
}

// NewSGTree
func NewSGTree(data []interface{}, merger func(a, b interface{}) interface{}) *SGTree {
	var tree = new(SGTree)
	tree.merger = merger
	tree.data = data
	tree.tree = make([]interface{}, 4*len(data))
	tree.build(0, 0, len(tree.data)-1)
	return tree
}

/**
 * private func
 * slice convert to SGTree by recursive
 * @param treeindex
 * @param l  left
 * @param r  right
 */
func (s *SGTree) build(treeindex, l, r int) {
	if l == r {
		s.tree[treeindex] = s.data[l]
		return
	}
	leftTreeIndex := s.leftChild(treeindex)
	rightTreeIndex := s.rightChild(treeindex)
	//mid = (l  + r )/ 2  notice: it's easy to core dump while hug int l and r so we do s:
	mid := l + (r-l)/2
	//recursive
	s.build(leftTreeIndex, l, mid)
	s.build(rightTreeIndex, mid+1, r)
	s.tree[treeindex] = s.merger(s.tree[leftTreeIndex], s.tree[rightTreeIndex])
}

/**
 * public func
 * interval query => query [0-5]
 * @params l r => left ,right
 */
func (s SGTree) Query(l, r int) (interface{}, error) {
	if l < 0 || l > len(s.data) || l > r || r < 0 || r > len(s.data) {
		return nil, errors.New("params invalid")
	}
	return s.query(0, 0, len(s.data)-1, l, r), nil
}

/**
 * private query recursive
 * params => treeIndex l=>0,r=>len(data),ql =>queryl,qr => queryr
 */
func (s SGTree) query(treeIndex, l, r, ql, qr int) interface{} {
	if l == ql && r == qr {
		return s.tree[treeIndex]
	}
	mid := l + (r-l)/2
	leftTreeIndex := s.leftChild(treeIndex)
	rightTreeIndex := s.rightChild(treeIndex)
	if ql >= mid+1 {
		return s.query(rightTreeIndex, mid+1, r, ql, qr)
	} else if qr <= mid {
		return s.query(leftTreeIndex, l, mid, ql, qr)
	}
	lefRes := s.query(leftTreeIndex, l, mid, ql, mid)
	rigRes := s.query(rightTreeIndex, mid+1, r, mid+1, qr)
	return s.merger(lefRes, rigRes)
}

/**
 * set data and recuresive
 */
func (s *SGTree) Set(i int, e interface{}) {
	s.data[i] = e
	s.set(0, 0, len(s.data)-1, i, e)
}
func (s *SGTree) set(treeIndex, l, r, i int, e interface{}) {
	if l == r {
		s.tree[treeIndex] = e
		return
	}
	mid := l + (r-l)/2
	lt := s.leftChild(treeIndex)
	rt := s.rightChild(treeIndex)
	if i >= mid+1 {
		s.set(rt, mid+1, r, i, e)
	} else {
		s.set(lt, l, mid, i, e)
	}
	s.tree[treeIndex] = s.merger(s.tree[lt], s.tree[rt])
}
func (s SGTree) GetSize() int {
	return len(s.data)
}
func (s SGTree) Get(i int) interface{} {
	if i < 0 || i >= len(s.data) {
		return nil
	}
	return s.data[i]
}

/**
 * get the left node index
 */
func (s SGTree) leftChild(treeIndex int) int {
	return 2*treeIndex + 1
}

/**
 * get the right node index
 */
func (s SGTree) rightChild(treeIndex int) int {
	return 2*treeIndex + 2
}
