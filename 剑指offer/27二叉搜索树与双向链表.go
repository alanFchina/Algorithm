package 剑指offer

// 题目: 输入一棵二叉搜索树, 将该二叉搜索树转化成一个排序的双向链表.
// 要求: 不能创建新的结点,只能调整树中的结点指针的指向

func Convert(root *BinaryTreeNode) *BinaryTreeNode {
	if root == nil {
		return nil
	}
	head, _ := convertNode(root)
	return head
}

func convertNode(root *BinaryTreeNode) (*BinaryTreeNode, *BinaryTreeNode) {
	var (
		head1 = root
		head2 = root
		tail1 = root
		tail2 = root
	)
	if root.left != nil {
		head1, tail1 = convertNode(root.left)
		tail1.right = root
		root.left = tail1
	}
	if root.right != nil {
		head2, tail2 = convertNode(root.right)
		root.right = head2
		head2.left = root
	}
	return head1, tail2
}
