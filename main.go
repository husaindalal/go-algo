package main

import (
	"fmt"
	"go-algo/tree"
)

func main() {

	//search := &tree.TreeNode{
	//	Val:   12,
	//	Left:  nil,
	//	Right: nil,
	//}

	root := &tree.TreeNode{
		Val: 10,
		Left: &tree.TreeNode{
			Val: 5,
			Left: &tree.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &tree.TreeNode{
			Val: 15,
			Left: &tree.TreeNode{
				Val:   12,
				Left:  nil,
				Right: nil,
			},
			Right: &tree.TreeNode{
				Val:  17,
				Left: nil,
				Right: &tree.TreeNode{
					Val:   20,
					Left:  nil,
					Right: nil,
				},
			},
		},
	}

	//root := &tree.TreeNextNode{
	//	Val:   1,
	//	Left:  &tree.TreeNextNode{
	//		Val: 11,
	//		Left: &tree.TreeNextNode{
	//			Val:   111,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &tree.TreeNextNode{
	//			Val:   112,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//	Right: &tree.TreeNextNode{
	//		Val:   12,
	//		Left:  &tree.TreeNextNode{
	//			Val:   121,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//		Right: &tree.TreeNextNode{
	//			Val:   122,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//	},
	//}

	result := tree.InOrderSuccessor(root, root)
	fmt.Printf("result %v\n", result)
	tree.Print(root, 0, "O")
}
