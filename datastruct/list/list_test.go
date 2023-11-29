package list

import "testing"

var l *List[int]

func TestNewList(t *testing.T) {
	l = NewList[int]()
	l.PushBack(3)
	l.PushBack(4)
	l.PushFront(2)
	l.PushFront(1)
	printList()
}

func TestList_Get(t *testing.T) {
	println(l.Get(2))
	println(l.Get(1))
	println(l.Get(3))
}

func TestList_Set(t *testing.T) {
	l.Set(2, 5)
	l.Set(1, 6)
	l.Set(3, 7)
	printList()
}

func TestList_Insert(t *testing.T) {
	l.Insert(0, 11)
	l.Insert(2, 12)
	l.Insert(4, 10)
	printList()
}

func TestList_Erase(t *testing.T) {
	l.Erase(6)
	l.Erase(0)
	l.Erase(3)
	printList()
}

func TestList_ForEach(t *testing.T) {
	l.ForEach(func(index int, _ int) bool {
		if index%2 == 0 {
			l.Set(index, index-1)
		} else {
			l.Set(index, index+1)
		}
		return true
	})
	printList()
}

func TestList_Sort(t *testing.T) {
	l.Sort(func(a, b int) int {
		return a - b
	})
	printList()
}

func TestList_Clear(t *testing.T) {
	l.Clear()
	printList()
}

func printList() {
	println("size:", l.size)
	for cur := l.head.Next; cur != l.tail; cur = cur.Next {
		print(cur.Data, " ")
	}
	println()
}
