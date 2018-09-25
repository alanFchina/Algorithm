package 剑指offer

import "fmt"

// 题目: 输入一棵二叉树和一个整数, 打印出二叉树中结点值的和为输入整数的所有路径.
// 从输的根节点开始往下一直到叶结点所经过的结点形成一条路径

func FindPath(root *BinaryTreeNode, expected int) {
	s := make(UltimateStack, 0)
	findPath(root, expected, &s)
}

func findPath(root *BinaryTreeNode, expected int, s *UltimateStack) {
	if root == nil {
		return
	}
	s.Push(root.val)
	if root.left == nil || root.right == nil  {
		if val := cal(s); val == expected {
			fmt.Println(*s)
		}
	}

	if root.left != nil {
		findPath(root.left, expected, s)
	}
	if root.right != nil {
		findPath(root.right, expected, s)
	}
	s.Pop()
}

func cal(path *UltimateStack) int {
	sum := 0
	for _, v := range *path {
		sum += v
	}
	return sum
}
