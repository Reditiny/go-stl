package container

type Compare[T any] func(a, b T) int

type Sortable[T any] interface {
	Base
	Sort(compare Compare[T])
}
