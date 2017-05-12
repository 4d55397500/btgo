package btgo;


import "fmt"


type node struct {
	data int;
	left *node;
	right *node;
}

// Check that the tree is a binary search tree
func (n *node) IsBinarySearchTree() bool {

	condLeft := true
	condRight := true

	if n == nil {
		return true
	}
	if n.left != nil {
		condLeft = n.left.IsBinarySearchTree() && n.data >= n.left.data;
	}
	if n.right != nil {
		condRight = n.right.IsBinarySearchTree() && n.data <= n.right.data;
	}
	return condLeft && condRight
}

// Check for existence of the given node
func (n *node) Lookup(target int) bool {
	if (n == nil) {
		return false
	}
	if (target == n.data) {
		return true
	} else if target < n.data {
		return n.left.Lookup(target)
	} else if target > n.data {
		return n.right.Lookup(target)
	}
	return false
}


// Insert a node and return the new tree
func (n *node) Insert(data int) *node {
	if n == nil {
		return &node{data, nil, nil}
	} else if data < n.data {
		n.left = n.left.Insert(data)
	} else if data > n.data {
		n.right = n.right.Insert(data)
	}
	return n
}


// Print the tree in order
func (n *node) PrintTree() {
	if n != nil && n.left != nil {
		n.left.PrintTree()
	}
	if n != nil {
		fmt.Println(n.data)
	}
	if n != nil && n.right != nil {
		n.right.PrintTree()
	}
}


// Return maximum depth of the tree
func (n *node) MaxDepth() int {
	if n == nil {
		return 0
	}
	return maxInt(n.left.MaxDepth() + 1, n.right.MaxDepth() + 1)
}



func maxInt(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
