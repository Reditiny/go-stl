package node

type LinkNode[T any] struct {
	Data T
	Next *LinkNode[T]
	Pre  *LinkNode[T]
}

func NewLinkNode[T any](data T) *LinkNode[T] {
	return &LinkNode[T]{
		Data: data,
		Next: nil,
		Pre:  nil,
	}
}

type TreeNode[T any] struct {
	data  T
	left  *TreeNode[T]
	right *TreeNode[T]
}

func NewTreeNode[T any](data T) *TreeNode[T] {
	return &TreeNode[T]{
		data:  data,
		left:  nil,
		right: nil,
	}
}
