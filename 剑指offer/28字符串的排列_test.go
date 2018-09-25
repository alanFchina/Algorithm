package 剑指offer

import "testing"

func TestPermutation(t *testing.T) {
	cases := []struct{
		input string
	} {
		{"abc"},
		{"a"},
		{""},
		{"abcd"},
	}

	for _, v := range cases {
		Permutation(v.input)
	}
}
