package data_structure_go

import "fmt"

type nodeSLL struct {
	Data int
	Next *nodeSLL
}

type SLLList struct {
	Head   *nodeSLL
	length int
}

func (n *SLLList) AddDataLast(data int) {
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

func (n *SLLList) AddDataFirst(data int) {
	newNodeSLL := nodeSLL{Data: data}
	if n.Head != nil {
		newNodeSLL.Next = n.Head
		n.Head = &newNodeSLL
		n.length++
	} else {
		temp := n.Head
		n.Head.Next = temp
		n.Head = &newNodeSLL
		n.length++
	}
}

func (n *SLLList) InsertDataAfter(data int, NodeSLLBefore nodeSLL) {
	if n.Head == nil {
		fmt.Println("Error, Null LinkedList")
	}
	//Masukkan NodeSLL yg baru
	newNodeSLL := nodeSLL{Data: data}
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

func (n *SLLList) InsertDataBefore(data int, NodeSLLAfter nodeSLL) {
	if n.Head == nil {
		fmt.Println("Error, Null LinkedList")
	}
	//Masukkan NodeSLL yg baru
	newNodeSLL := nodeSLL{Data: data}
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

func (n *SLLList) Print() {
	if n.Head == nil {
		return
	}
	currentNodeSLL := n.Head
	for currentNodeSLL != nil {
		fmt.Println(currentNodeSLL.Data)
		currentNodeSLL = currentNodeSLL.Next
	}
	fmt.Println()

}

func (n *SLLList) RemoveDataByValue(data int) {
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

func (n *SLLList) Len() int {
	return n.length
}

func (n *SLLList) Initialize() {
	n.Head = nil
}

func (n *SLLList) IsEmpty() bool {
	return n.Head == nil
}
