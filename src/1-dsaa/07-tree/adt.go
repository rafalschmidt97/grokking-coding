package main

import (
	"fmt"
)

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

func NewNode(value int) *Node {
	return &Node{Data: value, Left: nil, Right: nil}
}

type BST struct {
	Root *Node
}

func findMin(node *Node) *Node {
	if node.Left == nil {
		return node
	}
	return findMin(node.Left)
}

func deleteNode(root *Node, value int) *Node {
	if root == nil {
		return root
	}

	if value < root.Data {
		root.Left = deleteNode(root.Left, value)
	} else if value > root.Data {
		root.Right = deleteNode(root.Right, value)
	} else {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}

		temp := findMin(root.Right)
		root.Data = temp.Data
		root.Right = deleteNode(root.Right, temp.Data)
	}
	return root
}

func (bst *BST) DeleteMethod(value int) {
	bst.Root = deleteNode(bst.Root, value)
}

func (bst *BST) Insert(value int) {
	newNode := NewNode(value)

	if bst.Root == nil {
		bst.Root = newNode
	} else {
		current := bst.Root
		var parent *Node

		for current != nil {
			parent = current
			if value < current.Data {
				current = current.Left
			} else {
				current = current.Right
			}
		}

		if value < parent.Data {
			parent.Left = newNode
		} else {
			parent.Right = newNode
		}
	}
}

func search(root *Node, value int) bool {
	current := root
	for current != nil {
		if value == current.Data {
			return true
		} else if value < current.Data {
			current = current.Left
		} else {
			current = current.Right
		}
	}
	return false
}

func (bst *BST) Search(value int) bool {
	return search(bst.Root, value)
}

func inOrderHelper(node *Node) {
	if node != nil {
		inOrderHelper(node.Left)
		fmt.Print(node.Data, " ")
		inOrderHelper(node.Right)
	}
}

func (bst *BST) InOrder() {
	inOrderHelper(bst.Root)
	fmt.Println()
}

func preOrderHelper(node *Node) {
	if node != nil {
		fmt.Print(node.Data, " ")
		preOrderHelper(node.Left)
		preOrderHelper(node.Right)
	}
}

func (bst *BST) PreOrder() {
	preOrderHelper(bst.Root)
	fmt.Println()
}

func postOrderHelper(node *Node) {
	if node != nil {
		postOrderHelper(node.Left)
		postOrderHelper(node.Right)
		fmt.Print(node.Data, " ")
	}
}

func (bst *BST) PostOrder() {
	postOrderHelper(bst.Root)
	fmt.Println()
}

type Solution struct{}

func main() {
	obj := &BST{}
	obj.Insert(5)
	obj.Insert(3)
	obj.Insert(7)
	obj.Insert(2)
	obj.Insert(4)
	obj.Insert(6)
	obj.Insert(8)

	obj.InOrder()
	obj.PreOrder()
	obj.PostOrder()

	obj.DeleteMethod(2)

	obj.InOrder()
}
