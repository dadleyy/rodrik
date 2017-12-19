package algorithms

func Quick(input []int) []int {
	swapped := true
	size := len(input)
	output := make([]int, size)
	copy(output, input)

	for swapped {
		swapped = false

		for i := 0; i < size; i++ {
			for j := i + 1; j < size; j++ {

				if output[j] < output[i] {
					output[i], output[j] = output[j], output[i]
					swapped = true
					break
				}
			}
		}
	}

	return output
}
