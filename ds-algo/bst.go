package main

import "fmt"

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func PreOrder(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.Value)
	PreOrder(node.Left)
	PreOrder(node.Right)
}

func InOrder(node *Node) {
	if node == nil {
		return
	}
	InOrder(node.Left)
	fmt.Println(node.Value)
	InOrder(node.Right)
}

func PostOrder(node *Node) {
	if node == nil {
		return
	}

	PostOrder(node.Left)
	PostOrder(node.Right)
	fmt.Println(node.Value)
}

func (node *Node) Insert(val int) {
	if node.Value < val {
		if node.Right == nil {
			node.Right = &Node{Value: val}
		} else {
			node.Right.Insert(val)
		}
	} else if node.Value > val {
		if node.Left == nil {
			node.Left = &Node{Value: val}
		} else {
			node.Left.Insert(val)
		}
	}
}

func (node *Node) Search(val int) bool {

	if node == nil {
		return false
	}

	if node.Value < val {
		return node.Right.Search(val)
	} else if node.Value > val {
		return node.Left.Search(val)
	}
	return true
}

func main() {
	fmt.Println("Binary Search Tree")
	n := &Node{Value: 100}
	n.Insert(50)
	n.Insert(120)
	n.Insert(40)
	n.Insert(70)
	n.Insert(200)
	n.Insert(105)

	fmt.Println("Root: ", n)

	fmt.Println("PreOrder")
	PreOrder(n)
	fmt.Println("InOrder")
	InOrder(n)
	fmt.Println("PostOrder")
	PostOrder(n)
	fmt.Println("Search function")
	fmt.Println(n.Search(120))
	fmt.Println(n.Search(45))

}
