package BinaryTreeImplementation

import "fmt"

type BinaryTree struct {
	Root *Node
}

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (b *BinaryTree) Insert(value int) {
	if b.Root == nil {
		b.Root = &Node{Key:value}
	} else {
		b.Root.insert(value)
	}
}

func (n *Node) insert(value int) {
	// Left Insert
	if n.Key > value {
		if n.Left == nil {
			n.Left = &Node{Key: value}
		} else {
			n.Left.insert(value)
		}

		// Right Insert
	} else if n.Key < value {
		if n.Right == nil {
			n.Right = &Node{Key:value}
		} else {
			n.Right.insert(value)
		}
	}
}

func (b *BinaryTree) Search(value int) bool {
	if b.Root.Key == value {
		return true
	} else {
		return b.Root.search(value)
	}
	return false
}

func (n *Node) search(value int) bool {
	if n.Key == value {
		return true
	// Search Left
	} else if n.Key > value {
		if n.Left != nil {
			return n.Left.search(value)
		}
	// Search Right
	} else {
		if n.Right != nil {
			return n.Right.search(value)
		}
	}
	return false
}

func (b *BinaryTree) InOrder() {
	if b != nil {
		b.Root.inorder()
	} else {
		return
	}
}

func (n *Node) inorder() {
	if n.Left != nil {
		n.Left.inorder()
	}
	fmt.Println(n.Key)
	if n.Right != nil {
		n.Right.inorder()
	}
}

func (b *BinaryTree) PreOrder() {
	if b != nil {
		b.Root.preOrder()
	} else {
		return
	}
}

func (n *Node) preOrder() {
	fmt.Println(n.Key)
	if n.Left != nil {
		n.Left.preOrder()
	}
	if n.Right != nil {
		n.Right.preOrder()
	}
}

func (b *BinaryTree) PostOrder() {
	if b != nil {
		b.Root.postOrder()
	} else {
		return
	}
}

func (n *Node) postOrder() {
	if n.Left != nil {
		n.Left.postOrder()
	}
	if n.Right != nil {
		n.Right.postOrder()
	}
	fmt.Println(n.Key)
}
