package 剑指offer

type queue struct {
	vals                  []int
	size, len, head, tail int
}

func NewQueue(size int) queue {
	return queue{
		vals: make([]int, size, size),
		size: size,
		len:  0,
		head: -1,
		tail: 0,
	}
}

func (q *queue) Push(v int) bool {
	if q.len == q.size {
		return false
	}
	q.head = (q.head + 1) % q.size
	q.vals[q.head] = v
	q.len++
	return true
}

func (q *queue) Pop() (int, bool) {
	if q.len == 0 {
		return 0, false
	}
	popVal := q.vals[q.tail]
	q.tail = (q.tail + 1) % q.size
	q.len--
	return popVal, true
}
