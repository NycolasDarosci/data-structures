package hashtables

/*
	The purpose of these exercises is to make O(N) algorithms
	when it are most common on O(N * M) or O(N²) efficiencies
*/

// efficiency of O(N)
func IntersectionOfTwo(arr1 []int, arr2 []int) []int {
	hashTable := make(map[int]bool)

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
	strings := make(map[string]bool)
	for i, value := range arr {
		if strings[arr[i]] {
			return value
		} else {
			strings[value] = true
		}
	}
	return ""
}

// it will return the first letter missing on the word string accordingly the alphabet
func ContainsAllAlphabetLetters(word string) string {
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	hashTable := make(map[string]bool)

	if len(word) == 0 {
		return ""
	}

	for i := range word {
		// i is a int
		// word[i] return byte. string(word[a]) return the letter
		letter := string(word[i])
		hashTable[letter] = true
	}

	for i := range alphabet {
		letter := string(alphabet[i])
		if !hashTable[letter] {
			return letter
		}
	}
	return ""
}

//returns the first non-duplicated character in a string
func FirstNonDuplicated(str string) string {
	hashTable := make(map[string]int)
	for i := range str {
		character := string(str[i])
		if hashTable[character] >= 1 {
			hashTable[character]++
		} else {
			hashTable[character] = 1
		}
	}

	for j := range str {
		character := string(str[j])
		if hashTable[character] == 1 {
			return character
		}
	}
	return ""
}
