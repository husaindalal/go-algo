package tree

import (
	"fmt"
)

type Value string

type Node struct {
	Val   Value
	Left  *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) LevelOrder() {
	if t == nil {
		fmt.Println("nil")
		return
	}
	queue := []*Node{t.Root}
	results := [][]Value{}

	//levelOrderUtil(queue, 0, results)
	fmt.Println(results)

	levelOrderLoop(queue)
}

func levelOrderUtil(queue []*Node, level int, results [][]Value) {
	if len(queue) == 0 {
		return
	}

	node := queue[0] // front
	fmt.Println(node)

	if node.Left != nil {
		queue = append(queue, node.Left)
	}
	if node.Right != nil {
		queue = append(queue, node.Right)
	}
	queue = queue[1:] //remove
	levelOrderUtil(queue, level+1, results)

}

func levelOrderLoop(queue []*Node) {
	for len(queue) != 0 {
		node := queue[0] // front
		fmt.Println(node)

		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		queue = queue[1:] //remove
	}
}

func (t *Tree) InOrder() {
	stack := []*Node{t.Root}
	fmt.Println(inOrderUtil(stack))
}

// DFS
func inOrderUtil(stack []*Node) []Value {
	result := []Value{}
	if len(stack) == 0 {
		return result
	}
	// pop
	for len(stack) > 0 {
		n := len(stack) - 1
		node := stack[n]
		stack = stack[:n]
		if node.Right != nil {
			stack = append(stack, node.Right)
			continue
		}
		fmt.Println(node)
		result = append(result, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
			continue
		}

	}
	return result
}

func (t *Tree) IsSymmetric() bool {
	// validate
	if t == nil {
		return true
	}
	queueL := []*Node{t.Root}
	queueR := []*Node{t.Root}
	isSymmetric := true
	for len(queueL) != 0 && len(queueR) != 0 {
		nodeL := queueL[0]
		nodeR := queueR[0]
		queueL = queueL[1:]
		queueR = queueR[1:]

		if nodeL == nil && nodeR == nil {
			continue
		}
		if (nodeL == nil && nodeR != nil) ||
			(nodeL != nil && nodeR == nil) ||
			(nodeL.Val != nodeR.Val) {
			isSymmetric = false
			break
		}
		queueL = append(queueL, nodeL.Left)
		queueR = append(queueR, nodeR.Right)
		queueL = append(queueL, nodeL.Right)
		queueR = append(queueR, nodeR.Left)
	}

	return isSymmetric
}

func (t *Tree) LeftVisibleNodes() []*Node {
	// if t == nil ...
	// bfs with level
	level := 0
	levelCount := 1
	queue := []*Node{t.Root}
	results := []*Node{t.Root}

	for len(queue) != 0 {
		nextLevelCount := 0
		level++

		for levelCount > 0 {
			levelCount--

			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
				if nextLevelCount == 0 {
					results = append(results, node.Left)
				}
				nextLevelCount = nextLevelCount + 1
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				if nextLevelCount == 0 {
					results = append(results, node.Right)
				}
				nextLevelCount = nextLevelCount + 1
			}

		}
		levelCount = nextLevelCount
		fmt.Printf("level %v, levelCount %v \n", level, levelCount)
	}

	return results
}

//func main() {
//	mytree := tree.Tree{
//		Root: &tree.Node{
//			Val:   10,
//			Left:  &tree.Node{
//				Val: 20,
//				Left: &tree.Node{
//					Val: 40,
//				},
//			},
//			Right: &tree.Node{
//				Val: 30,
//				Left: &tree.Node{
//					Val:   50,
//				},
//			},
//		},
//	}
//
//	mytree.LevelOrder()
//	mytree.InOrder()
//}

//
//func main() {
//	mytree := tree.Tree{
//		Root: &tree.Node{
//			Val: "A",
//			Left: &tree.Node{
//				Val: "B",
//				Left: &tree.Node{
//					Val: "D",
//				},
//				Right: &tree.Node{
//					Val: "E",
//				},
//			},
//			Right: &tree.Node{
//				Val: "C",
//				Right: &tree.Node{
//					Val: "G",
//					Left: &tree.Node{
//						Val: "H",
//						Right: &tree.Node{
//							Val: "I",
//						},
//					},
//				},
//			},
//		},
//	}
//
//	//mytree.LevelOrder()
//	//mytree.InOrder()
//
//	//fmt.Println(mytree.IsSymmetric())
//	results := mytree.LeftVisibleNodes()
//	for _, res := range results {
//		fmt.Println(res.Val)
//	}
//}
