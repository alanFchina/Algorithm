package 剑指offer

// 题目: 定义栈的数据结构, 请在该类型中实现一个能够得到栈的最小元素的min函数,
// 在该栈中, 调用min push 及 pop的时间复杂度都是O(1)

type minStack struct {
	s1 stack // 存储栈
	s2 stack // 辅助栈
}

func NewMinStack(size int) minStack {
	s1 := NewStack(size)
	s2 := NewStack(size)
	return minStack{s1, s2}
}

func (s *minStack) push(v int) bool {
	ok := s.s1.Push(v)
	if !ok {
		return false
	}
	m, ok := s.min()
	if !ok {
		s.s2.Push(v)
	} else {
		min := m
		if v < m {
			min = v
		}
		s.s2.Push(min)
	}
	return true
}

func (s *minStack) pop() (int, bool) {
	v, ok := s.s1.Pop()
	if !ok {
		return 0, false
	}
	s.s2.Pop()
	return v, true
}

func (s *minStack) min() (int, bool) {
	m, ok := s.s2.Pop()
	if !ok {
		return 0, false
	}
	s.s1.Push(m)
	return m, true
}
