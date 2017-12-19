package main

import "log"
import "flag"
import "math/rand"
import "github.com/dadleyy/practice/trees/iterator"

func main() {
	options := struct {
		size int
	}{}

	flag.IntVar(&options.size, "size", 20, "the amount of elements in the tree")
	flag.Parse()

	head := &iterator.Node{Value: rand.Int() % options.size}
	top := head

	seen := map[int]bool{
		head.Value: true,
	}

	log.Printf("start: %v", head.Value)

	for i := 0; i < options.size-1; i++ {
		v := head.Value
		_, ok := seen[v]

		for ok {
			v = rand.Int() % options.size
			_, ok = seen[v]
		}

		seen[v] = true

		log.Printf("adding %v", v)

		if v > head.Value {
			head.Right = &iterator.Node{Value: v}
			head = head.Right
		} else {
			head.Left = &iterator.Node{Value: v}
			head = head.Left
		}
	}

	log.Printf("breadth first iterator")
	bfi := iterator.NewBreadthFirstIterator(top)

	for bfi.Done() == false {
		n := bfi.Next()
		log.Printf("found %v", n)
	}
}
