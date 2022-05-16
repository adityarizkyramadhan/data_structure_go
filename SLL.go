package data_structure_go

import "fmt"

type NodeSLL struct {
	Data int
	Next *NodeSLL
}

type NodeSLLList struct {
	Head   *NodeSLL
	length int
}

func (n *NodeSLLList) AddDataLast(data int) {
	nextNodeSLL := NodeSLL{Data: data}
	if n.Head == nil {
		n.Head = &nextNodeSLL
	}
	currentNodeSLL := n.Head
	for currentNodeSLL != nil {
		if currentNodeSLL.Next == nil {
			currentNodeSLL.Next = &nextNodeSLL
			n.length++
			break
		}
		currentNodeSLL = currentNodeSLL.Next
	}
}

func (n *NodeSLLList) AddDataFirst(data int) {
	newNodeSLL := NodeSLL{Data: data}
	if n.Head != nil {
		newNodeSLL.Next = n.Head
		n.Head = &newNodeSLL
		n.length++
	} else {
		n.Head = &newNodeSLL
		n.length++
	}
}

func (n *NodeSLLList) InsertDataAfter(data int, NodeSLLBefore NodeSLL) {
	if n.Head == nil {
		fmt.Println("Error, Null LinkedList")
	}
	//Masukkan NodeSLL yg baru
	newNodeSLL := NodeSLL{Data: data}
	//Bikin LinkedList
	currentNodeSLL := n.Head
	//Loop sampai nil
	for currentNodeSLL != nil {
		//Cek data yang sama
		if NodeSLLBefore.Data == currentNodeSLL.Data {
			//NodeSLL selanjutnya yang lama diputus terus disambung ke NodeSLL yg baru
			newNodeSLL.Next = currentNodeSLL.Next
			//NodeSLL lanjutan disambungkan ke NodeSLL yg baru
			currentNodeSLL.Next = &newNodeSLL
			//Tambah Panjang
			n.length++
			break
		}
		//ganti ganti data terus
		currentNodeSLL = currentNodeSLL.Next
	}
}

func (n *NodeSLLList) InsertDataBefore(data int, NodeSLLAfter NodeSLL) {
	if n.Head == nil {
		fmt.Println("Error, Null LinkedList")
	}
	//Masukkan NodeSLL yg baru
	newNodeSLL := NodeSLL{Data: data}
	//Masukkan si kepala linkedlist
	currentNodeSLL := n.Head
	//loop smpe nil
	for currentNodeSLL != nil {
		// Mengecek data yang next buat dipisah
		if currentNodeSLL.Next.Data == NodeSLLAfter.Data {
			//NodeSLL yg diafterkan dipisah dulu
			tempNodeSLL := currentNodeSLL.Next
			//Tempat yg data awal, diganti sama newNodeSLL
			currentNodeSLL.Next = &newNodeSLL
			//NodeSLL baru yg sudah terpasang, disambunglagi sama yg dipisah tadi
			currentNodeSLL.Next.Next = tempNodeSLL
			n.length++
			break
		}
		currentNodeSLL = currentNodeSLL.Next
	}

}

func (n *NodeSLLList) Print() {
	if n.Head == nil {
		return
	}
	currentNodeSLL := n.Head
	for currentNodeSLL != nil {
		fmt.Print(currentNodeSLL.Data, "->")
		currentNodeSLL = currentNodeSLL.Next
	}
	fmt.Println()

}

func (n *NodeSLLList) RemoveDataByValue(data int) {
	if n.Head == nil {
		fmt.Println("List kosong")
		return
	}
	if n.Head.Data == data {
		n.Head = n.Head.Next
		n.length--
		return
	}
	previousNodeSLL := n.Head
	for previousNodeSLL.Next.Data != data {
		if previousNodeSLL.Next.Next == nil {
			return
		}
		previousNodeSLL = previousNodeSLL.Next
	}
	previousNodeSLL.Next = previousNodeSLL.Next.Next
	n.length--
}

func (n *NodeSLLList) Len() int {
	return n.length
}

func (n *NodeSLLList) Initialize() {
	n.Head = nil
}

func (n *NodeSLLList) IsEmpty() bool {
	return n.Head == nil
}
