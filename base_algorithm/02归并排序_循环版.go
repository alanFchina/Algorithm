package base_algorithm

// 归并排序的循环实现版本, 使用到了递归版本中的Merge函数

func MergeSortLoop(arr []int) {
	// 存储已经排序好的序列
	q := make(queue, 0, len(arr))
	// 初始化为arr中的每个单独的值
	for i := range arr {
		q.push(i, i)
	}

	for q.len() > 1 {
		length := q.len()
		isEven := length%2 == 0 // 确定最后是否有一部pop push

		for i := 0; i < length/2; i++ {
			start, middle, _ := q.pop()
			_, end, _ := q.pop()
			Merge(arr, start, middle, end)
			q.push(start, end)
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

// 为slice这种重命名的结构添加方法, 涉及到修改原数据也要用指针接收者
// 原本以为因为参数传递的时候不用考虑slice指针, 所以接收者也一样(这是错的)
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
