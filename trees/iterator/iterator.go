package iterator

type Iterator interface {
	Next() *Node
	Done() bool
}
