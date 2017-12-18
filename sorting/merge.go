package main

import "log"
import "math/rand"

func merge(left, right []int) []int {
	out := make([]int, 0, len(left)+len(right))

	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(out, right...)
		}
		if len(right) == 0 {
			return append(out, left...)
		}

		if left[0] <= right[0] {
			out = append(out, left[0])
			left = left[1:]
		} else {
			out = append(out, right[0])
			right = right[1:]
		}
	}

	return out
}

func sort(input []int) []int {
	length := len(input)

	if length <= 1 {
		return input
	}

	mid := length / 2
	log.Printf("splits: %v | %v", input[:mid], input[mid:])

	left := sort(input[:mid])
	right := sort(input[mid:])

	return merge(left, right)
}

func main() {
	input := make([]int, 20)

	for i := 0; i < len(input); i++ {
		input[i] = rand.Int() % 20
	}

	log.Printf("input: %v", input)
	log.Printf("sorted: %v", sort(input))
}
