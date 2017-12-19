package iterator

import "fmt"
import "testing"
import "strings"
import "github.com/franela/goblin"

func Test_InorderIterator(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("depth-first inorder iterator", func() {
		g.It("sucessfully iterates a tree in-order and terminates when complete", func() {
			tree := treeFromJSON(example)
			results := make([]string, 0, 8)
			i := NewInorderIterator(tree)

			for i.Done() == false {
				results = append(results, fmt.Sprintf("%d", i.Next().Value))
			}

			g.Assert(strings.Join(results, ",")).Equal("2,6,8,10,17,18,20")
		})
	})
}
