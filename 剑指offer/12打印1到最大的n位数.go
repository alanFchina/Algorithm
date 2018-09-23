package 剑指offer

import (
	"fmt"
)

// 题目: 输入数字n, 按顺序打印出从1到最大的n位十进制数.
// 比如输入3, 则打印出1, 2, 3 一直到最大的3位数即999

func Print1ToMaxOfNDigits(n int) {
	if n == 0 {
		return
	}
	arr := make([]int8, n, n)
	until := make([]int8, n, n)
	for i := range arr {
		arr[i] = 0
		until[i] = 9
	}

	for !equalInt8(arr, until) {
		i := n - 1
		for flag := true; flag && i != -1; i-- {
			arr[i] = (arr[i] + 1) % 10
			if arr[i] != 0 {
				flag = false
			}
		}
		PrintArr(arr)
	}
}

func equalInt8(arr1 []int8, arr2 []int8) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func PrintArr(arr []int8) {
	begin := false
	for _, v := range arr {
		if v != 0 && !begin {
			begin = true
		}
		if begin {
			fmt.Printf("%d", v)
		}
	}
	fmt.Print("\n")
}
