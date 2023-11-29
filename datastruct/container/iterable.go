package container

/**
Iterable container interface for vector, list, set and map.
*/

type Consumer[K any, V any] func(indexOrKey K, value V) bool

type Iterable[K any, V any] interface {
	Base
	// ForEach traverse the container and call consumer for each element
	// K is the type of key or index, V is the type of value
	ForEach(consumer Consumer[K, V])
}
