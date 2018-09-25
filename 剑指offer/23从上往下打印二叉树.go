package 剑指offer

func FromTopToBottom(root *BinaryTreeNode, f func(node *BinaryTreeNode)) {
	if root == nil {
		return
	}
	queue := make([](*BinaryTreeNode), 0)
	queue = append(queue, root)
	for i := 0; i < len(queue); i++ {
		first := queue[i]
		f(first)
		if first.left != nil {
			queue = append(queue, first.left)
		}
		if first.right != nil {
			queue = append(queue, first.right)
		}
	}
}
