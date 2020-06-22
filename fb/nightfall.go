package fb

import "fmt"

type StretchNode struct {
	Value int
	Left  *StretchNode
	Right *StretchNode
}

func Stretch(root *StretchNode, amount int) {
	// validation TODO
	stretchHelper(root, amount, true)
}

func stretchHelper(node *StretchNode, amount int, isLeft bool) {
	leftChild := node.Left
	rightChild := node.Right
	node.Value = node.Value / amount
	node.Left = nil
	node.Right = nil
	for i := 0; i < amount-1; i++ {
		fmt.Printf("i %v, node %v\n", i, node)
		child := &StretchNode{
			Value: node.Value, // TODO remove i
			Left:  nil,
			Right: nil,
		}
		if isLeft {
			node.Left = child
		} else {
			node.Right = child
		}
		node = child
	}
	fmt.Printf("node %v, leftChild %v, rightChild %v\n", node, leftChild, rightChild)

	if leftChild != nil {
		node.Left = leftChild
		stretchHelper(leftChild, amount, true)
	}
	if rightChild != nil {
		node.Right = rightChild
		stretchHelper(rightChild, amount, false)
	}

}

//func printStretch(node *StretchNode) {
//
//	fmt.Printf(", %v ", node.Value)
//	if node.Left == nil && node.Right == nil {
//		return
//	}
//	if node.Left != nil {
//		print(node.Left)
//	}
//	if node.Right != nil {
//		print(node.Right)
//	}
//
//}

/*

	root := &Node{
		Value: 12,
		Left:  &Node{
			Value: 81,
			Left:  nil,
			Right: &Node{
				Value: 56,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &Node{
			Value: 34,
			Left:  &Node{
				Value: 19,
				Left:  nil,
				Right: nil,
			},
			Right: &Node{
				Value: 6,
				Left:  nil,
				Right: nil,
			},
		},
	}
	//
	//root := &Node{
	//	Value: 12,
	//	Left: nil,
	//	Right: &Node{
	//			Value: 56,
	//			Left:  nil,
	//			Right: nil,
	//		},
	//}
	//print(root)
	Stretch(root, 2)
	fmt.Printf("factor 2 %v\n", root)
	print(root)

*/
