package base_algorithm

// 快排的循环版本, 使用了递归里的Partition函数, 也可以使用RandomPartition

func QuickSortLoop(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	s := make(stack, 0, length)
	s.push(0, length-1)

	for s.len() != 0 {
		start, end, _ := s.pop()
		middle := Partition(arr, start, end)
		if end-(middle+1) > 0 {
			s.push(middle+1, end)
		}
		if (middle-1)-start > 0 {
			s.push(start, middle-1)
		}
	}
}

// 实现一个栈
type stack []value

func (s *stack) push(start, end int) {
	*s = append(*s, value{start, end})
}

func (s *stack) pop() (int, int, bool) {
	length := s.len()
	if length == 0 {
		return 0, 0, false
	}
	v := (*s)[length-1]
	*s = (*s)[:length-1]
	return v.start, v.end, true
}

func (s *stack) len() int {
	return len(*s)
}
