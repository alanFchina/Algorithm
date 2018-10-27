package base_algorithm

import (
	"log"
)

type AVLNode struct {
	Left, Right *AVLNode
	Val, High   int
}

func NewAVLNode(v int) *AVLNode {
	return &AVLNode{
		Left:  nil,
		Right: nil,
		Val:   v,
		High:  1,
	}
}

func (root *AVLNode) Search(key int) *AVLNode {
	if root == nil {
		return nil
	}
	if key == root.Val {
		return root
	} else if key < root.Val {
		return root.Left.Search(key)
	} else {
		return root.Right.Search(key)
	}
}

func (root *AVLNode) Insert(v int) *AVLNode {
	if root == nil {
		root = NewAVLNode(v)
	} else if v < root.Val {
		root.Left = root.Left.Insert(v)
		if root.Left.getHigh()-root.Right.getHigh() == 2 {
			if v < root.Left.Val {
				root = root.left_left_rotate()
				root.Right.High = max(root.Right.Right.getHigh(), root.Right.Left.getHigh()) + 1
			} else {
				root = root.left_right_rotate()
				root.Right.High = max(root.Right.Right.getHigh(), root.Right.Left.getHigh()) + 1
				root.Left.High = max(root.Left.Right.getHigh(), root.Left.Right.getHigh()) + 1
			}
		}
	} else if v > root.Val {
		root.Right = root.Right.Insert(v)
		if root.Right.getHigh()-root.Left.getHigh() == 2 {
			if v > root.Right.Val {
				root = root.right_right_rotate()
				root.Left.High = max(root.Left.Right.getHigh(), root.Left.Right.getHigh()) + 1
			} else {
				root = root.right_left_rotate()
				root.Left.High = max(root.Left.Right.getHigh(), root.Left.Right.getHigh()) + 1
				root.Right.High = max(root.Right.Right.getHigh(), root.Right.Left.getHigh()) + 1
			}
		}
	} else {
		// 等于的情况
		log.Printf("value: %d has already existed!")
	}
	root.High = max(root.Right.getHigh(), root.Left.getHigh()) + 1
	return root
}

func (root *AVLNode) Delete(key int) *AVLNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = root.Left.Delete(key)
		if root.Right.getHigh()-root.Left.getHigh() == 2 {
			if root.Right.Right.getHigh() >= root.Right.Left.getHigh() {
				root = root.right_right_rotate()
				root.Left.High = max(root.Left.Right.getHigh(), root.Left.Right.getHigh()) + 1
			} else {
				root = root.right_left_rotate()
				root.Left.High = max(root.Left.Right.getHigh(), root.Left.Right.getHigh()) + 1
				root.Right.High = max(root.Right.Right.getHigh(), root.Right.Right.getHigh()) + 1
			}
		}
	} else if key > root.Val {
		root.Right = root.Right.Delete(key)
		if root.Left.getHigh()-root.Right.getHigh() == 2 {
			if root.Left.Left.getHigh() >= root.Left.Right.getHigh() {
				root = root.left_left_rotate()
				root.Right.High = max(root.Right.Right.getHigh(), root.Right.Right.getHigh()) + 1
			} else {
				root = root.left_right_rotate()
				root.Left.High = max(root.Left.Right.getHigh(), root.Left.Right.getHigh()) + 1
				root.Right.High = max(root.Right.Right.getHigh(), root.Right.Right.getHigh()) + 1
			}
		}
	} else {
		// 找到需要删除的结点
		if root.Right == nil {
			return root.Left
		}
		if root.Left == nil {
			return root.Right
		}
		next := root.Right.getMinNode()
		root.Right = root.Right.Delete(next.Val)
		next.Left = root.Left
		next.Right = root.Right
		root = next
	}
	root.High = max(root.Right.getHigh(), root.Left.getHigh()) + 1
	return root
}


func (root *AVLNode) getMinNode() *AVLNode {
	if root.Left == nil {
		return root
	}
	return root.Left.getMinNode()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (root *AVLNode) getHigh() int {
	if root == nil {
		return 0
	}
	return root.High
}

//      X                  Y
//    Y      right-->        X
//     N     <--left       N
func (n *AVLNode) left_rotate() *AVLNode {
	Y := n
	X := Y.Right
	N := X.Left

	X.Left = Y
	Y.Right = N
	return X
}

func (n *AVLNode) right_rotate() *AVLNode {
	X := n
	Y := X.Left
	N := Y.Right

	Y.Right = X
	X.Left = N
	return Y
}

//      X             Y
//    Y     -->     Z   X
//  Z   N             N
func (n *AVLNode) left_left_rotate() *AVLNode {
	return n.right_rotate()
}

//         X            X            Z
//      Y       -->   Z       -->  Y   X
//        Z         Y
func (n *AVLNode) left_right_rotate() *AVLNode {
	X := n
	Y := X.Left

	X.Left = Y.left_rotate()
	return X.right_rotate()
}

//     X                       Y
//        Y          -->     X   Z
//           Z
func (n *AVLNode) right_right_rotate() *AVLNode {
	return n.left_rotate()
}

//     X                 X              Z
//       Y      -->        Z     -->  X   Y
//     Z                     Y
func (n *AVLNode) right_left_rotate() *AVLNode {
	X := n
	Y := X.Right
	X.Right = Y.right_rotate()
	return X.left_rotate()
}
