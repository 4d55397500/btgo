// Package btgo provides common operations over
// binary search trees
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
		t.Error("Expected Node with Data 1; not found")
	}
	if !twoExists {
		t.Error("Expected Node with Data 2; not found")
	}
	if threeExists {
		t.Error("Found unexpected Node with Data 3")
	}
}

func TestNode_Insert(t *testing.T) {
	n := SampleTree().Insert(-1)
	if n.Left.Data != -1 {
		t.Error("Expected inserted Left child Node with Data -1, got ", n.Left.Data)
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

func SampleTree() *Node {
	n0 := Node{1, nil, nil}
	n1 := Node{2, nil, nil}
	return n0.Insert(n1.Data)
}
