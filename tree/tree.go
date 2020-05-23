package tree

import "container/list"

type Value int

type Node struct {
	Val Value
	Left *Node
	Right *Node
}

type Tree struct {
	Root *Node
}

func (t *Tree) LevelOrder() {
	list.
}

func (t *Tree) InOrder() {
	inOrderUtil(t.Root, )
}

func inOrderUtil(node *Node, f func(Value)) {
	if node.Left != nil {
		inOrderUtil(node.Left, f)
	}
	f(node.Val)
	if node.Right != nil {
		inOrderUtil(node.Right, f)
	}
}