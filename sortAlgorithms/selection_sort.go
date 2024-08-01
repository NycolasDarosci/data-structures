package sortalgorithms

import "fmt"

/*
	__Selection sort -> O(n²)

	make description..


*/
func DoSelectionSort() {
	arr := []int{2, 3, 1, 6}
	selectionSort(arr)
	fmt.Println(arr)
}

// implement algorithm
func selectionSort(arr []int) []int {
	firstIndex := 0
	lowestValue := arr[0]

	for firstIndex <= len(arr)-1 {
		for i := range firstIndex {
			if arr[i] < lowestValue {
				lowestValue = arr[i]
			}
		}
		arr[firstIndex] = lowestValue
		firstIndex++
	}
	return arr
}
