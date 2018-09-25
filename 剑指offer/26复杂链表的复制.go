package 剑指offer

type ComplexListNode struct {
	Val           int
	Next, Sibling *ComplexListNode
}

func CloneNodes(head *ComplexListNode) *ComplexListNode {
	if head == nil {
		return nil
	}
	this := head
	// 复制结点
	for this != nil {
		newNode := &ComplexListNode{
			Val:  this.Val,
			Next: this.Next,
		}
		this.Next = newNode
		this = newNode.Next
	}

	// 复制Sibling
	this = head
	for this != nil {
		this.Next.Sibling = this.Sibling.Next
		this = this.Next.Next
	}

	// 分离
	this1 := head
	this2 := this1.Next
	newHead := this2
	for this2 != nil {
		this1.Next = this2.Next
		this2.Next = this2.Next.Next
		this1 = this1.Next
		this2 = this2.Next
	}
	this1 = nil
	return newHead
}
