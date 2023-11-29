package list

import (
	"go-stl/datastruct/container"
	"go-stl/datastruct/node"
)

type List[T any] struct {
	head *node.LinkNode[T]
	tail *node.LinkNode[T]
	size int
}

func NewList[T any]() *List[T] {
	dummyHead := node.LinkNode[T]{}
	dummyTail := node.LinkNode[T]{}
	dummyHead.Next = &dummyTail
	dummyTail.Pre = &dummyHead
	return &List[T]{
		head: &dummyHead,
		tail: &dummyTail,
		size: 0,
	}
}

func (l *List[T]) Get(index int) T {
	return l.locate(index).Data
}

func (l *List[T]) Set(index int, x T) {
	l.locate(index).Data = x
}

func (l *List[T]) Insert(index int, x T) {
	newNode := node.NewLinkNode[T](x)
	cur := l.locate(index)
	newNode.Next = cur
	newNode.Pre = cur.Pre
	cur.Pre.Next = newNode
	cur.Pre = newNode
	l.size++
}

func (l *List[T]) Erase(index int) {
	cur := l.locate(index)
	cur.Pre.Next = cur.Next
	cur.Next.Pre = cur.Pre
	l.size--
}

func (l *List[T]) locate(index int) *node.LinkNode[T] {
	if index >= l.size || index < 0 {
		panic("index " + string(rune(index)) + " out of range " + string(rune(l.size)))
	}
	cur := l.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur
}

func (l *List[T]) Front() T {
	if l.size == 0 {
		panic("List is empty")
	}
	return l.head.Next.Data
}

func (l *List[T]) Back() T {
	if l.size == 0 {
		panic("List is empty")
	}
	return l.tail.Pre.Data
}

func (l *List[T]) PushBack(x T) {
	newNode := node.NewLinkNode[T](x)
	if l.head == nil || l.tail == nil {
		panic("List is nil")
	}
	l.tail.Pre.Next = newNode
	newNode.Pre = l.tail.Pre
	newNode.Next = l.tail
	l.tail.Pre = newNode
	l.size++
}

func (l *List[T]) PopBack() {
	if l.head == nil || l.tail == nil {
		panic("List is nil")
	}
	if l.size != 0 {
		l.tail.Pre.Pre.Next = l.tail
		l.tail.Pre = l.tail.Pre.Pre
		l.size--
	}
}

func (l *List[T]) PushFront(x T) {
	newNode := node.NewLinkNode[T](x)
	if l.head == nil || l.tail == nil {
		panic("List is nil")
	}
	l.head.Next.Pre = newNode
	newNode.Next = l.head.Next
	newNode.Pre = l.head
	l.head.Next = newNode
	l.size++
}

func (l *List[T]) PopFront() {
	if l.head == nil || l.tail == nil {
		panic("List is nil")
	}
	if l.size != 0 {
		target := l.head.Next
		l.head.Next.Next.Pre = l.head
		l.head.Next = l.head.Next.Next
		l.size--
		target.Next = nil
		target.Pre = nil
	}
}

func (l *List[T]) Empty() bool {
	return l.size == 0
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) Clear() {
	if l.size != 0 {
		l.head.Next.Pre = nil
		l.tail.Pre.Next = nil
		l.head.Next = l.tail
		l.tail.Pre = l.head
		l.size = 0
	}
}

func (l *List[T]) ForEach(consumer container.Consumer[int, T]) {
	cur := l.head
	for i := 0; i < l.size; i++ {
		cur = cur.Next
		if !consumer(i, cur.Data) {
			break
		}
	}
}

func (l *List[T]) Sort(comparator container.Compare[T]) {
	if l.size == 0 {
		return
	}
	quickSort(l.head.Next, l.tail.Pre, comparator)
}

func quickSort[T any](start, end *node.LinkNode[T], comparator container.Compare[T]) {
	if start == end || start == end.Next {
		return
	}
	pivot := start.Data
	i := start
	j := end
	for i != j {
		for i != j && comparator(j.Data, pivot) > 0 {
			j = j.Pre
		}
		if i != j {
			i.Data = j.Data
		}
		for i != j && comparator(i.Data, pivot) < 0 {
			i = i.Next
		}
		if i != j {
			j.Data = i.Data
		}
	}
	i.Data = pivot
	quickSort(start, i.Pre, comparator)
	quickSort(i.Next, end, comparator)
}
