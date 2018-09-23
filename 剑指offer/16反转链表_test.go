package 剑指offer

import "testing"

func TestReverseList(t *testing.T) {
	cases := []struct {
		head   *ListNode
		wanted []int
	}{
		{NewListInReverseOrder(0), []int{}},
		{NewListInReverseOrder(1), []int{1}},
		{NewListInReverseOrder(10), []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}
	for i, v := range cases {
		newHead := ReverseList(v.head);
		result := getSliceFromNode(newHead)
		if !equal(result, v.wanted) {
			t.Errorf("第%d个测试用例失败! 期望: %d ,返回: %d", i, v.wanted, result)
		}
	}
}

func getSliceFromNode(head *ListNode) []int {
	numbers := make([]int, 0)
	for head != nil {
		numbers = append(numbers, head.m_nValue)
		head = head.m_pNext
	}
	return numbers
}
