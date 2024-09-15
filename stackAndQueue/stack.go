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

type stack[V any] struct {
	data []V
}

func NewStack[V any]() stack[V] {
	return stack[V]{data: make([]V, 0)}
}

func (s *stack[V]) Push(element V) {
	s.data = append(s.data, element)
}

func (s *stack[V]) Pop() {
	if s.IsEmpty() {
		return
	}
	s.data = s.data[:len(s.data)-1]
}

func (s *stack[V]) Read() (*V, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	return &s.data[len(s.data)-1], nil
}

func (s *stack[V]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *stack[V]) Print() {
	for _, value := range s.data {
		fmt.Println(value)
	}
}

// stack is used on linter - opening and closing braces - algorithm
// here is a example implementation
func (s *stack[V]) Linter(text string) {

}
