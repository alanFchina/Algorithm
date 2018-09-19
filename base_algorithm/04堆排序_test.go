package base_algorithm

import "testing"

func TestHeapSort(t *testing.T) {
	for i, v := range cases {
		if HeapSort(v.input); !equal(v.input, v.wanted) {
			t.Errorf("第%d个测试用例出错!", i)
		}
	}
}