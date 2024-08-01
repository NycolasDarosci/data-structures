package main

import s "data-structure/sortalgorithms"

func main() {
	//DoLinearSearch()
	//DoBinarySearch()
	s.DoBubbleSort()

}

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

func greatestNumberRefactor(arr []int) int {
	greatestNumber := arr[0]
	for i := range arr {
		if arr[i] > greatestNumber {
			greatestNumber = arr[i]
		}
	}
	return greatestNumber
}
