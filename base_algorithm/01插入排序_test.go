package base_algorithm

import "testing"

func TestInsertSort(t *testing.T) {
	cases := []struct {
		input  []int
		wanted []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{}, []int{}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{1, 5, 3, 6, 12, 8, 9, 32, 6, 0, 1, 4, 4, 5, 4,}, []int{0, 1, 1, 3, 4, 4, 4, 5, 5, 6, 6, 8, 9, 12, 32}},
	}
	for i, v := range cases {
		if InsertSort(v.input); !equal(v.input, v.wanted) {
			t.Errorf("第%d个测试用例失败!", i)
		}
	}
}

func equal(arr1 []int, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
