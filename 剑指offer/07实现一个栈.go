package 剑指offer

// 栈
type stack struct {
	vals []int
	size int
}

func NewStack(size int) stack {
	return stack{
		vals: make([]int, 0, size),
		size: size,
	}
}

func (s *stack) Push(v int) bool {
	if len(s.vals) == s.size {
		return false
	}
	s.vals = append(s.vals, v)
	return true
}

func (s *stack) Pop() (int, bool) {
	length := len(s.vals)
	if length == 0 {
		return 0, false
	}
	topVal := s.vals[length-1]
	s.vals = s.vals[:length-1]
	return topVal, true
}

func (s *stack) Top() (int, bool) {
	length := len(s.vals)
	if length == 0 {
		return 0, false
	}
	return s.vals[length-1], true
}
