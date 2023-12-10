package queue

import "testing"

func TestNewPriorityQueue(t *testing.T) {
	queue := NewPriorityQueue[int](func(a, b int) int {
		return b - a
	})
	data1 := []int{1, 3, 5, 7, 9}
	data2 := []int{4, 3, 2, 1, 0}
	for i, v := range data1 {
		queue.Push(v)
		if i%2 == 1 {
			println(queue.Top())
			queue.Pop()
		}
	}
	for i, v := range data2 {
		queue.Push(v)
		if i%2 == 0 {
			println(queue.Top())
			queue.Pop()
		}
	}
	for !queue.Empty() {
		println(queue.Top())
		queue.Pop()
	}
}
