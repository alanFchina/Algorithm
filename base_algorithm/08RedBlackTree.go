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
