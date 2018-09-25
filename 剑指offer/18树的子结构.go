package 剑指offer

type TreeNode struct {
	m_nValue          int
	m_pLeft, m_pRight *TreeNode
}

func HasSubTree(root1 *TreeNode, root2 *TreeNode) bool {
	result := false
	if root1 != nil && root2 != nil {
		if root1.m_nValue == root2.m_nValue {
			result = DoseTree1HaveTree2(root1, root2)
		}
		if !result {
			result = HasSubTree(root1.m_pLeft, root2)
		}
		if !result {
			result = HasSubTree(root1.m_pRight, root2)
		}
	}
	return result
}

func DoseTree1HaveTree2(root1 *TreeNode, root2 *TreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}
	if root1.m_nValue != root2.m_nValue {
		return false
	}
	return DoseTree1HaveTree2(root1.m_pLeft, root2.m_pLeft) &&
		DoseTree1HaveTree2(root1.m_pRight, root2.m_pRight)
}
