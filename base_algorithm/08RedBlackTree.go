package base_algorithm

type RBTree struct {
	root *RBNode
	nil  *RBNode
}

type RBNode struct {
	Left, Right, Parent *RBNode
	Value               int
	Color               string
}

func NewRBTree() RBTree {
	return RBTree{
		root: nil,
		nil: &RBNode{
			Left:   nil,
			Right:  nil,
			Parent: nil,
			Value:  0,
			Color:  black,
		},
	}
}

const red, black = "RED", "BLACK"

func (T *RBTree) Search(key int) *RBNode {
	if T.root == nil {
		return nil
	}
	this := T.root
	for this != T.nil {
		if this.Value == key {
			return this
		}
		if key < this.Value {
			this = this.Left
		} else {
			this = this.Right
		}
	}
	return nil
}

func (T *RBTree) Insert(key int) {
	newNode := &RBNode{
		Left:   T.nil,
		Right:  T.nil,
		Parent: nil,
		Color:  red,
		Value:  key,
	}
	if T.root == nil {
		T.root = newNode
		newNode.Parent = T.nil
		newNode.Color = black
		return
	}

	this := T.root
	next := this
	for next != T.nil {
		this = next
		if key < this.Value {
			next = this.Left
		} else {
			next = this.Right
		}
	}
	newNode.Parent = this
	if key < this.Value {
		this.Left = newNode
	} else {
		this.Right = newNode
	}
	T.insertFixUp(newNode)
}

func (T *RBTree) insertFixUp(z *RBNode) {
	for z.Parent.Color == red {
		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if y.Color == red {
				z = z.Parent.Parent
				z.Color = red
				z.Left.Color = black
				z.Right.Color = black
			} else {
				if z == z.Parent.Right {
					z = z.Parent
					T.leftRotate(z)
				}
				T.rightRotate(z.Parent.Parent)
				z.Parent.Color = black
				z.Parent.Right.Color = red
			}
		} else {
			y := z.Parent.Parent.Left
			if y.Color == red {
				z = z.Parent.Parent
				z.Color = red
				z.Left.Color = black
				z.Right.Color = black
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					T.rightRotate(z)
				}
				T.leftRotate(z.Parent.Parent)
				z.Parent.Color = black
				z.Parent.Left.Color = red
			}
		}
		T.root.Color = black
	}
}

//      y                    x
//   x   c                 a   y
// a   b    <--left          b   c
func (T *RBTree) leftRotate(x *RBNode) {
	y := x.Right
	x.Right = y.Left
	if y.Left != T.nil {
		y.Left.Parent = x
	}
	y.Left = x
	y.Parent = x.Parent
	if T.root == x {
		T.root = y
	} else if x.Parent.Left == x {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	x.Parent = y
}

//      y                    x
//   x   c   -->right      a   y
// a   b                     b   c
func (T *RBTree) rightRotate(y *RBNode) {
	x := y.Left
	y.Left = x.Right
	if x.Right != T.nil {
		x.Right.Parent = y
	}
	x.Right = y
	if T.root == y {
		T.root = x
	} else if y.Parent.Left == y {
		y.Parent.Left = x
	} else {
		y.Parent.Right = x
	}
	y.Parent = x
}

func (T *RBTree) Delete(n *RBNode) {
	row_color := n.Color
	rep_node := T.nil
	if n.Right == T.nil && n.Left == T.nil {
		T.replace(n, T.nil)
	} else if n.Right == T.nil {
		rep_node = n.Left
		T.replace(n, n.Left)
	} else if n.Left == T.nil {
		T.replace(n, n.Right)
		rep_node = n.Right
	} else {
		minNode := n.Right.Min()
		if minNode != n.Right {
			c := row_color
			row_color = minNode.Color
			minNode.Color = c
			rep_node = minNode.Right
			minNode.Parent.Left = minNode.Right
			minNode.Right = n.Right
			n.Right.Parent = minNode
		}
		T.replace(n, minNode)
		minNode.Left = n.Left
	}
	if row_color == black {
		T.deleteFixUp(rep_node)
	}
}

func (T *RBTree) deleteFixUp(x *RBNode) {
	for T.root != x && x.Color == black {
		if x == x.Parent.Left {
			if x.Parent.Right.Color == red {
				x = x.Parent
				x.Color = red
				x.Right.Color = black
			} else if x.Parent.Right.Color == black {
				r := x.Parent.Right
				if r.Right.Color == black {
					T.rightRotate(r)
					r.Parent.Color = black
					r.Color = red
				}
				T.leftRotate(x.Parent)
				r.Color = x.Parent.Color
				x.Parent.Color = black
				r.Right.Color = black
			}
		} else {
			if x.Parent.Left.Color == red {
				x = x.Parent
				x.Color = red
				x.Left.Color = black
			} else if x.Parent.Left.Color == black {
				r := x.Parent.Left
				if r.Left.Color == black {
					T.leftRotate(r)
					r.Parent.Color = black
					r.Color = red
				}
				T.rightRotate(x.Parent)
				r.Color = x.Parent.Color
				x.Parent.Color = black
				r.Left.Color = black
			}
		}

	}

	x.Color = black
}

func (T *RBTree) replace(n1 *RBNode, n2 *RBNode) {
	if T.root == n1 {
		T.root = n2
	} else {
		if n1.Parent.Left == n1 {
			n1.Parent.Left = n2
		} else {
			n1.Parent.Right = n2
		}
	}
	n1.Parent = n1.Parent
}

func (T *RBTree) Min() *RBNode {
	if T.root == nil {
		return nil
	}
	this := T.root
	for this.Left != T.nil {
		this = this.Left
	}
	return this
}

func (T *RBTree) Max() *RBNode {
	if T.root == nil {
		return nil
	}
	this := T.root
	for this.Right != T.nil {
		this = this.Right
	}
	return this
}

func (N *RBNode) Min() *RBNode {
	this := N
	for this.Left != nil {
		this = this.Left
	}
	return this
}
