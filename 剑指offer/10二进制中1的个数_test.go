package 剑指offer

import "testing"

func TestNumberOf1(t *testing.T) {
	cases := []struct {
		input, wanted int
	}{
		{input: 1, wanted: 1},
		{input: 0, wanted: 0},
		{input: 2, wanted: 1},
		{input: 3, wanted: 2},
		{input: 10, wanted: 2},
	}
	for i, v := range cases {
		if result := NumberOf1(v.input); result != v.wanted {
			t.Errorf("第%d个测试用例失败!", i)
		}
	}
}
