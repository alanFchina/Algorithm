package 剑指offer

import (
	"math/rand"
)

// 题目: 输入一个链表的头结点, 从尾到头反过来打印

func printListBack(head *Node) []int {
	if head == nil {
		return []int{}
	}
	var r []int
	if head.next != nil {
		r = printListBack(head.next)
	}
	r = append(r, head.val)
	return r
}

type Node struct {
	val  int
	next *Node
}

func NewList(size int) (*Node, []int) {
	var (
		head *Node
		tail *Node
	)
	nums := make([]int, size, size)
	for i := 0; i < size; i++ {
		v := rand.Int()
		nums[size-i-1] = v
		n := Node{
			val:  v,
			next: nil,
		}
		if head == nil {
			head = &n
			tail = &n
		} else {
			tail.next = &n
			tail = &n
		}
	}
	return head, nums
}
