package 剑指offer

import "testing"

type testCase struct {
	input [][]int
	search int
	wanted bool
}

func TestArraySearch(t *testing.T) {
	cases := []testCase{
		{
			[][]int{
				{1,2,3,4},
				{2,3,4,5},
				{3,4,5,6},
				{4,5,6,7},
			},
			4,
			true,
	},
	{
			[][]int{
				{1,2,3,4},
				{2,3,4,5},
				{3,4,5,6},
				{4,5,6,7},
			},
			6,
			true,
	},
	{
			[][]int{
				{1,2,3,4},
				{2,3,4,5},
				{3,4,5,6},
				{4,5,6,7},
			},
			7,
			true,
	},
	{
			[][]int{
				{1,2,3,4},
				{2,3,4,5},
				{3,4,5,6},
				{4,5,6,7},
			},
			8,
			false,
	},
	{
			[][]int{
				{1,2,3,4},
				{2,3,4,5},
				{3,4,5,6},
				{4,5,6,7},
			},
			9,
			false,
	},
	}

	for i, mem := range cases{
		if mem.wanted != ArraySearch(mem.input, mem.search){
			t.Errorf("第%d个测试用例失败!", i)
		}
	}
}