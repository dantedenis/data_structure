package list

type IList interface {
	PushBack(interface{})
	PushFront(interface{})
	PopBack() interface{}
	PopFront() interface{}
	Front() interface{}
	Back() interface{}
	//Swap
}

type List struct {
	size int
	head *node
	tail *node
}

type node struct {
	data interface{}
	prev *node
	next *node
}

func NewList() *List {
	head := &node{}
	tail := &node{}
	head.next = tail
	tail.prev = head

	return &List{
		size: 0,
		head: head,
		tail: tail,
	}
}

func (l *List) Len() int {
	return l.size
}

func (l *List) PushBack(data interface{}) {
	curr := &node{
		data: data,
		prev: l.tail.prev,
		next: l.tail,
	}

	l.tail.prev.next = curr
	l.tail.prev = curr
	l.size++
}

func (l *List) PushFront(data interface{}) {
	curr := &node{
		data: data,
		prev: l.head,
		next: l.head.next,
	}

	l.head.next = curr
	l.size++
}