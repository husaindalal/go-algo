package graph

import "fmt"

type Vertex struct {
	Label string
	Value int
}

type Graph struct {
	//vertices []*Vertex
	Edges map[*Vertex][]*Vertex
}

func New() Graph {
	return Graph{
		Edges: map[*Vertex][]*Vertex{},
	}
}

func (g *Graph) AddEdge(v1, v2 *Vertex) {
	g.Edges[v1] = append(g.Edges[v1], v2)
	g.Edges[v2] = append(g.Edges[v2], v1)
}

func (g *Graph) Dfs(root *Vertex) {
	// stack.push(root)
	stack := []*Vertex{root}
	visited := map[*Vertex]bool{}
	for len(stack) != 0 {
		// stack.pop
		n := len(stack) - 1
		vertex := stack[n]
		stack = stack[:n]
		if !visited[vertex] {
			// mark visited
			visited[vertex] = true
			fmt.Printf("visiting node %v \n", vertex.Label)
			// add children to stack
			for _, neighbor := range g.Edges[vertex] {
				stack = append(stack, neighbor)
			}
		}
	}
}

func (g *Graph) Bfs(root *Vertex) {
	queue := []*Vertex{root}
	visited := map[*Vertex]bool{}
	for len(queue) != 0 {
		vertex := queue[0]
		queue = queue[1:]

		if !visited[vertex] {
			visited[vertex] = true
			fmt.Printf("visiting %v \n", vertex)
			for _, neighbor := range g.Edges[vertex] {
				if !visited[neighbor] {
					queue = append(queue, neighbor)
				}
			}
		}
	}
}

func (g *Graph) Sort() {
	// https://www.geeksforgeeks.org/given-sorted-dictionary-find-precedence-characters/
	//https://github.com/yourbasic/graph
}

func (g *Graph) Print() {

}

//
//func main() {
//	mygraph := graph.New()
//
//	va := &graph.Vertex{
//		Label: "A",
//		Value: 1,
//	}
//	vb := &graph.Vertex{
//		Label: "B",
//		Value: 1,
//	}
//	vc := &graph.Vertex{
//		Label: "C",
//		Value: 1,
//	}
//	mygraph.AddEdge(va, vc)
//	mygraph.AddEdge(va, vb)
//
//	mygraph.Dfs(va)
//	mygraph.Bfs(va)
//}
//
