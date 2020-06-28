package tree

import "fmt"

type TreeNextNode struct {
	Val   int
	Left  *TreeNextNode
	Right *TreeNextNode
	Next  *TreeNextNode
}

func Connect(root *TreeNextNode) *TreeNextNode {
	queue := []*TreeNextNode{root}

	for len(queue) > 0 {
		// pop
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			node.Left.Next = node.Right
		}
		if node.Right != nil && node.Next != nil {
			node.Right.Next = node.Next.Left
		}

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return root
}

func PrintNextNode(node *TreeNextNode, space int, prefix string) {
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%v:%v->%v\n", prefix, node.Val, node.Next)

	if node.Left != nil {
		PrintNextNode(node.Left, space+2, "L")
	}
	if node.Right != nil {
		PrintNextNode(node.Right, space+2, "R")
	}
}
