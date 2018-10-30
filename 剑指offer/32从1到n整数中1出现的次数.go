package 剑指offer

import (
	"math"
	"strconv"
)

func NumOf1Betwween1AndN(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 9 {
		return 1
	}
	return NumOf1(strconv.Itoa(n))
}

func NumOf1(n string) int {
	if len(n) == 1 {
		return 0
	}

	l := len(n)
	sum := 0
	if n[0] == 1 {
		n, _ := strconv.Atoi(n[1:])
		sum += n + 1
	} else {
		sum += pow(10, l-1)
	}
	a,_ := strconv.Atoi(string(n[0]))
	b := pow(10, l-2)
	sum += a * b

	return sum + NumOf1(n[1:])
}

func pow(i, j int) int {
	return int(math.Pow(float64(i), float64(j)))
}
