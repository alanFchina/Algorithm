package 剑指offer

type UltimateStack []int

func (s *UltimateStack) Pop() int {
	topVal := s.Top()
	*s = (*s)[0 : s.Len()-1]
	return topVal
}

func (s *UltimateStack) Push(v int) {
	*s = append(*s, v)
}

func (s *UltimateStack) Len() int {
	return len(*s)
}

func (s *UltimateStack) Top() int {
	if len(*s) == 0 {
		panic("栈为空!")
	} else {
		return (*s)[s.Len()-1]
	}
}
