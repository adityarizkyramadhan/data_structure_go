package data_structure_go

import (
	"fmt"
	"testing"
)

type mahasiswa struct {
	Name  string
	Grade int
}

func TestSLL(t *testing.T) {
	sll := NewSingleLinkedList()
	sll.AddDataLast(mahasiswa{Name: "Aditya", Grade: 99})
	data := sll.RemoveDataByIndex(0)
	fmt.Println(data.(mahasiswa).Name)
	fmt.Println(sll.IsEmpty())
}
