package algorithms

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

func Merge(input []int) []int {
	length := len(input)

	if length <= 1 {
		return input
	}

	mid := length / 2

	left := Merge(input[:mid])
	right := Merge(input[mid:])

	return merge(left, right)
}
