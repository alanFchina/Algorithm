package 剑指offer

// 题目: 请完成一个函数, 输入一个二叉树, 该函数输出它的镜像

func MirrorRecursively(root *BinaryTreeNode) {
	if root == nil {
		return
	}
	root.left, root.right = root.right, root.left
	MirrorRecursively(root.left)
	MirrorRecursively(root.right)
}
