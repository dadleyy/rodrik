package iterator

import "fmt"

// Node represents a single integer point on a binary search tree.
type Node struct {
	Value int   `json:"value"`
	Left  *Node `json:"left"`
	Right *Node `json:"right"`
}

// String prints out the value of the node w/ some formatting.
func (n *Node) String() string {
	return fmt.Sprintf("value[%d]", n.Value)
}
