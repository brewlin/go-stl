package bst

import "fmt"
import "stl/queue/list"

type Node struct {
	e     int
	left  *Node
	right *Node
}

//NewNode return node
func NewNode(e int) *Node {
	var node Node
	node.e = e
	node.left = nil
	node.right = nil
	return &node
}

//Bst struct
type Bst struct {
	root *Node
	size int
}

//NewBst return bst
func NewBst() *Bst {
	var bst Bst
	bst.root = nil
	bst.size = 0
	return &bst
}

//Size get size
func (b Bst) Size() int {
	return b.size
}

//IsEmpty bool
func (b Bst) IsEmpty() bool {
	return b.size == 0
}

//Add add
func (b *Bst) Add(e int) {
	if b.root == nil {
		b.root = NewNode(e)
		b.size++
	} else {
		b.add(b.root, e)
	}
}

//add node by recursive
func (b *Bst) add(node *Node, e int) {
	if node.e == e {
		return
	} else if e < node.e && node.left == nil {
		node.left = NewNode(e)
		b.size++
		return
	} else if e > node.e && node.right == nil {
		node.right = NewNode(e)
		b.size++
	}

	if e < node.e {
		b.add(node.left, e)
	} else {
		b.add(node.right, e)
	}

}
func (b Bst) Contains(e int) bool {
	return b.contains(b.root, e)
}
func (b Bst) contains(node *Node, e int) bool {
	if node == nil {
		return false
	}
	if e == node.e {
		return true
	} else if e < node.e {
		return b.contains(node.left, e)
	} else {
		return b.contains(node.right, e)
	}
}
func (b Bst) PreOrder() {
	b.preOrder(b.root)
}
func (b Bst) preOrder(node *Node) {
	if node == nil {
		return
	}
	b.preOrder(node.left)
	fmt.Println(node.e)
	b.preOrder(node.right)
}
func (b Bst) LevelOrder() {
	queue := list.NewQueue()
	queue.Push(b.root)
	for !queue.IsEmpty() {
		cur := queue.Pop().(*Node)
		fmt.Println(cur.e)
		if cur.left != nil {
			queue.Push(cur.left)
		}
		if cur.right != nil {
			queue.Push(cur.right)
		}

	}
}

//mini get mini value
func (c Bst) Mini() int {
	if c.size == 0 {
		panic("bst is empty")
	}
	return c.mini(c.root).e
}
func (c Bst) mini(node *Node) *Node {
	if node.left == nil {
		return node
	}
	return c.mini(node.left)
}
func (c Bst) Max() int {
	if c.size == 0 {
		panic("bst is empty")
	}
	return c.max(c.root).e
}
func (c Bst) max(node *Node) *Node {
	if node.right == nil {
		return node
	}
	return c.max(node.right)
}
func (c *Bst) RemoveMin() int {
	ret := c.Mini()
	c.root = c.removeMin(c.root)
	return ret
}
func (c *Bst) removeMin(node *Node) *Node {
	if node.left == nil {
		rightNode := node.right
		node.right = nil
		c.size--
		return rightNode
	}
	node.left = c.removeMin(node.left)
	return node
}
func (c *Bst) RemoveMax() int {
	ret := c.Max()
	c.root = c.removeMax(c.root)
	return ret
}
func (c *Bst) removeMax(node *Node) *Node {
	if node.right == nil {
		leftNode := node.left
		node.left = nil
		c.size--
		return leftNode
	}
	node.right = c.removeMax(node.right)
	return node
}
func (c *Bst) Remove(e int) {
	c.root = c.remove(c.root, e)
}
func (c *Bst) remove(node *Node, e int) *Node {
	if node == nil {
		return nil
	}
	if e < node.e {
		node.left = c.remove(node.left, e)
		return node
	} else if e > node.e {
		node.right = c.remove(node.right, e)
		return node
	} else {
		if node.left == nil {
			rightNode := node.right
			node.right = nil
			c.size--
			return rightNode
		}
		if node.right == nil {
			leftNode := node.left
			node.left = nil
			c.size--
			return leftNode
		}
		successor := c.mini(node.right)
		successor.right = c.removeMin(node.right)
		successor.left = node.left
		node.left = nil
		node.right = nil
		return successor
	}

}
