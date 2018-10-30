package 剑指offer

import "testing"

func TestNumOf1Betwween1AndN(t *testing.T) {
	cases := []struct {
		input  int
		wanted int
	}{
		{21345, 1},
	}
	for i, v := range cases {
		if r := NumOf1Betwween1AndN(v.input); r != v.wanted {
			t.Errorf("第%d个测试用例失败, 期望%d, 返回%d", i, v.wanted, r)
		}
	}
}
