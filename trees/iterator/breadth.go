package iterator

// NewBreadthFirstIterator returns an iterator that will traverse down each generation of tree's heirarchy.
func NewBreadthFirstIterator(node *Node) Iterator {
	iter := &breadth{
		start: node,
		level: make([]*Node, 1),
		next:  make([]*Node, 0, 2),
	}

	iter.level[0] = node

	iter.reload()

	return iter
}

type breadth struct {
	start *Node
	level []*Node
	next  []*Node
}

func (b *breadth) Next() *Node {
	if len(b.level) == 0 {
		return nil
	}

	top := b.level[0]

	b.level = b.level[1:]

	// we've emptied out the level, use next and reload
	if len(b.level) == 0 {
		b.level = b.next
		b.reload()
	}

	return top
}

func (b *breadth) reload() {
	b.next = make([]*Node, 0, len(b.level)*2)

	for _, n := range b.level {
		if n.Left != nil {
			b.next = append(b.next, n.Left)
		}

		if n.Right != nil {
			b.next = append(b.next, n.Right)
		}
	}
}

func (b *breadth) Done() bool {
	return len(b.level) == 0
}
