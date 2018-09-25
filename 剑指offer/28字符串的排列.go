package 剑指offer

import "fmt"

// 题目: 输入一个字符串, 打印出该字符串中字符的所有排列

func Permutation(s string) {
	if s == "" {
		return
	}
	permutation([]byte(s), 0)
	// permutationLoop([]byte(s))
}

func permutation(c []byte, start int) {
	if start == len(c)-1 {
		fmt.Printf("%c\n", c)
	} else {
		for i := start; i < len(c); i++ {
			c[start], c[i] = c[i], c[start]
			permutation(c, start+1)
			c[start], c[i] = c[i], c[start] // 每次交换之后都交换回去, 所以经过两次交换之后还是原来的顺序
		}
	}
}

func permutationLoop(c []byte) {
	for i := 0; i < len(c); i++ {
		for j := i; j < len(c); j++ {
			for k := j; k < len(c); k++ {
				c[j], c[k] = c[k], c[j]
				fmt.Printf("%c\n", c)
				c[j], c[k] = c[k], c[j]
			}
		}
	}
}
