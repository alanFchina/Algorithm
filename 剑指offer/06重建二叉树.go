package 剑指offer

// 题目: 输入某二叉树的前序遍历和中序遍历的结果, 请重建出该二叉树.
// 假设输入的前序遍历和中序遍历的结果中都不含重复的数字.
// 例如输入前序遍历序列{1,2,4,7,3,5,6,8}和中序遍历序列{4,7,2,1,5,3,8,6}

func RebuildBinaryTree(arr1 []int, arr2 []int) *BinaryTreeNode {
	if len(arr1) == 0 {
		return nil
	}
	head := &BinaryTreeNode{
		val:   arr1[0],
		left:  nil,
		right: nil,
	}

	i, _ := IndexOf(arr2, arr1[0])
	head.left = RebuildBinaryTree(arr1[1:i+1], arr2[:i])
	head.right = RebuildBinaryTree(arr1[i+1:], arr2[i+1:])
	return head
}

func IndexOf(arr []int, f int) (int, bool) {
	for i, v := range arr {
		if v == f {
			return i, true
		}
	}
	return 0, false
}

type BinaryTreeNode struct {
	val   int
	left  *BinaryTreeNode
	right *BinaryTreeNode
}
