package data_structure_go

import (
	"fmt"
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	dll := NewDoubleLinkedList()
	dataHead := 20
	dll.AddFirst(dataHead)
	data := dll.getNodeByIndex(1)
	fmt.Println(data.Data.(int))
}
