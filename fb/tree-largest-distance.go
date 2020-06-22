package fb

// https://www.interviewbit.com/problems/largest-distance-between-nodes-of-a-tree/

type LDNode struct {
	Value int
	Child []*LDNode
}

func TreeLargestDistance(array []int) int {

	distMap := []int{}
	for i := 0; i < len(array); i++ {
		distMap = append(distMap, 1)
	}
	maxDist := 0
	for i := 0; i < len(array); i++ {
		//dist := distMap[i]
		j := array[i]
		for j != -1 {
			//if distMap[j] != 0 {
			//	dist = dist + distMap[j]
			//	break
			//}
			//dist++
			distMap[j]++
			distMap[i]++
			j = array[j]
		}
		if distMap[i] > maxDist {
			maxDist = distMap[i]
		}
		//distMap[i] = dist
	}

	return maxDist
}

func TreeLargestDistance2(root *LDNode) int {
	return treeDistHelper(root)
}

func treeDistHelper(node *LDNode) int {
	if node.Child == nil {
		return 1
	}
	// sum all children dist
	sum := 1
	for _, ch := range node.Child {
		sum = sum + treeDistHelper(ch)
	}
	return sum
}

/*
root := &fb.LDNode{
		Value: 0,
		Child: []*fb.LDNode{
			{
				Value: 1,
			},
			{
				Value: 2,
			},
			{
				Value: 3,
				Child: []*fb.LDNode{
					{
						Value: 4,
						Child: nil,
					},
				},
			},
		},
	}
*/
