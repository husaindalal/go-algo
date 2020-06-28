package graph

/*
equations = [ ["a", "b"], ["b", "c"] ],
values = [2.0, 3.0],
queries = [ ["a", "c"], ["b", "a"], ["a", "e"], ["a", "a"], ["x", "x"] ].
*/

type EqVertex struct {
	Node  string
	Value float64
}

type EqGraph struct {
	Edges map[string][]*EqVertex
}

func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	// validate

	graph := EqGraph{
		Edges: map[string][]*EqVertex{},
	}

	for i := 0; i < len(equations); i++ {
		eq := equations[i]
		graph.Edges[eq[0]] = append(graph.Edges[eq[0]], &EqVertex{
			Node:  eq[1],
			Value: values[i],
		})
		graph.Edges[eq[1]] = append(graph.Edges[eq[1]], &EqVertex{
			Node:  eq[0],
			Value: 1 / values[i],
		})
	}

	// BFS
	result := []float64{}
	for _, q := range queries {
		//queue := []string{}
		result = append(result, bfs(graph, q))
	}

	return result
}

func bfs(graph EqGraph, query []string) float64 {
	//validate

	result := -1.0

	if graph.Edges[query[0]] == nil || graph.Edges[query[1]] == nil {
		return -1.0
	}

	queue := []EqVertex{{
		Node:  query[0],
		Value: 1.0,
	}}
	visited := map[string]bool{query[0]: true}

	for len(queue) > 0 {
		node := queue[0].Node
		multiplier := queue[0].Value
		queue = queue[1:]

		for _, edge := range graph.Edges[node] {
			if !visited[edge.Node] {
				if edge.Node == query[1] {
					result = multiplier * edge.Value
					break
				}

				visited[edge.Node] = true
				queue = append(queue, EqVertex{
					Node:  edge.Node,
					Value: multiplier * edge.Value,
				})
			}
		}
	}

	return result
}

/*

	equations := [][]string{
		{"a", "b"},
		{"b", "c"},
	}
	values := []float64{2.0, 3.0}
	queries := [][]string {
		{"a", "c"},
		{"b", "a"},
		{"a", "e"},
		{"a", "a"},
		{"x", "x"},
	}


	fmt.Printf("result %v\n", graph.CalcEquation(equations, values, queries))


*/
