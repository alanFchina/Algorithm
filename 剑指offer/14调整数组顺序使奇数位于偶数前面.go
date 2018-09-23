package 剑指offer

// 题目: 输入一个整数数组, 实现一个函数来调整该数组中数字的顺序,
// 使得所有奇数位于数组的前半部分, 所有偶数位于数组的后半部分.

func ReorderOddEven(arr []int) {
	i := 0
	j := len(arr) - 1
	for i < j {
		if arr[i]%2 != 0 {
			i++
			continue
		}
		if arr[j]%2 == 0 {
			j--
			continue
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
}
