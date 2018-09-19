package 剑指offer

import "testing"

func TestReplaceSpace(t *testing.T) {
	// 0表示空余空间
	// l表示需要处理的字节长度
	cases := []struct{
		input []byte
		l int
		wanted []byte}{
		{
			[]byte("aasdf asdf asdf0000"),
			15,
			[]byte("aasdf%20asdf%20asdf"),
		},
		{
			[]byte(""),
			0,
			[]byte(""),
		},
		{
			[]byte(" 00"),
			1,
			[]byte("%20"),
		},
		{
			[]byte("asdas"),
			5,
			[]byte("asdas"),
		},
	}
	for i,v := range cases {
		if replaceSpace(v.input, v.l);!bytesEqual(v.input, v.wanted) {
			t.Errorf("第%d个测试用例失败!", i)
		}
	}
}

func bytesEqual(first, second []byte) bool {
	if len(first) != len(second) {
		return false
	}
	for i := range first {
		if first[i] != second[i]{
			return false
		}
	}
	return true
}