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
	nodel := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = nodel
		l.Tail = nodel
	} else {
		l.Tail.Next = nodel
		l.Tail = nodel
	}
}
