package 剑指offer

func Fibonacci(n int) int {
	first := 1
	second := 0
	for i := 0; i < n; i++ {
		first, second = second, first+second
	}
	return second
}
