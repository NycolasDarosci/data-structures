package hashtables

/*
	This approach method is to increase significantly
	the efficiency of an algorithm: verify if a value of a sub array
	is contained into a larger array. Like set theory.

	This method will replace the traditional array data structure
	to hash table structure and using it as an "index" hash table to
	reduce to just one step algorithm

	Remember: hash table is a list of pair values -> key: value
*/

func IsSubset(arr1 []int, arr2 []int) bool {
	var smallerArray []int
	var biggerArray []int

	if len(arr1) > len(arr2) {
		biggerArray = arr1
		smallerArray = arr2
	} else {
		biggerArray = arr2
		smallerArray = arr1
	}

	for i := range smallerArray {
		isSubset := false
		for j := range biggerArray {
			if smallerArray[i] == biggerArray[j] {
				isSubset = true
				break
			}
		}
		if !isSubset {
			return false
		}
	}
	return true
}

func IsSubsetWithHashTable(arr1 []int, arr2 []int) bool {
	var smallerArray []int
	var biggerArray []int

	if len(arr1) > len(arr2) {
		biggerArray = arr1
		smallerArray = arr2
	} else {
		biggerArray = arr2
		smallerArray = arr1
	}

	hashTable := map[int]bool{}

	for _, value := range biggerArray {
		hashTable[value] = true
	}

	for _, value := range smallerArray {
		if !hashTable[value] {
			return false
		}
	}

	return true
}
