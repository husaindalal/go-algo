package fb

import "fmt"

func SortItems(n int, m int, group []int, beforeItems [][]int) []int {

	// graph := buildGraph(group, beforeItems)

	// find indegree of each node
	// add indgree 0 nodes to queue
	// for queue != empty
	//.  pop latst and add to result
	//.  get all its before items and decrement its indegree by 1
	//.  if indegree 0 then add to queue

	groupGraph := buildGroupGraph(n, group, beforeItems)
	beforeGraph := buildBeforeGraph(n, beforeItems)

	fmt.Printf("groupGraph %v\n", groupGraph)
	fmt.Printf("beforeGraph %v\n", beforeGraph)

	return nil
}

type GroupGraph struct {
	Edges map[int][]int
}

func buildGroupGraph(n int, group []int, beforeItems [][]int) GroupGraph {
	graph := GroupGraph{
		Edges: map[int][]int{},
	}
	for i := 0; i < n; i++ {
		graph.Edges[group[i]] = append(graph.Edges[group[i]], beforeItems[i]...)

	}

	return graph
}

func buildBeforeGraph(n int, beforeItems [][]int) GroupGraph {
	graph := GroupGraph{
		Edges: map[int][]int{},
	}
	for i := 0; i < n; i++ {
		graph.Edges[i] = beforeItems[i]
		//for item := range beforeItems[i] {
		//    items = append(items, item)
		//}
	}

	return graph
}

/*

		fmt.Printf("sort items %v \n", fb.SortItems(8, 2, []int{-1,-1,1,0,0,1,0,-1}, [][]int{
			{},
			{6},
			{5},
			{6},
			{3,6},
			{},
			{},
			{},
	}))

*/
