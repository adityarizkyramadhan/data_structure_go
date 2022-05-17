package data_structure_go

import "testing"

func TestSLL(t *testing.T) {
	sll := &SLLList{}
	sll.AddDataLast(2)
	sll.Print()
}
