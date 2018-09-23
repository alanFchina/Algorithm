package 剑指offer

// 题目: 给定单项链表的头指针和一个结点指针,
// 定义一个函数在O(1)时间删除该结点
type ListNode struct {
	m_nValue int
	m_pNext *ListNode
}

func DeleteNode(node *ListNode) {
	node.m_nValue = node.m_pNext.m_nValue
	node.m_pNext = node.m_pNext.m_pNext
}
