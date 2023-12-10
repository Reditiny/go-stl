package container

type Sequence[T any] interface {
	Base
	// Get return the element at index idx
	// index should be scoped in [0,size)
	Get(index int) T
	// Set set the element at index idx to x
	// index should be scoped in [0,size)
	Set(index int, value T)
	// Insert insert element with value x before index idx
	// index should be scoped in [0,size)
	Insert(index int, value T)
	// Erase erase element at index idx
	// index should be scoped in [0,size).
	Erase(index int)
	// Front return the first element in vector
	// container must not be empty.
	Front() T
	// Back return the last element in vector.
	// container must not be empty.
	Back() T
	// PushBack add element x at the end of container
	PushBack(value T)
	// PopBack delete the last element if container is not empty
	PopBack()
}
