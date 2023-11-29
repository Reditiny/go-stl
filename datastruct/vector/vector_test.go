package vector

import "testing"

var v *Vector[int]

func TestNewVector(t *testing.T) {
	v = NewVector[int]()
	printV()
}

func TestNewVectorByArray(t *testing.T) {
	v = NewVectorByArray[int]([]int{1, 3, 2, 5, 4, 9, 7, 8, 6})
	printV()
}

func TestVector_PushBack(t *testing.T) {
	for i := 0; i < 5; i++ {
		v.PushBack(i + 10)
	}
	printV()
}

func TestVector_PopBack(t *testing.T) {
	for i := 0; i < 4; i++ {
		if i%2 == 0 {
			v.PopBack()
		}
	}
	printV()
}

func TestVector_Size(t *testing.T) {
	printV()
}

func TestVector_Reserve(t *testing.T) {
	v.Reserve(32)
	printV()
}

func TestVector_Resize(t *testing.T) {
	v.Resize(32, 0)
	printV()
	v.Resize(23, 0)
	printV()
}

func TestVector_Insert(t *testing.T) {
	v.Insert(0, 28)
	printV()
}

func TestVector_Sort(t *testing.T) {
	v.Sort(func(a, b int) int {
		return a - b
	})
	printV()
}

func TestVector_Erase(t *testing.T) {
	v.Erase(0)
	printV()
}

func TestVector_ForEach(t *testing.T) {
	v := NewVector[int]()
	for i := 0; i < 10; i++ {
		v.PushBack(i)
	}
	v.ForEach(func(index, _ int) bool {
		if v.Get(index)%2 == 0 {
			v.Insert(index, 0)
			v.Erase(index)
		}
		return true
	})
	v.ForEach(func(index, _ int) bool {
		if v.Get(index) != index {
			t.Error("ForEach failed")
		}
		return true
	})
}

func TestVector_Capacity(t *testing.T) {
	printV()
}

func TestVector_Shrink(t *testing.T) {
	v.Shrink()
	printV()
}

func TestVector_Clear(t *testing.T) {
	v.Clear()
	printV()
}

func TestVector_Empty(t *testing.T) {
	printV()
}

func printV() {
	println("size:", v.Size())
	println("capacity:", v.Capacity())
	for i := 0; i < v.Size(); i++ {
		print(v.data[i])
		print(" ")
	}
	println()
}
