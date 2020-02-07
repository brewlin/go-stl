package bm

//BMNode b- node
type BMNode struct {
	parent *BMNode
	size   int
	key    []int
	child  []*BMNode
}

func NewBMNode(m int) *BMNode {
	return &BMNode{
		parent: nil,
		size:   0,
		key:    make([]int, m+1),
		child:  make([]*BMNode, m+1),
	}
}

func (b *BMNode) split(m int) *BMNode {
	node := NewBMNode(m)
	parent := b.parent
	if parent == nil {
		parent = NewBMNode(m)
	}
	mid := b.size/2 + 1
	node.size = m - mid
	b.size = mid - 1
	j := 1
	k := mid + 1
	//新生成的右节点
	for ; k <= m; k++ {
		node.key[j] = b.key[k]
		node.child[j-1] = b.child[k-1]
		j += 1
	}
	node.child[j-1] = b.child[k-1]
	node.parent = parent
	b.parent = parent
	//将该节点中间节点插入到父节点
	k = parent.size
	for ; b.key[mid] < parent.key[k]; k-- {
		parent.key[k+1] = parent.key[k]
		parent.child[k+1] = parent.child[k]
	}
	parent.key[k+1] = b.key[mid]
	parent.child[k] = b
	parent.child[k+1] = node
	parent.size++

	if parent.size >= m {
		return parent.split(m)
	}
	return parent

}
func (b *BMNode) deleteNode(v, i, m int) *BMNode {
	if b.child[i] != nil {
		vt, nt := b.findAfterMinNode(i)
		b.key[i] = *vt
		nt.deleteNode(*vt, 1, m)
	} else {
		for k := i; k < b.size; k++ {
			b.key[k] = b.key[k+1]
			b.child[k] = b.child[k+1]
		}
		b.size--
		if b.size < (m-1)/2 && b.parent != nil {
			ok, b := b.restore(m)
			if !ok {
				b = b.mergeNode(m)
			}
		}
	}
	for b.parent != nil {
		b = b.parent
	}
	return b
}
func (b *BMNode) restore(m int) (bool, *BMNode) {
	p := b.parent
	j := 0
	for ; p.child[j] != b; j++ {
	}
	if j > 0 {
		left := p.child[j-1]
		if left.size > (m-1)/2 {
			for k := b.size; k >= 0; k-- {
				b.key[k+1] = b.key[k]
			}
			b.key[1] = p.key[j]
			p.key[j] = left.key[left.size]
			b.size++
			left.size--
			return true, p
		}
	}
	if j < p.size {
		right := p.child[j+1]
		if right.size > (m-1)/2 {
			b.key[b.size+1] = p.key[j+1]
			p.key[j+1] = right.key[1]
			for k := 1; k < right.size; k++ {
				right.key[k] = right.key[k+1]
			}
			b.size++
			right.size--
			return true, p
		}
	}
	return false, b
}
func (b *BMNode) mergeNode(m int) *BMNode {
	node := NewBMNode(m)
	p := b.parent

	j := 0
	for ; p.child[j] != b; j++ {
	}
	if j > 0 {
		node = p.child[j-1]
		node.size++
		node.key[node.size] = p.key[j]
		for k := 1; k <= b.size; k++ {
			node.size++
			node.key[node.size] = b.key[k]
		}
		p.size--
		for k := j; k < p.size; k++ {
			p.key[k] = p.key[k+1]
			p.child[k] = p.child[k+1]
		}
	} else {
		node = p.child[j+1]
		b.size++
		b.key[b.size] = p.key[j]
		for k := 1; k <= node.size; k++ {
			b.size++
			b.key[b.size] = b.key[k]
		}
		p.size--
		for k := j; k <= p.size; k++ {
			p.key[k] = p.key[k+1]
			p.child[k] = p.child[k+1]
		}
	}
	return p
}
func (b *BMNode) findAfterMinNode(i int) (*int, *BMNode) {
	leaf := b
	if leaf == nil {
		return nil, nil
	} else {
		leaf = leaf.child[i]
		for leaf.child[0] != nil {
			leaf = leaf.child[0]
		}
	}
	return &leaf.key[1], leaf
}
