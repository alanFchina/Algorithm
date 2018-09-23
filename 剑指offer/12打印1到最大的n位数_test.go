package 剑指offer

import "testing"

func TestPrint1ToMaxOfNDigits(t *testing.T) {
	cases := []struct {
		input int
	}{
		{0},
		{1},
		{4},
	}

	for _, v := range cases {
		Print1ToMaxOfNDigits(v.input)
	}
}
