package 剑指offer

import "testing"

func TestNewQueue(t *testing.T) {
	cases := []struct {
		input  []int
		wanted []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			wanted: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			input:  []int{},
			wanted: []int{},
		},
	}

	for i, v := range cases {
		resp := make([]int, 0, len(v.input))
		q := NewQueue(len(v.input))
		for _, val := range v.input {
			q.Push(val)
		}

		for v, ok := q.Pop(); ok; {
			resp = append(resp, v)
			v, ok = q.Pop()
		}
		if !equal(v.wanted, resp) {
			t.Errorf("第%d个测试用例失败, 期望: %d, 返回 %d", i, v.wanted, resp)
		}
	}
}

func TestNewQueueByStack(t *testing.T) {
	cases := []struct {
		input  []int
		wanted []int
	}{
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7, 8},
			wanted: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			input:  []int{},
			wanted: []int{},
		},
	}

	for i, v := range cases {
		resp := make([]int, 0, len(v.input))
		q := newQueueByStack(len(v.input))
		for _, val := range v.input {
			q.push(val)
		}

		for v, ok := q.pop(); ok; {
			resp = append(resp, v)
			v, ok = q.pop()
		}
		if !equal(v.wanted, resp) {
			t.Errorf("第%d个测试用例失败, 期望: %d, 返回 %d", i, v.wanted, resp)
		}
	}
}
