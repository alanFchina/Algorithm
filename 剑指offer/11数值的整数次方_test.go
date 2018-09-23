package 剑指offer

import "testing"

func TestPower(t *testing.T) {
	cases := []struct {
		base     float64
		exponent int
		wanted   float64
	}{
		{10, 0, 1},
		{10, 1, 10},
		{9, 2, 81},
		{9, 3, 729},
	}
	for i, v := range cases {
		if result := Power(v.base, v.exponent); result != v.wanted {
			t.Errorf("第%d个测试用例失败", i)
		}
	}
}
