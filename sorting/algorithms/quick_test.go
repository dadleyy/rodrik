package algorithms

import "testing"
import "github.com/franela/goblin"

func Test_QuickSort(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("quick sort", func() {
		g.It("sorts the input correctly", func() {
			a := []int{10, 50, 4, 20, 1, 2, 100}
			e := []int{1, 2, 4, 10, 20, 50, 100}
			r := Quick(a)
			g.Assert(len(r)).Equal(len(a))

			for i, _ := range r {
				g.Assert(r[i]).Equal(e[i])
			}
		})
	})
}

func Benchmark_Quick(b *testing.B) {
	a := []int{10, 50, 2, 4, 1, 30}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Quick(a)
	}
}
