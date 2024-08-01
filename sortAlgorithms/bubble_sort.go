package sortalgorithms

import "fmt"

func DoBubbleSort() {
	arr := []int{5, 2, 76, 34, 22, 108}
	bubbleSort(arr)
	fmt.Println(arr)
}

func bubbleSort(arr []int) []int {
	unsortedUntilIndex := len(arr) - 1
	sorted := false

	// each round of this loop represents a pass-through of the array
	for !sorted {
		sorted = true
		for index := range unsortedUntilIndex {
			current := arr[index]
			next := arr[index+1]
			if current > next {
				arr[index+1] = current
				arr[index] = next
				sorted = false
			}
		}
		unsortedUntilIndex -= 1
	}
	return arr
}
