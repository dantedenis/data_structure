package stack

import "fmt"

type IStack interface {
	Pop() interface{}
	Push(interface{})
}

type Stack struct {
	head *node
	size int
}

type node struct {
	value interface{}
	next  *node
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Pop() (data interface{}) {
	if s.size > 0 {
		data, s.head = s.head.value, s.head.next
		s.size--
		return
	}
	return nil
}

func (s *Stack) Push(data interface{}) {
	s.head = &node{data, s.head}
	s.size++
}

func (s *Stack) Len() int {
	return s.size
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) String() (result string) {
	temp := s.head
	space := ""
	for temp != nil {
		result += fmt.Sprint(space, temp.value)
		temp = temp.next
		space = " "
	}
	return
}
