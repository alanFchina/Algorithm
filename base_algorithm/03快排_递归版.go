package base_algorithm

// 快排是原址的不稳定排序
// 块排的最坏运行时间为O(n²), 也就是当输入为已经排序好的序列时
// 块排的平均运行时间为O(nlgn)

import (
	"math/rand"
)

func QuickSortRecursion(arr []int, start, end int) {
	if start >= end {
		return
	}

	middle := Partition(arr, start, end)
	QuickSortRecursion(arr, start, middle-1)
	QuickSortRecursion(arr, middle+1, end)

}

func QuickSortRecursionRandom(arr []int, start, end int) {
		if start >= end {
		return
	}

	middle := RandomPartition(arr, start, end)
	QuickSortRecursion(arr, start, middle-1)
	QuickSortRecursion(arr, middle+1, end)
}

func Partition(arr []int, start, end int) int {
	small := start
	for large := start + 1; large <= end; large++ {
		if arr[large] < arr[start] {
			small++
			arr[small], arr[large] = arr[large], arr[small]
		}
	}
	arr[start], arr[small] = arr[small], arr[start]
	return small
}

// 最常见的优化方法为随机化, 在分割过程中交换随机的一个元素和第一个元素
// 期望运行时间为O(nlgn)
// 对于数组中所有的数都是同一个值, 同样也会存在时间复杂度为n²的情况
func RandomPartition(arr []int, start, end int) int {
	r := rand.Intn(end-start+1) + start
	arr[start], arr[r] = arr[r], arr[start]
	return Partition(arr, start, end)
}
