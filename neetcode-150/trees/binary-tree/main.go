package main

import "fmt"

type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func (bs *BinaryTree) Insert(val int) {
	bs.InsertRec(bs.root, val)

}

func (bs *BinaryTree) InsertRec(node *TreeNode, val int) *TreeNode {
	if bs.root == nil {
		bs.root = &TreeNode{val: val}
		return bs.root
	}

	if node == nil {
		node = &TreeNode{val: val}
		return node
	}

	if node.val > val {
		node.left = bs.InsertRec(node.left, val)
	}
	if node.val < val {
		node.right = bs.InsertRec(node.right, val)
	}

	return node
}

func (bs *BinaryTree) Search(val int) bool {
	return bs.SearchRec(bs.root, val)
}

func (bs *BinaryTree) SearchRec(node *TreeNode, val int) bool {

	if node == nil {
		return false 
	}

	if node.val == val {
		return true
	}

	if val > node.val {
		return bs.SearchRec(node.right, val)
	} else if val < node.val {
		return bs.SearchRec(node.left, val)
	}

	return false
}

func (bs *BinaryTree) PreOrderTraversal() {
	bs.PreOrderTraversalH(bs.root)
}

func (bs *BinaryTree) PreOrderTraversalH(node *TreeNode) {

	if node == nil {
		return
	}

	root := node.val
	fmt.Printf("%v ->", root)

	bs.PreOrderTraversalH(node.left)
	bs.PreOrderTraversalH(node.right)

}

func (bs *BinaryTree) PostOrderTraversal() {
	bs.PostOrderTraversalH(bs.root)
}

func (bs *BinaryTree) PostOrderTraversalH(node *TreeNode) {

	if node == nil {
		return
	}

	bs.PostOrderTraversalH(node.left)
	bs.PostOrderTraversalH(node.right)
	root := node.val
	fmt.Printf("%v ->", root)

}

func (bs *BinaryTree) InOrderTraversal() {
	bs.InOrderTraversalH(bs.root)
}

func (bs *BinaryTree) InOrderTraversalH(node *TreeNode) {

	if node == nil {
		return
	}

	bs.InOrderTraversalH(node.left)
	root := node.val
	fmt.Printf("%v ->", root)

	bs.InOrderTraversalH(node.right)

}

func main() {

	bs := &BinaryTree{}
	bs.Insert(50)
	bs.Insert(5)
	bs.Insert(65)
	bs.Insert(89)
	bs.Insert(15)
	bs.Insert(12)

	fmt.Println("--Pre--")
	bs.PreOrderTraversal()
	fmt.Println()
	fmt.Println("--Post--")

	bs.PostOrderTraversal()
	fmt.Println()
	fmt.Println("--In--")

	bs.InOrderTraversal()

	// fmt.Printf("bs.Search(89): %v\n", bs.PreOrderTraversalH(89))
	fmt.Println()

	fmt.Printf("bs.Search(89): %v\n", bs.Search(8989))

}
