package 剑指offer

// 队列
type queueByStack struct {
	s_1 stack
	s_2 stack
}

func newQueueByStack(size int) queueByStack {
	return queueByStack{
		s_1: NewStack(size),
		s_2: NewStack(size),
	}
}

func (q *queueByStack) push(v int) bool {
	return q.s_1.Push(v)
}

func (q *queueByStack) pop() (int, bool) {
	if len(q.s_2.vals) == 0 {
		if len(q.s_1.vals) == 0 {
			return 0, false
		}
		for v, ok := q.s_1.Pop(); ok; {
			q.s_2.Push(v)
			v, ok = q.s_1.Pop()
		}
	}
	v, _ := q.s_2.Pop()
	return v, true
}
