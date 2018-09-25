package 剑指offer

// 题目: 输入一个整数数组, 判断该数组是不是某二叉搜索树的后序遍历的结果.
// 如果是返回true, 否则返回false. 假设输入的数组的任意两个数字都互不相同

func VerifySequenceOfBST(sequence []int) bool {
	size := len(sequence)
	if size == 2 || size == 1 || size == 0 {
		return true
	}
	root := sequence[size-1]

	flag, index := false, 0
	for i, v := range sequence {
		if v >= root && !flag {
			index, flag = i, true
		}
		if v < root && flag {
			return false
		}
	}
	return VerifySequenceOfBST(sequence[:index]) && VerifySequenceOfBST(sequence[index:size-1])
}
