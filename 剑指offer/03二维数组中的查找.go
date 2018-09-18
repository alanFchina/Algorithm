package 剑指offer

// 在一个二维数组中, 每一行都按照从左到右递增的顺序排序,
// 每一列都按照从上到下递增的顺序排序. 请完成一个函数,
// 输入这样的一个二维数组和一个整数, 判断数组中是否含有该整数

// 思路:
// 从数组的右上角开始查找, 如果等于则返回true,
// 如果右上角大于要查找的值, 则删除最后一列递归调用
// 如果右上角小雨要查找的值, 则删除第一行递归调用

// 也可以从左下角开始查找

func ArraySearch(arr [][]int, searched int) bool {
	rowNum := len(arr)
	// 一直删减行直到行数为0
	if rowNum == 0 {
		return false
	}

	colNum := len(arr[0])
	// 一直删减列直到列数为0
	if colNum == 0 {
		return false
	}

	key := arr[0][colNum-1]
	if key == searched {
		return true
	} else if key > searched {
		// 给arr的每一行删除一列
		for i, row := range arr {
			arr[i] = row[0 : colNum-1]
		}
	} else if key < searched {
		// 删除arr的第一行
		arr = arr[1:]
	}
	return ArraySearch(arr, searched)
}
