package 剑指offer

import "testing"

func TestMinNumber(t *testing.T) {
	cases := []struct {
		input  []int
		wanted int
	}{
		{
			input:  []int{1, 0, 1, 1, 1},
			wanted: 0,
		},
		{
			input:  []int{1, 1, 1, 0, 1},
			wanted: 0,
		},
		{
			input:  []int{3, 4, 0, 1, 2},
			wanted: 0,
		},
		{
			input:  []int{0, 0, 0, -1, -1},
			wanted: -1,
		},
		{
			input:  []int{3, 4, 5, 0, 0, 1, 1, 2, 2},
			wanted: 0,
		},
	}
	for i, v := range cases {
		if result := MinNumber(v.input); result != v.wanted {
			t.Errorf("第%d个测试用例失败, 期望:%d, 返回%d", i, v.wanted, result)
		}
	}
}
