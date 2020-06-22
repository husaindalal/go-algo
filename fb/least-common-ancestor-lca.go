package fb

import (
	"fmt"
)

type LCANode struct {
	Value int
	Left  *LCANode
	Right *LCANode
}

func LCA(node *LCANode, x, y int) *LCANode {
	if node.Value == x || node.Value == y {
		return node
	}
	if node.Left == nil && node.Right == nil {
		return nil
	}
	leftResult := &LCANode{}
	if node.Left != nil {
		leftResult = LCA(node.Left, x, y)
	} else {
		leftResult = nil
	}

	rightResult := &LCANode{}
	if node.Right != nil {
		rightResult = LCA(node.Right, x, y)
	} else {
		rightResult = nil
	}
	if leftResult != nil && rightResult != nil {
		return node
	} else if leftResult != nil {
		return leftResult
	} else {
		return rightResult
	}

}

func PrintLCA(node *LCANode, ns int, ch byte) {
	if node == nil {
		return
	}

	for i := 0; i < ns; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%c:%v\n", ch, node.Value)
	PrintLCA(node.Left, ns+2, 'L')
	PrintLCA(node.Right, ns+2, 'R')
}

/*
	root := &fb.LCANode{
		Value: 12,
		Left:  &fb.LCANode{
			Value: 81,
			Left:  nil,
			Right: &fb.LCANode{
				Value: 56,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &fb.LCANode{
			Value: 34,
			Left:  &fb.LCANode{
				Value: 19,
				Left:  nil,
				Right: nil,
			},
			Right: &fb.LCANode{
				Value: 6,
				Left:  nil,
				Right: nil,
			},
		},
	}
	//
	//root := &fb.LCANode{
	//	Value: 12,
	//	Left: nil,
	//	Right: &fb.LCANode{
	//			Value: 56,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//}
	//print(root)
	result := fb.LCA(root, 19, 6)
	fmt.Printf("factor %v\n", result)
	fb.PrintLCA(root, 2, 'O')
*/
