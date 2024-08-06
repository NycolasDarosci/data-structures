package sortalgorithms

import "fmt"

/*
	Insertion Sort -> N(O²)
*/
func DoInsertionSort() {
	arr := []int{4, 3, 1, 7, 6}
	insertionSort(arr)
	fmt.Println(arr)
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		tempValue := arr[i]
		previous := i - 1

		for previous >= 0 {
			if arr[previous] > tempValue {
				arr[previous+1] = arr[previous]
				previous--
			} else {
				break
			}
		}
		arr[previous+1] = tempValue
	}
	return arr
}
