package fb

import (
	"fmt"
)

type BiNode struct {
	Value string
	Left  *BiNode
	Right *BiNode
}

type BSTIterator struct {
	Root *BiNode
	//position *BiNode
	Stack []*BiNode
}

func (b *BSTIterator) Next() string {
	// validate
	if !b.HasNext() {
		return ""
	}

	n := len(b.Stack) - 1
	node := b.Stack[n]
	b.Stack = b.Stack[:n]

	b.AddAll(node.Right)
	//if node.Right != nil {
	//	b.Stack = append(b.Stack, node.Right)
	//	left := node.Right.Left
	//	for left != nil {
	//		b.Stack = append(b.Stack, left)
	//		left = left.Left
	//	}
	//}

	return node.Value
}

func (b *BSTIterator) AddAll(node *BiNode) {
	if node != nil {
		b.Stack = append(b.Stack, node)
		for node.Left != nil {
			b.Stack = append(b.Stack, node.Left)
			node = node.Left
		}
	}
}

func (b *BSTIterator) HasNext() bool {
	return len(b.Stack) != 0
}

func (b *BSTIterator) Print() {
	print(b.Root, 0, "O")
	for _, s := range b.Stack {
		fmt.Printf("stack %v\n", s)
	}
	fmt.Println()
}

func print(node *BiNode, space int, nodePos string) {
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}

	fmt.Printf("%v:%v\n", nodePos, node.Value)
	if node.Left != nil {
		print(node.Left, space+2, "L")
	}
	if node.Right != nil {
		print(node.Right, space+2, "R")
	}
}

/*
	root := &fb.BiNode{
		Value: "12",
		Left:  &fb.BiNode{
			Value: "81",
			Left:  nil,
			Right: &fb.BiNode{
				Value: "56",
				Left:  nil,
				Right: nil,
			},
		},
		Right: &fb.BiNode{
			Value: "34",
			Left:  &fb.BiNode{
				Value: "19",
				Left:  nil,
				Right: nil,
			},
			Right: &fb.BiNode{
				Value: "6",
				Left:  nil,
				Right: nil,
			},
		},
	}

	iterator := fb.BSTIterator{
		GetRoot:  root,
		Stack: []*fb.BiNode{},
	}

	//
	//result := fb.LCA(root, 19, 6)
	//fmt.Printf("factor %v\n", result)
	//fb.PrintLCA(root, 2, 'O')

	// []int{-5,-1,-1, 0, 3,5,9}

	iterator.AddAll(root)

	iterator.Print()
	result := iterator.Next()
	fmt.Printf("factor %v %v %v %v\n", result, "x","y","z")
	iterator.Print()

	result = iterator.Next()
	fmt.Printf("factor %v %v %v %v\n", result, "x","y","z")
	iterator.Print()

	result = iterator.Next()
	fmt.Printf("factor %v %v %v %v\n", result, "x","y","z")
	iterator.Print()

	result = iterator.Next()
	fmt.Printf("factor %v %v %v %v\n", result, "x","y","z")
	iterator.Print()

*/
