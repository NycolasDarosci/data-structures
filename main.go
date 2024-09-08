package main

import (
	h "data-structure/hashTables"
	s "data-structure/stackAndQueue"
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
	duplicateValue := h.FirstDuplicateValues([]string{"a", "b", "a", "h", "c", "i"})

	// phrase contains all alphabet letters, except 'f'
	phrase := "the quick brown box jumps over a lazy dog"
	exceptLetter := h.ContainsAllAlphabetLetters(phrase)

	nonDuplicated := h.FirstNonDuplicated("minimum")

	fmt.Println(isContained)
	fmt.Println(isContained2)
	fmt.Println(intersection)   // 2, 8, 5
	fmt.Println(duplicateValue) // a
	fmt.Println(exceptLetter)   // f
	fmt.Println(nonDuplicated)  // n -> first character non duplicated found

	//a := combineTwoCharacters([]string{"a", "b", "c", "d"})
	//fmt.Println(a)

	// Stack
	stack := s.New()
	stack.Push(1)
	stack.Push(3)
	stack.Push(5)
	stack.Push(2)

	stack.Pop()

	stack.Print()

	// lastElement := stack.Read()
	// fmt.Println(lastElement)
}
