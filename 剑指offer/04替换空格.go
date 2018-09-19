package 剑指offer

import "bytes"

// 题目: 请实现一个函数,把字符串中的每个空格替换成"%20".
// 例如: 输入"We are happy", 则输出"We%20are%20happy".

// 这个问题有个关键点在于原字符串是否有足够的空间, 这里假设有
// 则需要开辟一段空间来存储结果

func replaceSpace(str []byte, l int) {
	space := []byte(" ")
	num := bytes.Count(str, space)
	if num == 0 {
		return
	}
	l--
	for i := l + 2*num; l >= 0; l-- {
		if str[l] != ' ' {
			str[i] = str[l]
			i--
		} else {
			str[i] = '0'
			i--
			str[i] = '2'
			i--
			str[i] = '%'
			i--
		}
	}
}
