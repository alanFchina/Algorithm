package base_algorithm

import "testing"

func TestQuickSortRecursion(t *testing.T) {
	for i, v := range cases {
		if QuickSortRecursion(v.input, 0, len(v.input)-1); !equal(v.input, v.wanted) {
			t.Errorf("第%d个测试用例出错!", i)
		}
	}
}

func TestQuickSortRecursionRandom(t *testing.T) {
	for i, v := range cases {
		if QuickSortRecursionRandom(v.input, 0, len(v.input)-1); !equal(v.input, v.wanted) {
			t.Errorf("第%d个测试用例出错!", i)
		}
	}
}

func TestQuickSortLoop(t *testing.T) {
	for i, v := range cases {
		if QuickSortLoop(v.input); !equal(v.input, v.wanted) {
			t.Errorf("第%d个测试用例出错!", i)
		}
	}
}