package base_algorithm

func HeapSort(arr []int) {
	h := BuildMaxHeap(arr)
	for i := len(h) - 1; i >= 1; i-- {
		h[0], h[i] = h[i], h[0]
		h = h[:i]
		h.maxHeapify(0)
	}
}

// 实现堆
type heap []int

func BuildMaxHeap(arr []int) heap {
	h := heap(arr)
	for i := h.parent(len(h) - 1); i >= 0; i-- {
		h.maxHeapify(i)
	}
	return h
}

func (h *heap) left(i int) int {
	return 2*i + 1
}

func (h *heap) right(i int) int {
	return 2*i + 2
}

func (h *heap) parent(i int) int {
	return (i - 1) / 2
}

func (h *heap) maxHeapify(i int) {
	l := h.left(i)
	r := h.right(i)
	largest := i
	if l < len(*h) && (*h)[l] > (*h)[i] {
		largest = l
	}
	if r < len(*h) && (*h)[r] > (*h)[largest] {
		largest = r
	}
	if largest != i {
		(*h)[i], (*h)[largest] = (*h)[largest], (*h)[i]
		h.maxHeapify(largest)
	}
}

func (h *heap) maxHeapInsert(v int) {
	*h = append(*h, v)
	i := len(*h) - 1
	for p := h.parent(i); p > 0 && (*h)[p] < v; {
		(*h)[i], (*h)[p] = (*h)[p], (*h)[i]
		i = p
		p = h.parent(i)
	}
}

func (h *heap) heapExtractMax() int {
	max, size := (*h)[0], len(*h)
	(*h)[0], (*h)[size-1] = (*h)[size-1], (*h)[0]
	*h = (*h)[:size-1]
	h.maxHeapify(0)
	return max
}

func (h *heap) heapMaximum() int {
	return (*h)[0]
}
