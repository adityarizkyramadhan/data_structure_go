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
	// fmt.Println(sll.IsEmpty())
	sll.AddDataLast(mahasiswa{Name: "Aditya", Grade: 99})
	sll.AddDataLast(mahasiswa{Name: "Rizky", Grade: 99})
	sll.AddDataLast(mahasiswa{Name: "Ramadhan", Grade: 99})
	fmt.Println(sll.Len())
	fmt.Println(sll.GetDataByIndex(1).(mahasiswa).Name)
	for i := 0; i < sll.Len(); i++ {
		fmt.Println(sll.GetDataByIndex(i).(mahasiswa).Name)
	}
}
