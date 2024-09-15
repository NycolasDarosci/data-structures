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
	stack := s.NewStack[int]()
	stack.Push(2)
	stack.Push(4)
	stack.Push(6)
	stack.Print() // 2, 4, 6

	stack.Pop()
	stack.Print() // 2, 4

	// Queue
	q := s.NewQueue[int]()
	q.Enqueue(1)
	q.Enqueue(3)
	q.Enqueue(5)
	q.Print() // 1, 3, 5

	q.Dequeue()
	q.Print() // 3, 5

	i, err := q.Read()
	if err == nil {
		fmt.Println("ultimo: ", *i)
	}

	// lastElement := stack.Read()
	// fmt.Println(lastElement)
}
