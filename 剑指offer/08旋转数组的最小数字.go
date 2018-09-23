package 剑指offer

import (
	"math"
	"fmt"
)

// 题目: 把一个数组最开始的若干个元素搬到数组的末尾, 我们称之为数组的旋转.
// 输入一个递增的数组的一个旋转, 输出旋转数组的最小元素.
// 例如, 数组{3,4,5,1,2}为{1,2,3,4,5}的一个旋转, 该数组的最小值为1

func MinNumber(arr []int) int {
	length := len(arr)
	index1 := 0
	index2 := length - 1
	indexMid := index1
	for arr[index1] >= arr[index2] {
		if index2-index1 == 1 {
			indexMid = index2
			break
		}
		indexMid = (index1 + index2) / 2
		if arr[index1] == arr[index2] && arr[indexMid] == arr[index1] {
			return MinInOrder(arr[index1 : index2+1])
		}
		if arr[indexMid] >= arr[index1] {
			index1 = indexMid
		} else if arr[indexMid] <= arr[index2] {
			index2 = indexMid
		} else {
			fmt.Println("111111")
		}
	}
	return arr[indexMid]
}

func MinInOrder(arr []int) int {
	min := math.MaxInt64
	for _, v := range arr {
		if v < min {
			min = v
		}
	}
	return min
}
