package main

type Node struct {
	Data int
	Next *Node
}

type List struct {
	size int
	head *Node
	tail *Node
}

func New() *List {
	return &List{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (l *List) Append(node *Node) bool {
	if node == nil {
		return false
	}

	if l.size == 0 {
		l.head = node
	} else {
		oldTail := l.tail
		oldTail.Next = node
	}

	l.tail = node
	l.size++

	return true
}
