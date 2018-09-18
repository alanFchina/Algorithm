package base_algorithm

import "testing"

func TestInsertSort(t *testing.T) {
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
