package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	if l == nil {
		return
	}

	node := &NodeL{data, nil}
	if l.Head == nil && l.Tail == nil {
		l.Head = node
		l.Tail = l.Head
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}
