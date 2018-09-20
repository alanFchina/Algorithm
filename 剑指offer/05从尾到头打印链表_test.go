package 剑指offer

import "testing"

func TestPrintListBack(t *testing.T) {
	for i := 0; i < 10; i++ {
		head, wanted := NewList(i)
		if result := printListBack(head); !equal(wanted, result) {
			t.Errorf("第%d个测试用例失败! 期望得到%d, 返回%d", i, wanted, result)
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
