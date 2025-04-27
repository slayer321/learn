package main

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Data  int
}

type BinaryTree struct {
	root *Node
}

func NewNode(data int) *Node {
	return &Node{
		Data: data,
	}
}

func (bt *BinaryTree) Insert(val int) {
	bt.InsertRec(bt.root, val)
}

func (bt *BinaryTree) InsertRec(node *Node, val int) *Node {
	if bt.root == nil {
		bt.root = NewNode(val)
		return bt.root
	}

	if node == nil {
		return NewNode(val)
	}
	if val > bt.root.Data {
		node.Right = bt.InsertRec(node.Right, val)
	}
	if val <= bt.root.Data {
		node.Left = bt.InsertRec(node.Left, val)
	}
	return node
}

func (bt *BinaryTree) Search(val int) bool {
	return bt.SearchRec(bt.root, val)
}

func (bt *BinaryTree) SearchRec(node *Node, val int) bool {

	if node == nil {
		return false
	}
	if node.Data == val {
		return true
	}

	if val < node.Data {
		return bt.SearchRec(node.Left, val)
	}
	if val > node.Data {
		return bt.SearchRec(node.Right, val)
	}

	return false

}

func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%v\n", root.Data)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

func Inorder(root *Node) {
	if root == nil {
		return
	}
	Inorder(root.Left)
	fmt.Printf("%v\n", root.Data)
	Inorder(root.Right)
}

func PostOrder(root *Node) {
	if root == nil {
		return
	}
	PostOrder(root.Left)
	PostOrder(root.Right)
	fmt.Printf("%v\n", root.Data)

}

func main() {

	rootNode := BinaryTree{}
	rootNode.Insert(50)
	rootNode.Insert(5)
	rootNode.Insert(65)
	rootNode.Insert(89)
	rootNode.Insert(15)
	rootNode.Insert(12)

	PreOrderTraversal(rootNode.root)
	fmt.Println("-----------------")
	Inorder(rootNode.root)
	fmt.Println("-----------------")
	PostOrder(rootNode.root)

	fmt.Println(rootNode.Search(112))

}
