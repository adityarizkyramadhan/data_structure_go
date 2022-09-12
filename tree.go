package data_structure_go

import "fmt"

type node struct {
	Data  interface{}
	Right *node
	Left  *node
}

type Tree struct {
	Root *node
}

func (t *Tree) MakeTree(data interface{}) {
	newNode := &node{Data: data}
	t.Root = newNode
}

func (t *Tree) AddRight(data interface{}, parent *node) {
	newNode := &node{Data: data}
	if parent.Right == nil {
		parent.Right = newNode
	} else {
		fmt.Println("Node kanan sudah ada")
	}
}

func (t *Tree) AddLeft(data interface{}, parent *node) {
	newNode := &node{Data: data}
	if parent.Left == nil {
		parent.Left = newNode
	} else {
		fmt.Println("Node kiri sudah ada")
	}
}

func (t *Tree) DeleteAll(parent *node) {
	if parent != nil {
		t.DeleteAll(parent.Left)
		t.DeleteAll(parent.Right)
		parent = nil
	}
}

func (t *Tree) PreOrder(parent *node) {
	if parent != nil {
		fmt.Print(parent.Data.(string), " ")
		t.PreOrder(parent.Left)
		t.PreOrder(parent.Right)
	}
}

func (t *Tree) InOrder(parent *node) {
	if parent != nil {
		t.InOrder(parent.Left)
		fmt.Print(parent.Data.(string), " ")
		t.InOrder(parent.Right)
	}
}

func (t *Tree) PostOrder(parent *node) {
	if parent != nil {
		t.PostOrder(parent.Left)
		t.PostOrder(parent.Right)
		fmt.Print(parent.Data.(string), " ")
	}
}
