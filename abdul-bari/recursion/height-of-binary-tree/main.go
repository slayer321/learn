package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func CreateNode(val int) *TreeNode {
	return &TreeNode{
		Val:   val,
		Left:  nil,
		Right: nil,
	}
}

func (root *TreeNode) InsertLeft(val int) *TreeNode {
	root.Left = CreateNode(val)
	return root.Left
}

func (root *TreeNode) InsertRight(val int) *TreeNode {
	root.Right = CreateNode(val)
	return root.Right
}

// var count int = 1

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	// if root.Left == nil && root.Right == nil {
	// 	return 1
	// }

	//count++
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	gg := 1 + max(left, right)
	return gg
}

func main() {

	// root := CreateNode(3)
	// root.Left = CreateNode(9)
	// root.Right = CreateNode(20)
	// root.Right.Left = CreateNode(18)
	// root.Right.Right = CreateNode(7)

	// -----------------------------------------

	root := CreateNode(1)
	root.Right = CreateNode(2)

	// gg := root.InsertLeft(9)
	// fmt.Printf("gg: %v\n", gg)
	// root.InsertRight(20)
	// root.InsertLeft(18)
	// root.InsertRight(7)

	// fmt.Printf("root: %+v\n", root)

	fmt.Printf("count: %v\n", maxDepth(root))

}
