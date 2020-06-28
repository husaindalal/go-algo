package tree

func ZigzagLevelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	zig := true
	result := [][]int{}
	for len(queue) > 0 {
		cou := len(queue)
		res := []int{}
		for i := 0; i < cou; i++ {
			// pop
			node := queue[0]
			queue = queue[1:]
			res = append(res, node.Val)
			// if zig, left then right
			if zig {
				if node.Left != nil {
					//
				}
				if node.Right != nil {
					//
				}
			} else {
				if node.Right != nil {
					//
				}
				if node.Left != nil {
					//
				}
			}

			// if !zig right then left
		}
		result = append(result, res)
	}

	return result
}
