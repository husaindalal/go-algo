package tree

func InOrderSuccessor(root *TreeNode, search *TreeNode) *TreeNode {
	result := &TreeNode{}
	stack := []*TreeNode{}
	searchFound := false

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
		if searchFound == true {
			result = cur
			break
		}
		if searchFound == false && cur == search {
			searchFound = true
		}

		cur = cur.Right // important

	}

	return result
}
