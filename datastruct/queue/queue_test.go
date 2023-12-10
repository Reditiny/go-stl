package queue

import "testing"

func TestNewQueue(t *testing.T) {
	queue := NewQueue[int](10)
	for i := 0; i < 5; i++ {
		queue.EnQueue(i)
	}
	println(queue.Front())
	queue.DeQueue()
	for i := 0; i < 10; i++ {
		queue.EnQueue(i)
	}
	for !queue.Empty() {
		println(queue.Front())
		queue.DeQueue()
	}

}
