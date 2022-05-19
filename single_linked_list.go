package data_structure_go

import "fmt"

type nodeSll struct {
	Data interface{}
	Next *nodeSll
}

type singleLinkedList struct {
	Head *nodeSll
}

func NewSingleLinkedList() *singleLinkedList {
	return &singleLinkedList{}
}

func (n *singleLinkedList) AddDataLast(data interface{}) {
	nextNode := nodeSll{Data: data}
	if n.Head == nil {
		n.Head = &nextNode
		return
	}
	currNode := n.Head
	for currNode.Next != nil {
		currNode = currNode.Next
	}
	currNode.Next = &nextNode
}

func (n *singleLinkedList) AddDataFirst(data interface{}) {
	newnodeSll := &nodeSll{Data: data}
	if n.Head != nil {
		newnodeSll.Next = n.Head
		n.Head = newnodeSll

	} else {
		temp := n.Head
		newnodeSll.Next = temp
		n.Head = newnodeSll

	}
}

func (n *singleLinkedList) InsertDataAfter(index int, data interface{}) {
	if n.Head == nil {
		fmt.Println("Error, Null LinkedList")
		return
	}
	newnodeSll := &nodeSll{Data: data}
	currentnodeSll := n.Head
	for i := 0; i < index; i++ {
		if currentnodeSll.Next == nil {
			fmt.Println("Error, Index out of range")
			return
		}
		currentnodeSll = currentnodeSll.Next
	}
	newnodeSll.Next = currentnodeSll.Next
	currentnodeSll.Next = newnodeSll
}

func (n *singleLinkedList) InsertDataBefore(index int, data interface{}) {
	if index == 0 {
		fmt.Println("Error, you can't insert data before zero")
		return
	}
	if n.Head == nil {
		fmt.Println("Error, Null LinkedList")
		return
	}
	newNode := &nodeSll{Data: data}
	currentNode := n.Head
	for i := 0; i < index-1; i++ {
		if currentNode.Next == nil {
			fmt.Println("Error, index out of range")
			return
		}
		currentNode = currentNode.Next
	}
	newNode.Next = currentNode.Next
	currentNode.Next = newNode.Next
}

//Not used because data type is interface{} so i don't give method print in this case
// func (n *singleLinkedList) Print() {
// 	if n.Head == nil {
// 		return
// 	}
// 	currentnodeSll := n.Head
// 	for currentnodeSll != nil {
// 		fmt.Println(currentnodeSll.Data)
// 		currentnodeSll = currentnodeSll.Next
// 	}
// 	fmt.Println()
// }

func (n *singleLinkedList) RemoveDataByIndex(index int) interface{} {
	if n.Head == nil {
		fmt.Println("Empty linked list")
		return nil
	}
	var temp *nodeSll
	currNode := n.Head
	if index == 0 {
		temp = n.Head
		n.Head = nil
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
	return temp.Data
}

func (n *singleLinkedList) GetDataByIndex(index int) interface{} {
	if n.Head == nil {
		fmt.Println("Empty linked list")
		return nil
	}
	currNode := n.Head
	if index == 0 {
		return n.Head.Data
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

func (n *singleLinkedList) Len() int {
	// TODO: count elements in linked list
	curr := n.Head
	count := 0
	for curr != nil {
		count++
		curr = curr.Next
	}
	return count
}

func (n *singleLinkedList) IsEmpty() bool {
	return n.Len() == 0
}
