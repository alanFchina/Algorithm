package 剑指offer

import (
	"testing"
)

func TestFindKthToTail(t *testing.T) {
	cases := []struct {
		head   *ListNode
		k      int
		wanted *ListNode
	}{
		{NewListInReverseOrder(0), 1, nil},
		{NewListInReverseOrder(1), 1, &ListNode{1, nil}},
		{NewListInReverseOrder(1), 2, nil},
		{NewListInReverseOrder(10), 9, &ListNode{9, nil}},
	}
	for i, v := range cases {
		result := FindKthToTail(v.head, v.k)
		if result != nil && v.wanted == nil || result == nil && v.wanted != nil {
			t.Errorf("第%d个测试用例失败", i)
			continue
		}
		if result != nil && v.wanted != nil && result.m_nValue != v.wanted.m_nValue {
			t.Errorf("第%d个测试用例失败", i)
		}
	}
}

func NewListInReverseOrder(n int) *ListNode {
	var head *ListNode
	var tail *ListNode
	for i := 0; i < n; i++ {
		this := ListNode{
			m_nValue: n - i,
			m_pNext:  nil,
		}
		if head == nil {
			head = &this
			tail = &this
		} else {
			tail.m_pNext = &this
			tail = &this
		}
	}
	return head
}
