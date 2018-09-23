package 剑指offer

func NumberOf1(n int) int {
	num := 0
	for n != 0 {
		n = n&(n-1)
		num++
	}
	return num
}
