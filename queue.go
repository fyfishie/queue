package queue

type Node[D any] struct {
	pre  *Node[D]
	Data D
	next *Node[D]
}
type Queue[D any] struct {
	Head   *Node[D]
	Tail   *Node[D]
	Length int
}

func (q *Queue[D]) Push(Data D) {
	newNode := Node[D]{Data: Data}
	if q.Length == 0 {
		q.Head = &newNode
		q.Tail = &newNode
		q.Length++
	} else {
		q.Tail.next = &newNode
		q.Tail = &newNode
		q.Length++
	}
}
func (q *Queue[D]) Pop() *D {
	if q.Length == 0 {
		return nil
	} else {
		ret := q.Head
		q.Head = ret.next
		q.Length--
		return &ret.Data
	}
}
func NewQueue[D any]() *Queue[D] {
	q := Queue[D]{
		Head:   nil,
		Tail:   nil,
		Length: 0,
	}
	return &q
}
