package vector

/**
Vector is a sequence container representing array, which can automatically expand size.
Vector implements Iterable and Sortable interface.
*/

import (
	"go-stl/datastruct/container"
)

type Vector[T any] struct {
	size int
	data []T
}

func NewVectorByArray[T any](tArray []T) *Vector[T] {
	return &Vector[T]{
		size: len(tArray),
		data: tArray,
	}
}

func NewVector[T any]() *Vector[T] {
	return &Vector[T]{
		size: 0,
		data: make([]T, 0),
	}
}

// Reserve reserve memory for vector
// len is floor of capacity as capacity expansion is automatically implemented by append()
func (v *Vector[T]) Reserve(len int) {
	if len < 0 {
		panic("len " + string(rune(len)) + " is invalid")
	}
	if v.Capacity() < len {
		v.data = append(v.data, make([]T, len-v.size)...)
	}
}

func (v *Vector[T]) Get(index int) T {
	if index >= v.size || index < 0 {
		panic("index " + string(rune(index)) + " out of range " + string(rune(v.size)))
	}
	return v.data[index]
}

func (v *Vector[T]) Set(index int, x T) {
	if index >= v.size || index < 0 {
		panic("index " + string(rune(index)) + " out of range " + string(rune(v.size)))
	}
	v.data[index] = x
}

func (v *Vector[T]) Front() T {
	if v.size == 0 {
		panic("vector is empty")
	}
	return v.data[0]
}

func (v *Vector[T]) Back() T {
	if v.size == 0 {
		panic("vector is empty")
	}
	return v.data[v.size-1]
}

func (v *Vector[T]) Insert(index int, x T) {
	if index >= v.size || index < 0 {
		panic("index " + string(rune(index)) + " out of range " + string(rune(v.size)))
	}
	v.PushBack(x)
	for i := v.size - 1; i > index; i-- {
		v.data[i] = v.data[i-1]
	}
	v.data[index] = x
}

func (v *Vector[T]) Erase(index int) {
	if index >= v.size || index < 0 {
		panic("index " + string(rune(index)) + " out of range " + string(rune(v.size)))
	}
	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.PopBack()
}

func (v *Vector[T]) PushBack(x T) {
	v.data = append(v.data, x)
	v.size++
}

func (v *Vector[T]) PopBack() {
	if v.size != 0 {
		v.size--
		v.data = v.data[:v.size]
	}
}

func (v *Vector[T]) Size() int {
	return v.size
}

func (v *Vector[T]) Empty() bool {
	return v.size == 0
}

func (v *Vector[T]) Clear() {
	v.size = 0
}

// Resize resize the size of vector and truncate or expand with defaultValue
func (v *Vector[T]) Resize(size int, defaultValue T) {
	if size < 0 {
		panic("size " + string(rune(size)) + " is invalid")
	}
	if v.Capacity() < size {
		v.data = append(v.data, make([]T, size-v.size)...)
	}
	if v.size < size {
		for i := v.size; i < size; i++ {
			v.data[i] = defaultValue
		}
	}
	v.size = size
}

// Capacity return the capacity of vector
func (v *Vector[T]) Capacity() int {
	return cap(v.data)
}

// Shrink try to shrink the capacity of vector to fit the size of vector
func (v *Vector[T]) Shrink() {
	if v.Capacity() > 2*v.Size() {
		newData := make([]T, 2*v.size)
		copy(newData, v.data)
		v.data = newData
	}
}

func (v *Vector[T]) ForEach(consumer container.Consumer[int, T]) {
	for i := 0; i < v.size; i++ {
		if !consumer(i, v.data[i]) {
			break
		}
	}
}

func (v *Vector[T]) Sort(compare container.Compare[T]) {
	quickSort(v.data, compare, 0, v.size-1)
}

func quickSort[T any](arr []T, compare container.Compare[T], left int, right int) {
	if left < right {
		pivot := arr[left]
		i := left
		j := right
		for i < j {
			for i < j && compare(arr[j], pivot) > 0 {
				j--
			}
			if i < j {
				arr[i] = arr[j]
				i++
			}
			for i < j && compare(arr[i], pivot) < 0 {
				i++
			}
			if i < j {
				arr[j] = arr[i]
				j--
			}
		}
		arr[i] = pivot
		quickSort(arr, compare, left, i-1)
		quickSort(arr, compare, i+1, right)
	}
}
