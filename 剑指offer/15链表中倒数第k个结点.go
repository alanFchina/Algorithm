package 剑指offer

// 题目: 输入一个链表,输出该链表中倒数第k个结点.
// 为了符合大多数人的习惯, 本体从1开始计数, 即链表的结尾结点是倒数第一个结点
// 例如一个链表有6个结点, 从头结点开始他们的值依次是1, 2, 3, 4, 5, 6.
// 这个链表的倒数第三个结点的值是4

func FindKthToTail(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	first := head
	second := head
	for i := 1; i < k; i++ {
		first = first.m_pNext
		if first == nil {
			return nil
		}
	}
	for first.m_pNext != nil {
		first = first.m_pNext
		second = second.m_pNext
	}
	return second
}
