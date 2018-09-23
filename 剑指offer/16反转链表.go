package 剑指offer

// 题目: 定义一个函数输入一个链表的头结点,
// 反转该链表并输出反转后链表的头结点

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.m_pNext == nil {
		return head
	}

	i, j, k := head, head.m_pNext, head.m_pNext.m_pNext
	i.m_pNext = nil
	for {
		j.m_pNext = i
		if k == nil {
			break
		}
		i, j, k = j, k, k.m_pNext
	}
	return j
}
