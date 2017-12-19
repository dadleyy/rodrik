package iterator

// Iterator represents a basic tree traversal iterator.
type Iterator interface {
	Next() *Node
	Done() bool
}
