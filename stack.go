package data_structure_go

import "fmt"

type node_stack struct {
	Data string
	Next *node_stack
}

type Stack struct {
	Top *node_stack
}

func (s *Stack) IsEmpty() bool {
	return s.Top == nil
}

func (s *Stack) CountElement() int {
	temp := s.Top
	jumlah := 0
	if temp == nil {
		return jumlah
	}
	for temp != nil {
		jumlah++
		temp = temp.Next
	}
	return jumlah
}

func (s *Stack) CreateEmpty() {
	s.Top = nil
}

func (s *Stack) Push(input string) {
	node_stackIn := &node_stack{Data: input}
	if s.Top == nil {
		s.Top = node_stackIn
		return
	}
	node_stackIn.Next = s.Top
	s.Top = node_stackIn
}

func (s *Stack) Pop() {
	temp := s.Top
	temp2 := s.Top.Data
	s.Top = temp.Next
	fmt.Println("Data dibuang ->", temp2)
}

func (s *Stack) Print() {
	temp := s.Top
	for temp != nil {
		fmt.Println("Data : ", temp.Data)
		temp = temp.Next
	}
}
