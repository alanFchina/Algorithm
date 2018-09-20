package base_algorithm

func CountSort(arr []int, n int) []int {
	c := make([]int, n, n)
	for _, v := range arr {
		c[v]++
	}
	for i := 1; i < len(c); i++ {
		c[i] = c[i-1] + c[i]
	}
	r := make([]int, len(arr), len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		r[c[arr[i]]-1] = arr[i]
		c[arr[i]]--
	}
	return r
}
