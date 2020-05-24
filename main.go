package main

import (
	"go-algo/tree"
)

func main() {
	mytree := tree.Tree{
		Root: &tree.Node{
			Val: 'A',
			Left: &tree.Node{
				Val: 'B',
				Left: &tree.Node{
					Val: 'D',
				},
				Right: &tree.Node{
					Val: 'E',
				},
			},
			Right: &tree.Node{
				Val: 'C',
				Right: &tree.Node{
					Val: 'G',
				},
			},
		},
	}

	//mytree.LevelOrder()
	//mytree.InOrder()

	//fmt.Println(mytree.IsSymmetric())
	mytree.LeftVisibleNodes()
}
