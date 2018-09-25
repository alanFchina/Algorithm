package 剑指offer

import "testing"

func TestIsPopOrder(t *testing.T) {
	cases := []struct {
		push, pop []int
		wanted    bool
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}, true},
		{[]int{1}, []int{1}, true},
		{[]int{}, []int{}, true},
	}

	for i, v := range cases {
		if result := IsPopOrder(v.push, v.pop); result != v.wanted {
			t.Errorf("第%d个测试用例失败!", i)
		}
	}
}
