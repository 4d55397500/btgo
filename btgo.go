// Package btgo contains typical operations
// over binary search trees
package btgo

import "fmt"

// A Node holds an int value and pointers
// to left and right child nodes.
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

// Check that the tree is a binary search tree
func (n *Node) IsBinarySearchTree() bool {

	condLeft := true
	condRight := true

	if n == nil {
		return true
	}
	if n.Left != nil {
		condLeft = n.Left.IsBinarySearchTree() && n.Data >= n.Left.Data
	}
	if n.Right != nil {
		condRight = n.Right.IsBinarySearchTree() && n.Data <= n.Right.Data
	}
	return condLeft && condRight
}

// Check for existence of the given Node
func (n *Node) Lookup(target int) bool {
	if n == nil {
		return false
	}
	if target == n.Data {
		return true
	} else if target < n.Data {
		return n.Left.Lookup(target)
	} else if target > n.Data {
		return n.Right.Lookup(target)
	}
	return false
}

// Insert a Node and return the new tree
func (n *Node) Insert(Data int) *Node {
	if n == nil {
		return &Node{Data, nil, nil}
	} else if Data < n.Data {
		n.Left = n.Left.Insert(Data)
	} else if Data > n.Data {
		n.Right = n.Right.Insert(Data)
	}
	return n
}

// Print the tree in order
func (n *Node) PrintTree() {
	if n != nil && n.Left != nil {
		n.Left.PrintTree()
	}
	if n != nil {
		fmt.Println(n.Data)
	}
	if n != nil && n.Right != nil {
		n.Right.PrintTree()
	}
}

// Return maximum depth of the tree
func (n *Node) MaxDepth() int {
	if n == nil {
		return 0
	}
	return maxInt(n.Left.MaxDepth()+1, n.Right.MaxDepth()+1)
}

func maxInt(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
