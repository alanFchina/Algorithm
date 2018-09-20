package base_algorithm

import (
	"testing"
)

func TestCountSort(t *testing.T) {
	cases := []struct {
		input  []int
		n      int
		wanted []int
	}{
		{
			[]int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2},
			3,
			[]int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2},
		},
		{
			[]int{},
			0,
			[]int{},
		},
		{
			[]int{0, 0, 0, 0, 0, 0, 1, 1, 0, 1, 1, 2, 2, 2, 2, 2},
			3,
			[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 2},
		},
		{
			[]int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 0, 3},
			4,
			[]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3},
		},
	}
	for i, v := range cases {
		if r := CountSort(v.input, v.n); !equal(r, v.wanted) {
			t.Errorf("第%d个测试用例失败, 期望得到:%d, 返回%d", i, v.wanted, r)
		}
	}
}
