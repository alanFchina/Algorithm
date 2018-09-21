package 剑指offer

import "testing"

func TestRebuildBinaryTree(t *testing.T) {
	cases := []struct {
		firstTravelResult  []int
		secondTravelResult []int
	}{
		{
			[]int{1, 2, 4, 7, 3, 5, 6, 8},
			[]int{4, 7, 2, 1, 5, 3, 8, 6},
		},
		{
			[]int{},
			[]int{},
		},
	}
	for i, v := range cases {
		head := RebuildBinaryTree(v.firstTravelResult, v.secondTravelResult)
		getFirstResult := firstTravel(head, make([]int, 0, 10))
		getSecondResult := secondTravel(head, make([]int, 0, 10))
		if !equal(v.firstTravelResult, getFirstResult) ||
			!equal(v.secondTravelResult, getSecondResult) {
			t.Errorf("第%d个测试用例失败, first, second 分别返回:%d, %d", i, getFirstResult, getSecondResult)
		}
	}
}

func firstTravel(head *BinaryTreeNode, resp []int) []int {
	if head == nil {
		return resp
	}
	resp = append(resp, head.val)
	resp = firstTravel(head.left, resp)
	resp = firstTravel(head.right, resp)
	return resp
}

func secondTravel(head *BinaryTreeNode, resp []int) []int {
	if head == nil {
		return resp
	}

	resp = secondTravel(head.left, resp)
	resp = append(resp, head.val)
	resp = secondTravel(head.right, resp)
	return resp
}
