package iterator

import "fmt"
import "testing"
import "strings"
import "github.com/franela/goblin"

func Test_BreadthIterator(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("breadth-first iterator", func() {
		g.It("sucessfully iterates a tree along each hierarchical level", func() {
			tree := treeFromJSON(example)
			results := make([]string, 0, 8)
			i := NewBreadthFirstIterator(tree)

			for i.Done() == false {
				results = append(results, fmt.Sprintf("%d", i.Next().Value))
			}

			g.Assert(strings.Join(results, ",")).Equal("10,6,20,2,8,18,17")
		})
	})
}
