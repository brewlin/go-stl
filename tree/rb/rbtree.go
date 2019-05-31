package rb

import (
	"fmt"
)

const (
	RED   = true
	BLACK = false
)

type Node struct {
	key   interface{}
	value interface{}
	left  *Node
	right *Node
	//default RED
	color bool
}

//NewNode return a node
func NewNode(key, value interface{}) *Node {
	node := new(Node)
	node.key = key
	node.value = value
	node.left = nil
	node.right = nil
	node.color = RED
	return node
}

//RBTree struct
type RBTree struct {
	root  *Node
	size  int
	less  func(a, b interface{}) bool
	more  func(a, b interface{}) bool
	equal func(a, b interface{}) bool
}

//NewRBTree return a tree
func NewRBTree(less, more, equal func(a, b interface{}) bool) *RBTree {
	rb := new(RBTree)
	rb.root = nil
	rb.size = 0
	rb.less = less
	rb.more = more
	rb.equal = equal
	return rb
}
func (r RBTree) GetSize() int {
	return r.size
}
func (r RBTree) IsEmpty() bool {
	return r.size == 0
}

//leftRotate
func (r *RBTree) leftRotate(node *Node) *Node {
	x := node.right
	node.right = x.left
	x.left = node
	x.color = node.color
	node.color = RED
	return x

}

//rightRotate
func (r *RBTree) rightRotate(node *Node) *Node {
	x := node.left
	node.left = x.right
	x.right = node
	x.color = node.color
	node.color = RED
	return x
}

//filipColors
func (r *RBTree) filipColors(node *Node) {
	node.color = RED
	node.left.color = BLACK
	node.right.color = BLACK
}
func (r RBTree) isRed(node *Node) bool {
	if node == nil {
		return BLACK
	}
	return node.color == RED
}

//Add add one node
func (r *RBTree) Add(key, value interface{}) {
	r.root = r.add(r.root, key, value)
	//root always Black
	r.root.color = BLACK
}
func (r *RBTree) add(node *Node, key, value interface{}) *Node {
	if node == nil {
		r.size++
		return NewNode(key, value)
	}
	if r.less(key, node.key) {
		node.left = r.add(node.left, key, value)
	} else if r.more(key, node.key) {
		node.right = r.add(node.right, key, value)
	} else if r.equal(key, node.key) { //equal
		node.value = value
	}

	if r.isRed(node.right) && !r.isRed(node.left) {
		node = r.leftRotate(node)
	}
	if r.isRed(node.left) && r.isRed(node.left.left) {
		node = r.rightRotate(node)
	}
	if r.isRed(node.left) && r.isRed(node.right) {
		r.filipColors(node)
	}
	return node
}

//Contains check
func (r RBTree) Contains(key interface{}) bool {
	return r.getNode(r.root, key) != nil
}
func (r RBTree) getNode(node *Node, key interface{}) *Node {
	if node == nil {
		return nil
	}
	if r.equal(key, node.key) {
		return node
	} else if r.less(key, node.key) {
		return r.getNode(node.left, key)
	} else {
		return r.getNode(node.right, key)
	}
}

//Get one key
func (r *RBTree) Get(key interface{}) interface{} {
	node := r.getNode(r.root, key)
	if node == nil {
		return nil
	}
	return node.value
}

//Set update value
func (r *RBTree) Set(key, value interface{}) bool {
	node := r.getNode(r.root, key)
	if node == nil {
		return false
	}
	node.value = value
	return true
}

//Mini get
func (r RBTree) Mini() *Node {
	if r.size == 0 {
		return nil
	}
	return r.mini(r.root)
}
func (r RBTree) mini(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return r.mini(node.left)
}

//Max get
func (r RBTree) Max() *Node {
	if r.size == 0 {
		return nil
	}
	return r.max(r.root)
}
func (r RBTree) max(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return r.max(node.right)
}

//RmoveMin remove min node
func (r *RBTree) RmoveMin() *Node {
	ret := r.Mini()
	r.root = r.removeMin(ret)
	return ret
}
func (r *RBTree) removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		r.size--
		return rightNode
	}
	node.left = r.removeMin(node.left)
	return node
}

//RmoveMax return node
func (r *RBTree) RmoveMax() *Node {
	ret := r.Max()
	r.root = r.removeMax(ret)
	return ret
}
func (r *RBTree) removeMax(node *Node) *Node {
	if node.right == nil {
		leftNode := node.left
		node.left = nil
		r.size--
		return leftNode
	}
	node.right = r.removeMax(node.right)
	return node
}

//Remove remove key
func (r *RBTree) Remove(key interface{}) interface{} {
	node := r.getNode(r.root, key)
	if node != nil {
		r.root = r.remove(r.root, key)
		return node.value
	}
	return nil
}
func (r *RBTree) remove(node *Node, key interface{}) *Node {
	if node == nil {
		return nil
	}
	if r.less(key, node.key) {
		node.left = r.remove(node.left, key)
		return node
	} else if r.more(key, node.key) {
		node.right = r.remove(node.right, key)
		return node
	} else {
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			r.size--
			return rightNode
		}
		if node.right == nil {
			leftNode := node.right
			node.left = nil
			r.size--
			return leftNode
		}
		//left and right node all not nil
		successor := r.mini(node.right)
		successor.right = r.removeMin(node.right)
		successor.left = node.left
		node.left = nil
		node.right = nil
		return successor
	}
}
func (r RBTree) Prints() {
	r.prints(r.root)
}
func (r RBTree) prints(node *Node) {
	if node == nil {
		return
	}
	r.prints(node.left)
	fmt.Printf("%d ", node.value)
	r.prints(node.right)
}
