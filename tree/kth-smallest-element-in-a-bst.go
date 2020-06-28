package tree

import (
	"fmt"
	"math"
)

func KthSmallest(root *TreeNode, k int) int {
	node := root
	stack := []*TreeNode{}
	result := math.MinInt32
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		// pop
		node = stack[len(stack)-1]
		stack = stack[0 : len(stack)-1]

		fmt.Printf("node val %v\n", node.Val)

		if k == 1 {
			result = node.Val
			break
		}
		k--
		node = node.Right

	}

	return result
}
