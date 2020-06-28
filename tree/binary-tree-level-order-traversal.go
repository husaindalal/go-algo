package tree

func LevelOrder(root *TreeNode) [][]int {
	// BFS

	queue := []*TreeNode{root}
	result := [][]int{}

	for len(queue) > 0 {
		// pop
		cou := len(queue)

		res := []int{}
		for i := 0; i < cou; i++ {
			node := queue[0]
			queue = queue[1:]

			res = append(res, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, res)

	}

	return result
}
