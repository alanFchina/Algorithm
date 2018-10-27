package base_algorithm

import (
	"fmt"
	"testing"
)

func TestAVLNode_Insert(t *testing.T) {
	input := []int{9, 7, 8, 10, 1, 2, 3, 4}
	root := NewAVLNode(0)
	fmt.Println("Insert:")
	for _, val := range input {
		root = root.Insert(val)
		printTree(root)
	}

	fmt.Println("Delete:")
	for _, val := range []int{2, 1} {
		root = root.Delete(val)
		printTree(root)
	}
}

type TreeNode interface {
	GetLeft() TreeNode
	GetRight() TreeNode
	GetValue() int
}

func (n *AVLNode) GetLeft() TreeNode {
	return n.Left
}

func (n *AVLNode) GetRight() TreeNode {
	return n.Right
}

func (n *AVLNode) GetValue() int {
	return n.Val
}

func printTree(root TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}
	fmt.Printf("%d\n", root.GetValue())
	Nodes := make([]TreeNode, 0)
	Nodes = append(Nodes, root)
	for len(Nodes) != 0 {
		NextNodes := make([]TreeNode, 0)
		for _, v := range Nodes {
			s1 := getStr(v.GetLeft())
			s2 := getStr(v.GetRight())
			fmt.Printf("[%s,%s] ", s1, s2)
			if !isNil(v.GetLeft()) {
				NextNodes = append(NextNodes, v.GetLeft())
			}
			if !isNil(v.GetRight()) {
				NextNodes = append(NextNodes, v.GetRight())
			}
		}
		fmt.Print("\n")
		Nodes = NextNodes
	}
}

func getStr(n TreeNode) string {
	isnil := isNil(n)
	if isnil {
		return "nil"
	} else {
		return fmt.Sprintf("%d", n.GetValue())
	}
}

func isNil(n TreeNode) (isNil bool) {
	isNil = false
	defer func() {
		if err := recover(); err != nil {
			isNil = true
		}
	}()
	n.GetLeft()
	return isNil
}
