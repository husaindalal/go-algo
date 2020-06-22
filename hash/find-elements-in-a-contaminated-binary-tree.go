package hash

import "fmt"

type BinaryTreeNode struct {
	Value int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type FindElements struct {
	Root *BinaryTreeNode
}

func Constructor(root *BinaryTreeNode) FindElements {
	// level order
	//elements := FindElements{
	//	Root: &BinaryTreeNode{
	//		Value: 0,
	//		Left:  nil,
	//		Right: nil,
	//	},
	//}
	root.Value = 0

	queue := []*BinaryTreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			node.Left.Value = 2*node.Value + 1
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			node.Right.Value = 2*node.Value + 2
			queue = append(queue, node.Right)
		}

	}
	return FindElements{
		Root: root,
	}
}

func (f *FindElements) Find(target int) bool {
	queue := []*BinaryTreeNode{f.Root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Value == target {
			return true
		}
		if node.Value > target {
			continue
		}

		if node.Left != nil {
			//node.Left.Value = 2 * node.Value + 1
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			//node.Right.Value = 2 * node.Value + 2
			queue = append(queue, node.Right)
		}

	}
	return false //TODO
}

func (f *FindElements) Print() {
	print(f.Root, 0, "O")
}

func print(node *BinaryTreeNode, spaces int, char string) {
	for i := 0; i < spaces; i++ {
		fmt.Print(" ")
	}
	//queue := []*BinaryTreeNode{}
	fmt.Printf("%v:%v\n", char, node.Value)
	if node.Left != nil {
		print(node.Left, spaces+2, "L")
	}

	if node.Right != nil {
		print(node.Right, spaces+2, "R")
	}
}

/*

	smal := hash.Constructor(&hash.BinaryTreeNode{
		Value: -1,
		Left:  nil,
		Right: &hash.BinaryTreeNode{
			Value: -1,
			Left:  &hash.BinaryTreeNode{
				Value: -1,
				Left:  &hash.BinaryTreeNode{
					Value: -1,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			Right: nil,
		},
	})

	smal.Print()

	result := smal.Find(3)
	fmt.Printf("result %v %v %v %v\n", result, "x","y","z")

*/
