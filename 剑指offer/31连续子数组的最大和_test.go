package 剑指offer

import "testing"

func TestFindGreatestSumOfSubArray(t *testing.T) {
	cases := []struct {
		input  []int
		wanted int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 28},
		{[]int{-1, 2, 3, 4, 5, 6, 7}, 27},
		{[]int{-1}, 0},
		{[]int{-1, -2, -3, 4, -1, 5, -4}, 8},
	}

	for i, v := range cases {
		if r := FindGreatestSumOfSubArray(v.input); r != v.wanted {
			t.Errorf("第%d个测试用例出错, 期望%d, 返回%d", i, v.wanted, r)
		}
	}
}
