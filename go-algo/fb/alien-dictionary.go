package fb

import (
	"fmt"
	"math"
)

/*
  eterq
  errsd
  yrtwe
  yrzdvs
  fhhs
  fff
  tjsg
  tjadsd
*/

type Vertex struct {
	Label string
}

type Graph struct {
	Edges map[Vertex][]Vertex
	//Vertices map[Vertex]bool
}

func ArrangeLetters(dictionary []string) []Vertex {
	// for each word and next word, navigate to the first different letter
	// add two letters and relationship in graph
	// topological sort the graph by first letter

	graph := Graph{
		Edges: map[Vertex][]Vertex{},
		//Vertices: map[Vertex]bool{},
	}

	for i := 0; i < len(dictionary)-1; i++ {
		smallLen := int(math.Min(float64(len(dictionary[i])), float64(len(dictionary[i+1]))))
		j := 0
		for j < smallLen {
			if dictionary[i][j] == dictionary[i+1][j] {
				j++
			} else {
				break
			}
		}
		if j < smallLen {
			fmt.Printf("i %v j %v smallLen %v \n", i, j, smallLen)
			vertexBefore := Vertex{
				Label: string(dictionary[i][j]),
			}
			vertexAfter := Vertex{
				Label: string(dictionary[i+1][j]),
			}
			fmt.Printf("first %c second %c \n", vertexBefore, vertexAfter)
			graph.Edges[vertexBefore] = append(graph.Edges[vertexBefore], vertexAfter)
		}

	}

	return sortGraph(graph)

}

// Topological sort
//https://www.geeksforgeeks.org/topological-sorting/

func sortGraph(graph Graph) []Vertex {

	// find the indegree of each vertex
	// collect vertices with indegree 0 in queue
	// for each node in queue, append to result, reduce indegree of adjacent nodes. If indegree is 0 add to queue
	indegree := map[Vertex]int{}

	for vertex, val := range graph.Edges {
		fmt.Printf("graph %c %c\n", vertex, val)
		indegree[vertex] = 0
	}

	for _, neighbors := range graph.Edges {
		for _, v := range neighbors {
			indegree[v]++

		}
	}

	queue := []Vertex{}
	for v, degree := range indegree {
		fmt.Printf("Indegree %v %d \n", v, degree)
		if degree == 0 {
			queue = append(queue, v)
		}
	}

	result := []Vertex{}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]

		result = append(result, v)

		neighbors := graph.Edges[v]
		for _, n := range neighbors {
			indegree[n]--
			if indegree[n] == 0 {
				queue = append(queue, n)
			}
		}

		fmt.Printf("queue node %c \n", v)

	}

	return result
}

/*
	result := fb.ArrangeLetters([]string{
		"eterq",
		"errsd",
		"yrtwe",
		"yrzdvs",
		"fhhs",
		"fff",
		"tjsg",
		"tjadsd",
		"tjadsdrr",
	})

	for _, r := range result {
		fmt.Printf("arranged %c\n", r)

	}
*/
