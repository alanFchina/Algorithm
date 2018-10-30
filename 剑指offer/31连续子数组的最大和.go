package 剑指offer

func FindGreatestSumOfSubArray(nums []int) int {
	maxSum := 0
	sumIncludeLast := 0
	for _, v := range nums {
		if sumIncludeLast + v > v {
			sumIncludeLast += v
		} else {
			sumIncludeLast = v
		}
		if sumIncludeLast > maxSum {
			maxSum = sumIncludeLast
		}
	}
	return maxSum
}
