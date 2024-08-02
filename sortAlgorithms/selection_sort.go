package sortalgorithms

import "fmt"

/*
	__Selection sort -> O(n²)

	selection sort is O(n²) because it takes two steps to perform:
	COMPARISONS AND SWAPS

	This algorithm make a pass-through to the whole array, and
	take the first index as the lowest value and the algorithm
	effects a second pass-through to the array to compare each
	value with the lowest value previouslly accept
	if it is lowest than the previouslly accept, the lowest value is changed
	if it is changed the value, the swap is made with the lowest index
	to actual index

*/
func DoSelectionSort() {
	arr := []int{2, 3, 1, 6, 22, 21, 98}
	selectionSort(arr)
	fmt.Println(arr)
}

func selectionSort(arr []int) []int {
	for i := range len(arr) - 1 {
		currentlyLowestIndex := i
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[j] < arr[currentlyLowestIndex] {
				currentlyLowestIndex = j
			}
		}

		if currentlyLowestIndex != i {
			tmp := arr[i]
			arr[i] = arr[currentlyLowestIndex]
			arr[currentlyLowestIndex] = tmp
		}
	}
	return arr
}
