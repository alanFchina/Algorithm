package 剑指offer

import "testing"

func TestFibonacci(t *testing.T) {
	cases := []struct {
		input, wanted int
	} {
		{
			input: 0,
			wanted: 0,
		},
		{
			input: 1,
			wanted: 1,
		},
		{
			input: 2,
			wanted: 1,
		},
		{
			input: 3,
			wanted: 2,
		},
		{
			input:4,
			wanted: 3,
		},
	}
	for i, v := range cases {
		if result:=Fibonacci(v.input); result!=v.wanted{
			t.Errorf("第%d个用例失败!", i)
		}
	}
}
