package list

import (
	_ "fmt"
	"reflect"
	"testing"
)

func TestList_New(t *testing.T) {

	s := NewList()

	if s.head != s.tail.prev {
		t.Error("Head != Tail.Prev")
	}
	if s.tail != s.head.next {
		t.Error("Tail != Head.Next")
	}

	if s.Len() != 0 {
		t.Error("Error len size expected 0, actual", s.Len())
	}
}

func TestList_PushBack(t *testing.T) {
	s := NewList()

	m := map[int]interface{}{
		0: "123",
		1: 53,
		2: -1231,
		3: "  ",
		4: 0.123123,
	}

	for i := 0; i < len(m); i++ {
		s.PushBack(m[i])
	}

	i := 0
	for node := s.head.next; node != s.tail; node = node.next {
		if node.data != m[i] {
			t.Error("Actual:", reflect.TypeOf(node.data), node.data, "\nExcpected:", reflect.TypeOf(m[i]), m[i])
		}
		i++
	}

	if s.Len() != len(m) {
		t.Error("Error len size expected", len(m), ", actual", s.Len())
	}
}

func TestList_PushFront(t *testing.T) {
	s := NewList()

	m := map[int]interface{}{
		0: "123",
		1: 53,
		2: -1231,
		3: "  ",
		4: 0.123123,
	}

	for i := 0; i < len(m); i++ {
		s.PushFront(m[i])
	}

	i := len(m) - 1
	for node := s.head.next; node != s.tail; node = node.next {
		if node.data != m[i] {
			t.Error("Actual:", reflect.TypeOf(node.data), node.data, "\nExcpected:", reflect.TypeOf(m[i]), m[i])
		}
		i--
	}
}
