package tree

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil { //|| (root.Left == nil && root.Right == nil)
		return nil
	}

	if root == p {
		return root
	}
	if root == q {
		return root
	}
	if root.Left == nil && root.Right == nil {
		return nil
	}

	leftLCA := &TreeNode{}
	if root.Left != nil {
		leftLCA = LowestCommonAncestor(root.Left, p, q)
	} else {
		leftLCA = nil
	}

	rightLCA := &TreeNode{}
	if root.Right != nil {
		rightLCA = LowestCommonAncestor(root.Right, p, q)
	} else {
		rightLCA = nil
	}

	if leftLCA != nil && rightLCA != nil {
		return root
	}

	if leftLCA == nil && rightLCA != nil {
		return rightLCA
	}
	if rightLCA == nil && leftLCA != nil {
		return leftLCA
	}
	return nil

}
