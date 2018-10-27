package base_algorithm

type Node struct {
	Val    int
	Parent *Node
	Left   *Node
	Right  *Node
}

func NewBST() *Node {
	return nil
}

// 插入
func (root *Node) Insert(v int) *Node {
	if root == nil {
		return &Node{v, nil, nil, nil}
	}

	if v < root.Val {
		newLeft := root.Left.Insert(v)
		newLeft.Parent = root
		root.Left = newLeft
		return root
	}
	newRight := root.Right.Insert(v)
	newRight.Parent = root
	root.Right = newRight
	return root
}

// 删除
func (root *Node) Delete(n *Node) {
	if n.Left == nil {
		Replace(n.Right, n)
		return
	}
	if n.Right == nil {
		Replace(n.Left, n)
		return
	}

	next := Next(n)
	next.Parent.Left = next.Right
	next.Right.Parent = next.Parent

	Replace(next, n)
	next.Left = n.Left
	next.Right = n.Right
}

// 查找
func (root *Node) Search(v int) *Node {
	if root == nil {
		return nil
	}
	if v == root.Val {
		return root
	}
	if v < root.Val {
		return root.Left.Search(v)
	} else {
		return root.Right.Search(v)
	}
}

func Next(n *Node) *Node {
	this := n
	for this.Left != nil {
		this = this.Left
	}
	return this
}

func Replace(n1 *Node, n2 *Node) {
	if n2.Parent.Left == n2 {
		n2.Parent.Left = n1
	} else {
		n2.Parent.Right = n1
	}
	n1.Parent = n2.Parent
}
