package graph

//https://leetcode.com/problems/validate-binary-tree-nodes/

/*
Input: n = 6, leftChild = [1,-1,-1,4,-1,-1], rightChild = [2,-1,-1,5,-1,-1]
Output: false

Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,3,-1,-1]
Output: false

Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]
Output: true
*/

// Assuming 0 as the root node
func ValidateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	indegree := make([]int, n)

	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			indegree[leftChild[i]]++
		}
		if rightChild[i] != -1 {
			indegree[rightChild[i]]++
		}
	}

	numOfRoots := 0
	for i := 0; i < n; i++ {
		if indegree[i] > 1 {
			return false
		}
		if indegree[i] == 0 {
			numOfRoots++
		}
	}
	if numOfRoots != 1 {
		return false
	}

	return true
}

/*
	Input: n = 6, leftChild = [1,-1,-1,4,-1,-1], rightChild = [2,-1,-1,5,-1,-1]
	Output: false

	Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,3,-1,-1]
	Output: false

	Input: n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]
	Output: true

fmt.Printf("result %v\n", graph.ValidateBinaryTreeNodes(6, []int{1,-1,-1,4,-1,-1}, []int{2,-1,-1,5,-1,-1}))

fmt.Printf("result %v\n", graph.ValidateBinaryTreeNodes(4,[]int{1,-1,3,-1}, []int{2,3,-1,-1} ))
fmt.Printf("result %v\n", graph.ValidateBinaryTreeNodes(4,[]int{1,-1,3,-1}, []int{2,-1,-1,-1} ))
fmt.Printf("result %v\n", graph.ValidateBinaryTreeNodes(1,[]int{-1}, []int{-1} ))

*/
