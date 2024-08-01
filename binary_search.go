package main

import "fmt"

/*
	__Binary search

	eliminate a cells comparing its values with the input value.
	The algorithm do this getting the midpoint value and comparing
	if it is greater or less or equal than the input value.

	Binary search increase proportionally
	when the array doubles increase and binary search takes plus one step to perform
	O(log n)

*/
func DoBinarySearch() {
	arr := []int{1, 3, 7, 15, 22, 78, 99, 108}
	searchValue := 99 // 5
	value := binarySearch(arr, searchValue)
	fmt.Println(*value)
}

func binarySearch(arr []int, searchValue int) *int {
	lowerIndex := 0
	upperIndex := len(arr) - 1

	for lowerIndex <= upperIndex {
		middleIndex := (lowerIndex + upperIndex) / 2
		middleValue := arr[middleIndex]

		if middleValue == searchValue {
			return &middleIndex
		} else if middleValue > searchValue {
			upperIndex = middleIndex - 1
		} else if middleValue < searchValue {
			lowerIndex = middleIndex + 1
		}
	}
	return nil
}
