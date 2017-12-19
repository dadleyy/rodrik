package algorithms

import "testing"
import "github.com/franela/goblin"

func Test_MergeSort(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("merge sort", func() {
		g.It("successfully sorts the input", func() {
			a := []int{10, 50, 2, 4, 1, 30}
			e := []int{1, 2, 4, 10, 30, 50}
			m := Merge(a)

			g.Assert(len(a)).Equal(len(m))

			for i, _ := range a {
				g.Assert(e[i]).Equal(m[i])
			}
		})
	})
}

func Benchmark_Merge(b *testing.B) {
	a := []int{10, 50, 2, 4, 1, 30}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Merge(a)
	}
}
