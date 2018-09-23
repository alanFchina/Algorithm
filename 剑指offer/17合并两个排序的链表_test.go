package 剑指offer

import "testing"

func TestMergeList(t *testing.T) {
	cases := []struct {
		head1, head2 *ListNode
		wanted       []int
	}{
		{NewListInOrder(0), NewListInOrder(1), []int{0}},
		{NewListInOrder(1), NewListInOrder(1), []int{0, 0}},
		{NewListInOrder(0), NewListInOrder(0), []int{}},
		{NewListInOrder(2), NewListInOrder(3), []int{0, 0, 1, 1, 2}},
	}
	for i, v := range cases {
		newHead := MergeList(v.head1, v.head2)
		result := getSliceFromNode(newHead)
		if !equal(result, v.wanted) {
			t.Errorf("第%d个测试用例失败! 期望: %d, 返回: %d", i, v.wanted, result)
		}
	}
}

func NewListInOrder(n int) *ListNode {
	var head *ListNode
	var tail *ListNode
	for i := 0; i < n; i++ {
		this := ListNode{
			m_nValue: i,
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
