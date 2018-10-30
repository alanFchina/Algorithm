package 剑指offer

func MinKValue(input []int, k int) []int {
	if len(input) <= k {
		return input
	}
	h := New(k)
	for i, v := range input {
		if i < k{
			h.insert(v)
		} else {
			if v < h.max() {
				h[0] = v
				h.heapify()
			}
		}
	}
	return h
}

//  有Heapify操作的最大堆
type heap []int

func New(size int) heap {
	return make(heap, 0, size)
}

func (h *heap) max() int {
	return (*h)[0]
}

func (h *heap) heapify() {
	for i := 0; ; {
		maxIndex := i
		l := h.left(i)
		r := h.right(i)
		if l <= len(*h) && (*h)[l] > (*h)[maxIndex] {
			maxIndex = l
		}
		if r <= len(*h) && (*h)[r] > (*h)[maxIndex] {
			maxIndex = r
		}
		if maxIndex == i {
			break
		}
		(*h)[i], (*h)[maxIndex] = (*h)[maxIndex], (*h)[i]
		i = maxIndex
	}
}

func (h *heap) insert(v int) {
	*h = append(*h, v)
	for i := len(*h) - 1; ; {
		p := h.parent(i)
		if p >= 0 && (*h)[p] < (*h)[i] {
			(*h)[p], (*h)[i] = (*h)[i], (*h)[p]
			i = p
		} else {
			break
		}
	}
}

func (h *heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *heap) left(i int) int {
	return 2*i + 1
}

func (h *heap) right(i int) int {
	return 2*i + 2
}


