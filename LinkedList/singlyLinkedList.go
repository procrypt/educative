package singlyLinkedList

type List struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	Key  int
	Next *Node
}

func NewNode(value int) *Node {
	return &Node{Key:value}
}

func (l *List) Append(value int) {
	node := NewNode(value)
	if l.Size == 0 {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
	l.Size++
}

func (l *List) ListSize() int {
	return l.Size
}
