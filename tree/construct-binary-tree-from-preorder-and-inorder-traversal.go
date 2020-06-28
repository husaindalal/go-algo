package tree

import "fmt"

// NOTE prestart+1+i-instart
func BuildTree(preorder []int, inorder []int) *TreeNode {
	root := &TreeNode{
		Val: preorder[0],
	}

	root = helper(preorder, inorder, 0, 0, len(inorder)-1)

	return root
}

func helper(preorder []int, inorder []int, prestart int, instart int, inend int) *TreeNode {
	fmt.Printf("first %v %v %v \n", prestart, instart, inend)
	if prestart > len(preorder)-1 || instart > inend {
		return nil
	}

	node := &TreeNode{
		Val: preorder[prestart],
	}
	i := instart
	for i <= inend && inorder[i] != node.Val {
		i++
	}

	fmt.Printf("third %v %v %v %v \n", prestart, instart, inend, i)
	node.Left = helper(preorder, inorder, prestart+1, instart, i-1)
	fmt.Printf("node.left %v\n", node.Left)
	node.Right = helper(preorder, inorder, prestart+1+i-instart, i+1, inend)
	fmt.Printf("node.Right %v\n", node.Right)

	return node
}

func Print(node *TreeNode, space int, nodePos string) {
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}

	fmt.Printf("%v:%v\n", nodePos, node.Val)
	if node.Left != nil {
		Print(node.Left, space+2, "L")
	}
	if node.Right != nil {
		Print(node.Right, space+2, "R")
	}
}
