package data_structure_go

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type Queue struct {
	First *Node
	Tail  *Node
	Next  *Node
}

func (q *Queue) Print() {
	curr := q.First
	for curr != nil {
		fmt.Println(curr.Data)
		curr = curr.Next
	}
}

func (q *Queue) Enqueue(data string) {
	node := &Node{
		Data: data,
	}
	if q.First == nil && q.Tail == nil {
		q.First = node
		q.Tail = node
	}
	q.Tail.Next = node
	q.Tail = node
}

func (q *Queue) Dequeue() *Node {
	if q.First == nil {
		panic("Queue empty")
	}
	result := q.First
	q.First = q.First.Next
	return result
}
