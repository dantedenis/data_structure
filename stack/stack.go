package stack

import "fmt"

type Stack interface {
	fmt.Stringer

	Pop() any
	Push(any)
	Len() int
	IsEmpty() bool
}

type stack struct {
	head *node
	size int
}

type node struct {
	value any
	next  *node
}

func NewStack() Stack {
	return &stack{}
}

func (s *stack) Pop() (data any) {
	if s.size > 0 {
		data, s.head = s.head.value, s.head.next
		s.size--
		return
	}
	return nil
}

func (s *stack) Push(data any) {
	s.head = &node{data, s.head}
	s.size++
}

func (s *stack) Len() int {
	return s.size
}

func (s *stack) IsEmpty() bool {
	return s.size == 0
}

func (s *stack) String() (result string) {
	temp := s.head
	space := ""
	for temp != nil {
		result += fmt.Sprint(space, temp.value)
		temp = temp.next
		space = " "
	}
	return
}
