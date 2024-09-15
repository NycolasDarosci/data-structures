package stackandqueue

import "fmt"

/*
	QUEUE structure

	FIFO - "First In, First Out" -The first item added is the first item
	to be removed

	Some real examples: perfect tool for handling asynchronous requests—they
	ensure that the requests are processed in the order in which they were
	received. They are also commonly used to model real-world scenarios where
	events need to occur in a certain order, such as airplanes waiting for takeoff
	and patients waiting for their doctor.
*/
type queue[V any] struct {
	data []V
}

func NewQueue[V any]() queue[V] {
	return queue[V]{data: make([]V, 0)}
}

func (s *queue[V]) Enqueue(element V) {
	s.data = append(s.data, element)
}

func (s *queue[V]) Dequeue() {
	if s.IsEmpty() {
		return
	}
	s.data = s.data[1:len(s.data)]
}

func (s *queue[V]) Read() (*V, error) {
	if s.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	return &s.data[0], nil
}

func (s *queue[V]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *queue[V]) Print() {
	for _, value := range s.data {
		fmt.Println(value)
	}
}
