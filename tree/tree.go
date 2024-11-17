package tree

type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

func NewNode[T any](value T) *TreeNode[T] {
	node := &TreeNode[T]{
		Val:   value,
		Left:  nil,
		Right: nil,
	}

	return node
}

func (node *TreeNode[T]) InsertLeft(value T) *TreeNode[T] {
	node.Left = &TreeNode[T]{
		Val:   value,
		Left:  nil,
		Right: nil,
	}

	return node
}

func (node *TreeNode[T]) InsertRight(value T) *TreeNode[T] {
	node.Right = &TreeNode[T]{
		Val:   value,
		Left:  nil,
		Right: nil,
	}

	return node
}

// InOrder performs In-order traversal of tree
// Left-Node-Right
func (node *TreeNode[T]) InOrder() []T {
	var orderedItems []T
	if node.Left != nil {
		node.Left.InOrder()
	}
	orderedItems = append(orderedItems, node.Val)
	if node.Right != nil {
		node.Right.InOrder()
	}
	return orderedItems
}
