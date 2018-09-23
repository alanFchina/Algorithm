package 剑指offer

// 题目: 输入两个递增排序的链表, 合并这两个链表并使新链表中的结点仍然是按照递增排序的.

func MergeList(head1, head2 *ListNode) *ListNode {
	if head1 == nil && head2 == nil {
		return nil
	}
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}

	var newHead *ListNode

	if head1.m_nValue <= head2.m_nValue {
		newHead = head1
		newHead.m_pNext = MergeList(head1.m_pNext, head2)
	} else {
		newHead = head2
		newHead.m_pNext = MergeList(head1, head2.m_pNext)
	}
	return newHead
}
