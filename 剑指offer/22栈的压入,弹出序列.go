package 剑指offer

// 题目: 输入两个整数序列, 第一个序列表示栈的压入顺序, 请判断
// 第二个序列是否为该栈的弹出顺序. 假设压入栈的所有数字均不相等,

func IsPopOrder(push, pop []int) bool {
	if len(push) != len(pop) {
		return false
	}

	length := len(push)
	if length == 0 {
		return true
	}

	s := NewStack(length)
	s.Push(push[0])
	i, j := 1, 0
	for ; i < length; {
		top, _ := s.Top()
		if top == pop[j] {
			s.Pop()
			j++
		} else {
			s.Push(push[i])
			i++
		}
	}
	for pushVal, ok := s.Pop(); ok; j++ {
		if pushVal != pop[j] {
			return false
		}
		pushVal, ok = s.Pop()
	}
	return true
}
