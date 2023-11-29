package container

/**
Base container interface for all containers
*/

type Base interface {
	// Empty return true if container is empty
	Empty() bool
	// Size return the number of elements in container
	Size() int
	// Clear clear the all the elements of container
	Clear()
}
