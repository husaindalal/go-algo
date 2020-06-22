package hash

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindFrequentTreeSum(root *TreeNode) []int {
	hashmap := map[int]int{} // sum -> count of subtrees with sum

	postOrder(root, hashmap)

	fmt.Println(hashmap)

	//sort and return top frequent
	sortedSum := []int{}
	for _, v := range hashmap {
		sortedSum = append(sortedSum, v)
	}
	sort.SliceStable(sortedSum, func(i, j int) bool {
		return sortedSum[i] > sortedSum[j]
	})
	maxCount := sortedSum[0]
	//for h := range hashmap {
	//
	//}

	return []int{maxCount}
}

// returns sum
func postOrder(node *TreeNode, hashmap map[int]int) int {
	if node == nil {
		return 0
	}
	leftSum := postOrder(node.Left, hashmap)
	rightSum := postOrder(node.Right, hashmap)

	sum := node.Val + leftSum + rightSum
	hashmap[sum]++

	return sum
}

/*
	root := &hash.TreeNode{
		Val:   10,
		Left:  &hash.TreeNode{
			Val:   10,
			Left:  nil,
			Right: nil,
		},
		Right: &hash.TreeNode{
			Val:   10,
			Left:  &hash.TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
			Right: &hash.TreeNode{
				Val:   10,
				Left:  nil,
				Right: nil,
			},
		},
	}
	result := hash.FindFrequentTreeSum(root)

	fmt.Printf("result %v %v %v %v\n", result, "x","y","z")

*/
