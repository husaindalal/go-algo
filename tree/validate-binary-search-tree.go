package tree

import "math"

func IsValidBST(root *TreeNode) bool {
	result := true
	stack := []*TreeNode{}
	//result := helperValidBST(root, math.MinInt32)
	lastVal := math.MinInt32

	cur := root
	for cur != nil || len(stack) > 0 {
		// add all left
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		// pop
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if cur.Val < lastVal {
			result = false
			break
		}
		lastVal = cur.Val

		cur = cur.Right // important

	}

	return result
}

//func helperValidBST(node *TreeNode, prevVal int, result bool) bool {
//	if result == false {
//		return result
//	}
//	if node == nil {
//		return result
//	}
//	if node.Left != nil {
//		result = helperValidBST(node.Left, prevVal, result)
//	}
//	if prevVal > node.Val {
//		return false
//	}
//	if result && node.Right != nil {
//		result = helperValidBST(node.Right, prevVal, result)
//	}
//	return result
//}
