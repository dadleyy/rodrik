package main

import "log"
import "fmt"
import "flag"
import "sort"
import "bytes"
import "math/rand"
import "github.com/dadleyy/rodrik/trees/iterator"

func tree(pool []int) *iterator.Node {
	size := len(pool)

	if size == 0 {
		return nil
	}

	if size == 1 {
		v := &iterator.Node{Value: pool[0]}
		return v
	}

	mid := size / 2
	v := &iterator.Node{Value: pool[mid]}

	v.Left = tree(pool[:mid])

	if mid < size {
		v.Right = tree(pool[mid+1:])
	}

	return v
}

func main() {
	options := struct {
		size int
	}{}

	flag.IntVar(&options.size, "size", 20, "the amount of elements in the tree")
	flag.Parse()

	seen, pool := make(map[int]bool), make([]int, options.size)

	for i := 0; i < options.size; i++ {
		v := rand.Int() % options.size
		_, ok := seen[v]

		for ok {
			v = rand.Int() % options.size
			_, ok = seen[v]
		}

		seen[v] = true
		pool[i] = v
	}

	sort.Ints(pool)

	children, stack := make([]*iterator.Node, 0, options.size), make([]*iterator.Node, 1, 10)
	stack[0] = tree(pool)
	printer := new(bytes.Buffer)

	for len(stack) > 0 {
		if stack[0].Left != nil {
			children = append(children, stack[0].Left)
		}

		fmt.Fprintf(printer, " |%v| ", stack[0].Value)

		if stack[0].Right != nil {
			children = append(children, stack[0].Right)
		}

		// pop off the next node
		stack = stack[1:]

		// if we're out of items but we had children, move to next level
		if len(stack) == 0 {
			stack = children
			children = make([]*iterator.Node, 0)
			log.Printf("level: %v", printer.String())
			printer.Reset()
		}
	}

	log.Printf("iterating breadth-first:")
	iter := iterator.NewBreadthFirstIterator(tree(pool))
	for iter.Done() == false {
		log.Printf("next: %v", iter.Next().Value)
	}

	log.Printf("iterating depth-first inorder:")
	iter = iterator.NewInorderIterator(tree(pool))
	for iter.Done() == false {
		log.Printf("next: %v", iter.Next().Value)
	}
}
