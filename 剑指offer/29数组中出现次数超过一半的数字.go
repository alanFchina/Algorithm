package 剑指offer

func MoreThanHalfNum(nums []int) int {
	var num, count int
	for _, n := range nums {
		if n==num {
			count++
			continue
		}
		if count == 0 {
			num = n
			count++
		} else if count > 0 {
			count--
		}
	}
	return num
}
