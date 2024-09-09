package stackandqueue

import "fmt"

//
/*
	STACK structure

	LIFO - “Last In, First Out.” The last item pushed onto a
	stack is always the first item popped from it.

	Traditionals arrays work directly with the computer's memory
	Stack ises and a set of rul procer goal and
	it written on top of othesses of how we interact with array
	to achieve a particular built in data structures

	stack is a base for RECURSION
*/

type stack struct {
	data []int
}

func NewStack() stack {
	return stack{data: make([]int, 0)}
}

func (s *stack) Push(element int) {
	s.data = append(s.data, element)
}

func (s *stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func (s *stack) Read() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.data[len(s.data)-1], nil
}

func (s *stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *stack) Print() {
	for _, value := range s.data {
		fmt.Println(value)
	}
}

// stack is used on linter - opening and closing braces - algorithm
// here is a example implementation
func (s *stack) Linter(text string) {

}
