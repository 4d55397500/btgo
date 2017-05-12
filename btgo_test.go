package btgo


import (
	"testing"
)


func TestNode_IsBinarySearchTree(t *testing.T) {
	n := SampleTree()
	if !n.IsBinarySearchTree() {
		t.Error("Expected sample tree to be a binary search tree")
	}
}

func TestNode_Lookup(t *testing.T) {
	n := SampleTree()
	oneExists := n.Lookup(1)
	twoExists := n.Lookup(2)
	threeExists := n.Lookup(3)
	if !oneExists {
		t.Error("Expeted node with data 1; not found")
	}
	if !twoExists {
		t.Error("Expected node with data 2; not found")
	}
	if threeExists {
		t.Error("Found unexpected node with data 3")
	}
}


func TestNode_Insert(t *testing.T) {
	n := SampleTree().Insert(-1)
	if n.left.data != -1 {
		t.Error("Expected inserted left child node with data -1, got ", n.left.data)
	}
}


func TestNode_MaxDepth(t *testing.T) {
	n := SampleTree()
	if n.MaxDepth() != 2 {
		t.Error("Expected max depth of 2, got ", n.MaxDepth())
	}
}

func TestNode_PrintTree(t *testing.T) {
	n := SampleTree()
	n.PrintTree()
}


func SampleTree() *node {
	n0 := node{1, nil, nil}
	n1 := node{2, nil, nil}
	return n0.Insert(n1.data)
}