package iterator

func NewBreadthFirstIterator(node *Node) Iterator {
	iter := &breadth{
		start: node,
		level: make([]*Node, 1),
		next:  make([]*Node, 0, 2),
	}

	iter.level[0] = start

	if iter.start.Right != nil {
		iter.next = append(iter.start.Right)
	}

	if iter.start.Left != nil {
		iter.next = append(iter.start.Left)
	}

	return iter
}

type breadth struct {
	start *Node
	level []*Node
	next  []*Node
}

func (b *breadth) Next() *Node {
	return nil
}

func (b *breadth) Done() bool {
	return true
}
