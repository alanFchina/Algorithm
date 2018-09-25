package 剑指offer

import (
	"testing"
	"fmt"
)

func TestPrintMatrixClockwisely(t *testing.T) {
	cases := []struct {
		input [][]int
	}{
		{[][]int{}},
		{
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{2, 3, 4, 5},
			},
		},
		{
			[][]int{
				[]int{1},
				[]int{2},
			},
		},
		{
			[][]int{
				[]int{1, 2, 3, 4},
			},
		},
		{
			[][]int{
				[]int{1},
			},
		},
	}
	for _, v := range cases {
		fmt.Printf("%d start: \n", v)
		PrintMatrixClockwisely(v.input)
	}
}
