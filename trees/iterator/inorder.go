package iterator

// NewInorderIterator returns an in-order depth first traversal.
func NewInorderIterator(node *Node) Iterator {
	iter := &inorder{
		trace: []*Node{node},
	}

	iter.trace = append(iter.trace, iter.descendants(node)...)

	iter.head = iter.trace[len(iter.trace)-1]
	iter.trace = iter.trace[:len(iter.trace)-1]

	return iter
}

type inorder struct {
	head  *Node
	trace []*Node
}

func (i *inorder) Done() bool {
	return len(i.trace) == 0 && i.head == nil
}

func (i *inorder) Next() *Node {
	h := i.head

	if h.Right != nil {
		i.trace = append(i.trace, h.Right)
		i.trace = append(i.trace, i.descendants(h.Right)...)
	}

	next := len(i.trace) - 1

	if next < 0 {
		i.head = nil
		return h
	}

	i.head = i.trace[next]
	i.trace = i.trace[:next]

	return h
}

func (i *inorder) descendants(n *Node) []*Node {
	out := make([]*Node, 0)

	for n.Left != nil {
		out = append(out, n.Left)
		n = n.Left
	}

	return out
}
