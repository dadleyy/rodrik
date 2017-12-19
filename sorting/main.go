package main

import "log"
import "math/rand"
import "github.com/dadleyy/rodrik/sorting/algorithms"

func main() {
	input := make([]int, 20)

	for i := 0; i < len(input); i++ {
		input[i] = rand.Int() % 20
	}

	log.Printf("input: %v", input)
	log.Printf("sorted: %v", algorithms.Merge(input))
}
