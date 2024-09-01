package main

import (
	h "data-structure/hashTables"
	"fmt"
)

func main() {
	//DoLinearSearch()
	//DoBinarySearch()
	//s.DoBubbleSort()
	//s.DoSelectionSort()
	//s.DoInsertionSort()

	arr1 := []int{1, 2, 5, 6, 7, 4, 8}
	arr2 := []int{2, 8, 5, 3}

	isContained := h.IsSubset(arr1, arr2)
	isContained2 := h.IsSubsetWithHashTable(arr1, arr2)
	intersection := h.IntersectionOfTwo(arr1, arr2)

	fmt.Println(isContained)
	fmt.Println(isContained2)
	fmt.Println(intersection)

	//a := combineTwoCharacters([]string{"a", "b", "c", "d"})
	//fmt.Println(a)
}
