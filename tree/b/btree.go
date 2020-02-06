package b

const M int = 4

const Min int = M/2 - 1

type BT struct {
	parent *BT
	keynum int
	key    [M + 1]int
	child  [M + 1]*BT
}

func (b *BT) Search(v int) (bool, int, *BT) {
	node := &BT{}
	var i int
	if b == nil {
		return false, 0, nil
	}
	for b != nil {
		for i = b.keynum; i > 0 && v <= b.key[i]; i-- {
			if v == b.key[i] {
				return true, i, b
			}
		}
		if b.child[i] == nil {
			node = b
		}
		b = b.child[i]
	}
	return false, i, node
}

func (b *BT) Split() *BT {
	node := BT{}
	parent := b.parent
	if parent == nil {
		parent = &BT{}
	}
	mid := b.keynum/2 + 1
	node.keynum = M - mid
	b.keynum = mid - 1
}
