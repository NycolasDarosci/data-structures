package bigonotations

// pratical example
// a function that verify if there are duplicate values
// in two differents algorithms
// how substitute O(n²) to O(n) -> more efficiancy

/*
	O(n²) -> nested for loops
*/
func hasDuplicateValues(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

/*
	O(n)
*/
func hasDuplicateValuesRefactor(arr []int) bool {
	existingValues := []int{}
	for i := 0; i < len(arr); i++ {
		if existingValues[arr[i]] == 1 {
			return true
		} else {
			existingValues[arr[i]] = 1
		}
	}
	return false
}

// pratical example
// a function that verify which is the greatest number
// in two differents algorithms

/*
	O(n²)
*/
func greatestNumber(arr []int) *int {
	for i := range arr {
		isIValTheGreatest := true
		for j := range arr {
			if j > i {
				isIValTheGreatest = false
			}
		}
		if isIValTheGreatest {
			return &i
		}
	}
	return nil
}

/*
	O(n)
*/
func greatestNumberRefactor(arr []int) int {
	greatestNumber := arr[0]
	for i := range arr {
		if arr[i] > greatestNumber {
			greatestNumber = arr[i]
		}
	}
	return greatestNumber
}
