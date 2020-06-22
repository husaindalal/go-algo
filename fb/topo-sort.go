package fb

type TopoGraph struct {
	edges map[string][]string
}

func (t *TopoGraph) TopoSort() []string {

	// find indegree of all nodes
	indegree := map[string]int{}

	for from := range t.edges {
		indegree[from] = 0
	}

	for _, tos := range t.edges {
		for _, to := range tos {
			indegree[to]++
		}
	}

	// add 0 indegree nodes vertex queue
	queue := []string{}
	for vertex, degree := range indegree {
		if degree == 0 {
			queue = append(queue, vertex)
		}
	}

	result := []string{}
	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]

		result = append(result, vertex)

		neighbors := t.edges[vertex]
		for _, neighbor := range neighbors {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}

	}

	return nil
}
