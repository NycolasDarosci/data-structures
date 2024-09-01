package hashtables

/*
	The purpose of these exercises is to make O(N) algorithms
	when it are most common on O(N * M) or O(N²) efficiencies
*/

// efficiency of O(N)
func IntersectionOfTwo(arr1 []int, arr2 []int) []int {
	hashTable := map[int]bool{}

	for _, value := range arr1 {
		hashTable[value] = true
	}

	intersection := []int{}
	for _, value := range arr2 {
		if hashTable[value] {
			intersection = append(intersection, value)
		}
	}
	return intersection
}

func FirstDuplicateValues(arr []string) string {
	strings := map[string]bool{}
	for _, value := range arr {
		strings[value] = true
	}

	// continue
	for _, value := range arr {
		if strings[value] {
			return value
		}
	}
	return ""
}
