package 剑指offer

// 队列
type queue struct {
	s_1 stack
	s_2 stack
}

func NewQueue(size int) queue {
	return queue{
		s_1: NewStack(size),
		s_2: NewStack(size),
	}
}

func (q *queue) push(v int) bool {
	return q.s_1.push(v)
}

func (q *queue) pop() (int, bool) {
	if q.s_2.top == -1 {
		if q.s_1.top == -1 {
			return 0, false
		}
		for v, ok := q.s_1.pop(); ok; {
			q.s_2.push(v)
			v, ok = q.s_1.pop()
		}
	}
	v, _ := q.s_2.pop()
	return v, true
}

// 栈
type stack struct {
	vals []int
	top  int
	size int
}

func NewStack(size int) stack {
	return stack{
		vals: make([]int, size, size),
		top:  -1,
		size: size,
	}
}

func (s *stack) push(v int) bool {
	if s.top == s.size-1 {
		return false
	}
	s.top++
	s.vals[s.top] = v
	return true
}

func (s *stack) pop() (int, bool) {
	if s.top == -1 {
		return 0, false
	}
	topVal := s.vals[s.top]
	s.top--
	return topVal, true
}
