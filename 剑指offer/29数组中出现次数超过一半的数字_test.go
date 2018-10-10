package 剑指offer

import "testing"

func TestMoreThanHalfNum(t *testing.T) {
	cases := []struct {
		input  []int
		wanted int
	}{
		{[]int{1, 2, 3, 4, 5, 5, 5, 5, 5, 5, 5}, 5},
		{[]int{1, 1, 1, 1, 2, 2, 2, 2, 2}, 2},
		{[]int{1, 2, 2}, 2},
	}

	for i, v := range cases {
		if output := MoreThanHalfNum(v.input); output != v.wanted {
			t.Errorf("第%d个此时用例出错, 期望: %d, 返回: %d\n", i, v.wanted, output)
		}
	}
}
