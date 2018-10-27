package 剑指offer

import (
	"sort"
	"testing"
)

func TestMinKValue(t *testing.T) {
	inputs := []struct {
		input  []int
		k      int
		wanted []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 1, []int{1}},
		{[]int{1, 3, 9, 2, 5, 1}, 3, []int{1, 1, 2}},
	}

	for i, v := range inputs {
		if result := MinKValue(v.input, v.k); !valueEqual(result, v.wanted) {
			t.Errorf("第%d个测试用例失败, 期望%v, 返回%v", i, v.wanted, result)
		}
	}
}

func valueEqual(input, wanted []int) bool {
	sort.Ints(input)
	sort.Ints(wanted)
	return equal(input, wanted)
}
