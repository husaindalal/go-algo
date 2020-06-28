package tree

import "fmt"

func InorderTraversal(root *TreeNode) []int {
	result := []int{}

	//result = inorderRec(root, result)

	result = inorderIter(root)
	fmt.Println()
	result = inorderIter2(root)
	return result
}

func inorderRec(node *TreeNode, result []int) []int {
	if node == nil {
		return result
	}

	if node.Left != nil {
		result = inorderRec(node.Left, result)
	}
	result = append(result, node.Val)
	if node.Right != nil {
		result = inorderRec(node.Right, result)
	}

	return result
}

func inorderIter(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := []*TreeNode{}
	result := []int{}

	node := root
	for node != nil || len(stack) > 0 {
		fmt.Printf("first %v %v \n", node, stack)
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		fmt.Printf("second %v %v \n", node, stack)
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Printf("third %v %v \n", node, stack)
		result = append(result, node.Val)
		node = node.Right
		fmt.Printf("forth %v %v \n", node, stack)
	}
	return result
}

func inorderIter2(root *TreeNode) []int {
	res := []int{}
	stack := []*TreeNode{}
	cur := root

	for cur != nil || len(stack) != 0 {
		for cur != nil {
			fmt.Printf("first %v %v \n", cur, stack)

			stack = append(stack, cur)
			cur = cur.Left
		}

		fmt.Printf("second %v %v \n", cur, stack)
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("third %v %v \n", cur, stack)

		res = append(res, cur.Val)
		cur = cur.Right
		fmt.Printf("forth %v %v \n", cur, stack)
	}

	return res
}
