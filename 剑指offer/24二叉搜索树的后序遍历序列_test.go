package 剑指offer

import "testing"

func TestVerifySquenceOfBST(t *testing.T) {
	cases := []struct {
		input  []int
		wanted bool
	}{
		{[]int{5, 7, 6, 9, 11, 10, 8}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}, true},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, true},
		{[]int{6,5,2,1,3}, false},
	}

	for i, v := range cases {
		if result := VerifySequenceOfBST(v.input); result != v.wanted {
			t.Errorf("第%d个测试用例失败!", i)
		}
	}
}
