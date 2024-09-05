package stackandqueue

import "fmt"

// implement of a Stack structure

type Stack struct {
	data []int
}

func Initialize() Stack {
	return Stack{data: []int{}}
}

func (s *Stack) Push(element int) {
	s.data = append(s.data, element)
}

func (s *Stack) Pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *Stack) Read() int {
	return s.data[len(s.data)-1]
}

func (s *Stack) Print() {
	for _, value := range s.data {
		fmt.Println(value)
	}
}
