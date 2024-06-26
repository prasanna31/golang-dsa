package binaryTree

import "fmt"

// Node represents a node in the tree
type Node struct {
	Value any
	Left  *Node
	Right *Node
}

// Tree represents a binary tree
type BinaryTree struct {
	Root *Node
}

// InitializeTree initializes a binary tree
func InitializeTree() *BinaryTree {
	return &BinaryTree{}
}

// Comparator defines a function that compares two elements and returns true if a <= b
type Comparator func(a, b any) bool

// intComparator is a comparator function for int type elements
func intComparator(a, b any) bool {
	return a.(int) < b.(int)
}

// Insert inserts a value into the binary tree
func (bt *BinaryTree) Insert(value any, cmp Comparator) {
	newNode := &Node{Value: value}

	if bt.Root == nil {
		bt.Root = newNode
		return
	}

	current := bt.Root
	for {
		if cmp(value, current.Value) {
			if current.Left == nil {
				current.Left = newNode
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				return
			}
			current = current.Right
		}
	}
}

func (bt *BinaryTree) Search(value any, cmp Comparator) *Node {
	current := bt.Root
	for current != nil {
		if cmp(value, current.Value) {
			current = current.Left
		} else if cmp(current.Value, value) {
			current = current.Right
		} else {
			return current
		}
	}
	return nil
}

// Delete deletes a value from the binary tree
func (bt *BinaryTree) Delete(value any, cmp Comparator) {
	bt.Root = deleteNode(bt.Root, value, cmp)
}

func deleteNode(root *Node, value any, cmp Comparator) *Node {
	if root == nil {
		return nil
	}

	if cmp(value, root.Value) {
		root.Left = deleteNode(root.Left, value, cmp)
	} else if cmp(root.Value, value) {
		root.Right = deleteNode(root.Right, value, cmp)
	} else {
		// Case 1: No child or only one child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		// Case 2: Two children
		// Find the minimum value in the right subtree
		min := findMin(root.Right)
		root.Value = min.Value

		// Delete the minimum node from the right subtree
		root.Right = deleteNode(root.Right, min.Value, cmp)
	}

	return root
}

func findMin(node *Node) *Node {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func preOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.Value)
	preOrderTraversal(node.Left)
	preOrderTraversal(node.Right)
}

func (bt *BinaryTree) PreOrder() {
	preOrderTraversal(bt.Root)
}

func inOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	inOrderTraversal(node.Left)
	fmt.Println(node.Value)
	inOrderTraversal(node.Right)
}

func (bt *BinaryTree) InOrder() {
	inOrderTraversal(bt.Root)
}

func postOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	postOrderTraversal(node.Left)
	postOrderTraversal(node.Right)
	fmt.Println(node.Value)
}

func (bt *BinaryTree) PostOrder() {
	postOrderTraversal(bt.Root)
}

func levelOrderTraversal(node *Node) {
	if node == nil {
		return
	}

	queue := []*Node{node}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		fmt.Println(current.Value)

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
}
