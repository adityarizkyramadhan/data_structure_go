package data_structure_go

type nodeDL struct {
	Data interface{}
	Next *nodeDL
	Prev *nodeDL
}

//Make double linked list
type nodeListDL struct {
	Head   *nodeDL
	Tail   *nodeDL
	length int
}

func (n *nodeListDL) IsEmpty() bool {
	return n.Head == nil && n.length == 0
}

func NewDoubleLinkedList() *nodeListDL {
	return &nodeListDL{
		Head:   nil,
		Tail:   nil,
		length: 0,
	}
}

func (n *nodeListDL) AddFirst(data interface{}) {
	newNode := &nodeDL{Data: data}
	if n.Head == nil {
		n.Head = newNode
		n.Tail = newNode
		n.length++
		return
	}
	newNode.Next = n.Head
	newNode.Prev = nil
	n.Head = newNode
	n.length++
}

func (n *nodeListDL) AddLast(data interface{}) {
	newNode := &nodeDL{Data: data}
	if n.Head == nil && n.Tail == nil {
		n.Head = newNode
		n.Tail = newNode
		n.length++
		return
	}
	n.Tail.Next = newNode
	newNode.Prev = n.Tail
	n.Tail = newNode
	n.length++
}

func (n *nodeListDL) GetDataByIndex(index int) interface{} {
	return n.getNodeByIndex(index).Data
}

func (n *nodeListDL) getNodeByIndex(index int) *nodeDL {
	temp := n.Head
	for i := 0; i < index; i++ {
		if i == index-1 {
			return temp
		}
		temp = temp.Next
	}
	return nil
}

func (n *nodeListDL) DeleteData(index int) bool {
	node := n.getNodeByIndex(index)
	if node == n.Head {
		n.Head = n.Head.Next
		n.Head.Prev = nil
		n.length--
		return !(node == n.Head)
	}
	return false

}
