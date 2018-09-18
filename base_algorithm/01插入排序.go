package base_algorithm

// 插入排序是稳定的原址排序, 但是对于不同的输入, 时间复杂度不同
// f(n) = O(n²); f(n) = Ω(n)

// 使用二分法对插入排序进行优化:
//     不可行, 因为移动元素所花费的时间省不掉

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
