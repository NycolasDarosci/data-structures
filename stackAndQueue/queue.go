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
type queue struct {
	data []int
}

func NewQueue() queue {
	return queue{data: make([]int, 0)}
}

func (s *queue) Enqueue(element int) {
	s.data = append(s.data, element)
}

func (s *queue) Dequeue() {
	if s.IsEmpty() {
		return
	}
	s.data = s.data[1:len(s.data)]
}

func (s *queue) Read() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.data[0], nil
}

func (s *queue) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *queue) Print() {
	for _, value := range s.data {
		fmt.Println(value)
	}
}
