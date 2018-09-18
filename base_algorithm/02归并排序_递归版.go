package base_algorithm

// 归并排序是稳定的非原址排序, 之所以非原址是因为合并过程(merge)需要用到额外的存储空间
// f(n) = Θ(nlgn)
// 归并排序是渐进最优的比较模型下的排序, 没有坏的输入增加他的运行时间的量级
// 下面是归并排序的递归版本

import (
	"math"
)

func MergeSortRecursion(arr []int, i, j int) {
	if i >= j {
		return
	}
	q := (i + j) / 2
	MergeSortRecursion(arr, i, q)
	MergeSortRecursion(arr, q+1, j)
	Merge(arr, i, q, j)
}

func Merge(arr []int, i, q, j int) {
	arr1 := make([]int, 0, q-i+1+1)
	arr2 := make([]int, 0, j-q+1)
	for index := i; index <= q; index++ {
		arr1 = append(arr1, arr[index])
	}
	arr1 = append(arr1, math.MaxInt64)
	for index := q + 1; index <= j; index++ {
		arr2 = append(arr2, arr[index])
	}
	arr2 = append(arr2, math.MaxInt64)

	for index := i; index <= j; index++ {
		if arr1[0] <= arr2[0] {
			arr[index] = arr1[0]
			arr1 = arr1[1:]
		} else {
			arr[index] = arr2[0]
			arr2 = arr2[1:]
		}
	}
}
