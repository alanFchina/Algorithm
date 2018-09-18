package base_algorithm

// 归并排序的循环实现版本, 使用到了递归版本中的Merge函数

func MergeSortLoop(arr []int) {
	// 存储已经排序好的序列
	q := make(queue, 0, len(arr))
	// 初始化为arr中的每个单独的值
	for i := 0; i < len(arr); i++ {
		q.push(i, i)
	}

	for q.len() > 1 {
		var isEven bool // 确定最后是否有一部pop push
		length := q.len()
		if length%2 == 0 {
			isEven = true
		} else {
			isEven = false
		}

		for i := 0; i < length/2; i++ {
			start1, end1, _ := q.pop()
			_, end2, _ := q.pop()
			Merge(arr, start1, end1, end2)
			q.push(start1, end2)
		}
		
		if !isEven {
			start, end, _ := q.pop()
			q.push(start, end)
		}
	}
}

// 实现一个队列

type queue []value

type value struct {
	start, end int
}

func (q *queue) push(start, end int) {
	*q = append(*q, value{start, end})
}

func (q *queue) pop() (int, int, bool) {
	if q.len() == 0 {
		return 0, 0, false
	}
	v := (*q)[0]
	*q = (*q)[1:]
	return v.start, v.end, true
}

func (q *queue) len() int {
	return len(*q)
}
