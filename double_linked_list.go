package data_structure_go

type NodeDL struct {
	Data int
	Next *NodeDL
	Prev *NodeDL
}

//Make double linked list
type NodeListDL struct {
	Head   *NodeDL
	Tail   *NodeDL
	length int
}

func (n *NodeListDL) IsEmpty() bool {
	return n.Head == nil
}

func (n *NodeListDL) Initialize() {
	n.Head = nil
	n.Tail = nil
	n.length = 0
}

// Belum selesai
