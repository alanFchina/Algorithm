package 剑指offer

import (
	"testing"
	"fmt"
)

func TestReorderOddEven(t *testing.T) {
	cases := []struct {
		input []int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{1, 5, 2, 8, 23, 0, 1, 11, 5, 77}},
		{[]int{}},
		{[]int{1}},
		{[]int{1, 2}},
		{[]int{1, 2, 3}},
		{[]int{1, 1, 11,}},
	}
	for i, v := range cases {
		ReorderOddEven(v.input)
		fmt.Printf("输出:%d\n", v.input)
		if !isOddAllBeforeEven(v.input) {
			t.Errorf("第%d个测试用例失败!", i)
		}
	}
}

func isOddAllBeforeEven(arr []int) bool {
	even := false
	for _, v := range arr {
		if even && v%2 != 0 {
			return false
		}
		if v%2 == 0 {
			even = true
		}
	}
	return true
}
