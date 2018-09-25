package 剑指offer

import "fmt"

// 题目: 输入一个矩阵, 按照从外向里以顺时针的顺序依次打印出每一个数字.

func PrintMatrixClockwisely(mat [][]int) {
	row := len(mat)
	if row == 0 {
		return
	}
	col := len(mat[0])
	if col == 0 {
		return
	}

	// 按一圈打印
	var i, j int
	for ; j < col; j++ {
		fmt.Println(mat[i][j])
	}
	j--
	for i++; i < row; i++ {
		fmt.Println(mat[i][j])
	}
	i--
	for j--; i != 0 && j >= 0; j-- {
		fmt.Println(mat[i][j])
	}
	j++
	for i--; j != 0 && i >= 1; i-- {
		fmt.Println(mat[i][j])
	}

	if row > 2 && col > 2 {
		mat = mat[1 : row-2]
		for i := range mat {
			mat[i] = mat[i][1 : col-2]
		}
		PrintMatrixClockwisely(mat)
	}

}
