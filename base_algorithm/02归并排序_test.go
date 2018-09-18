package base_algorithm

import "testing"

func TestMergeSortRecursion(t *testing.T) {
	for i, v := range cases {
		if MergeSortRecursion(v.input, 0, len(v.input)-1); !equal(v.input, v.wanted) {
			t.Errorf("第%d个测试用例出错!", i)
		}
	}
}

func TestMergeSortLoop(t *testing.T) {
	for i, v := range cases {
		if MergeSortLoop(v.input); !equal(v.input, v.wanted) {
			t.Errorf("第%d个测试用例出错!", i)
		}
	}
}
