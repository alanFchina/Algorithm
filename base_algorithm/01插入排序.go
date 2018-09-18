package base_algorithm

func InsertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i
		for ; j >0 && arr[j-1] > key; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = key
	}
}
