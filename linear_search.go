package main

import "fmt"

/*
	__Linear Search
	check each cell one at a time — from left to right -
	until find the value

	__Linear Search on an ordered array
	check each cell one at a time — from left to right -
	until find the value, but the value is compared with the
	values within the array.
		Is value input greater than the value? search for next
		Is value input greater than the value? no
		stop the search

	Search -> find for the index given a value
	Read -> find for the value given a index
*/

// linear search on an ordered array
func DoLinearSearch() {
	arr := []int{1, 3, 7, 15, 22, 78}
	searchValue := 1 // 2
	value := linearSearch(arr, searchValue)
	fmt.Println(*value)
}

func linearSearch(arr []int, searchValue int) *int {
	for index, value := range arr {
		if value == searchValue {
			return &index
		} else if value > searchValue {
			break
		}
	}
	return nil
}
