package data_structure_go

import "fmt"

type nodeSLL struct {
	Data interface{}
	Next *nodeSLL
}

type sLList struct {
	Head   *nodeSLL
	length int
}

func NewSingleLinkedList() *sLList {
	return &sLList{}
}

func (n *sLList) AddDataLast(data interface{}) {
	nextNode := nodeSLL{Data: data}
	if n.Head == nil {
		n.Head = &nextNode
		return
	}
	currentNodeSLL := n.Head
	for currentNodeSLL != nil {
		currentNodeSLL = currentNodeSLL.Next
		//add data to last
		if currentNodeSLL.Next == nil {
			currentNodeSLL.Next = &nextNode
			n.length++
			break
		}
	}
}

func (n *sLList) AddDataFirst(data interface{}) {
	newNodeSLL := &nodeSLL{Data: data}
	if n.Head != nil {
		newNodeSLL.Next = n.Head
		n.Head = newNodeSLL
		n.length++
	} else {
		temp := n.Head
		newNodeSLL.Next = temp
		n.Head = newNodeSLL
	}
}

func (n *sLList) InsertDataAfter(index int, data interface{}) {
	if n.Head == nil {
		fmt.Println("Error, Null LinkedList")
		return
	}
	newNodeSLL := &nodeSLL{Data: data}
	currentNodeSLL := n.Head
	for i := 0; i < index; i++ {
		if currentNodeSLL.Next == nil {
			fmt.Println("Error, Index out of range")
			return
		}
		currentNodeSLL = currentNodeSLL.Next
	}
	newNodeSLL.Next = currentNodeSLL.Next
	currentNodeSLL.Next = newNodeSLL
	n.length++
}

func (n *sLList) InsertDataBefore(index int, data interface{}) {
	if index == 0 {
		fmt.Println("Error, you can't insert data before zero")
		return
	}
	if n.Head == nil {
		fmt.Println("Error, Null LinkedList")
		return
	}
	newNode := &nodeSLL{Data: data}
	currentNode := n.Head
	for i := 0; i < index-1; i++ {
		if currentNode.Next == nil {
			fmt.Println("Error, index out of range")
			return
		}
		currentNode = currentNode.Next
	}
	newNode.Next = currentNode.Next
	currentNode.Next = newNode

}

//Not used because data type is interface{} so i don't give method print in this case
// func (n *sLList) Print() {
// 	if n.Head == nil {
// 		return
// 	}
// 	currentNodeSLL := n.Head
// 	for currentNodeSLL != nil {
// 		fmt.Println(currentNodeSLL.Data)
// 		currentNodeSLL = currentNodeSLL.Next
// 	}
// 	fmt.Println()
// }

func (n *sLList) RemoveDataByIndex(index int) interface{} {
	if n.Head == nil {
		fmt.Println("Empty linked list")
		return nil
	}
	var temp *nodeSLL
	currNode := n.Head
	if index == 0 {
		temp = n.Head
		n.Head = nil
		n.length--
		return temp.Data
	}
	for i := 0; i < index-1; i++ {
		if currNode.Next == nil {
			fmt.Println("Error, index is out of range")
			return nil
		}
		currNode = currNode.Next
	}
	temp = currNode.Next
	currNode.Next = currNode.Next.Next
	n.length--
	return temp.Data
}

func (n *sLList) GetDataByIndex(index int) interface{} {
	if n.Head == nil {
		fmt.Println("Empty linked list")
		return nil
	}
	currNode := n.Head
	if index == 0 {
		return currNode.Data
	}
	for i := 0; i < index; i++ {
		if currNode.Next == nil {
			fmt.Println("Error, index is out of range")
			return nil
		}
		currNode = currNode.Next
	}
	return currNode.Data
}

func (n *sLList) Len() int {
	return n.length
}

func (n *sLList) IsEmpty() bool {
	return n.Head == nil
}
