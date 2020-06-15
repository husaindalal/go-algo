package graph

type Matrix struct {
	Edges [][]bool
}

func NewMatrix() Matrix {
	return Matrix{
		Edges: [][]bool{},
	}
}

type Coordinate struct {
	X int
	Y int
}

var directions = []Coordinate{
	{
		X: 0,
		Y: 1,
	},
	{
		X: 1,
		Y: 0,
	},
	{
		X: -1,
		Y: 0,
	},
	{
		X: 0,
		Y: -1,
	},
}

func (m *Matrix) MazeDistance(start, end Coordinate) (bool, error) {
	// validate: start and end within the matrix
	if len(m.Edges) == 0 {
		return false, nil
	}
	visited := map[Coordinate]bool{} // [][]bool{}
	maxRows := len(m.Edges)
	maxCols := len(m.Edges[0])
	// distance := 0

	queue := []Coordinate{start}

	for len(queue) != 0 {
		vertex := queue[0]
		queue = queue[1:]

		for _, dir := range directions {
			nextCoordinate := Coordinate{
				X: vertex.X + dir.X,
				Y: vertex.Y + dir.Y,
			}
			//if nextCoordinate withing range and not visited
			if nextCoordinate.X >= 0 && nextCoordinate.X < maxRows &&
				nextCoordinate.Y >= 0 && nextCoordinate.Y < maxCols &&
				m.Edges[nextCoordinate.X][nextCoordinate.Y] &&
				!visited[nextCoordinate] {

				if nextCoordinate == end {
					return true, nil
				}
				visited[nextCoordinate] = true
				queue = append(queue, nextCoordinate)
			}
		}
	}

	return false, nil
}

//
//func main() {
//	mygraph := graph.NewMatrix()
//
//	mygraph.Edges = [][]bool{
//		{true, false, false, true},
//		{true, true, false, true},
//		{false, true, false, true},
//	}
//
//	result, _ := mygraph.MazeDistance(graph.Coordinate{X: 0, Y:0}, graph.Coordinate{X:2, Y:1})
//	fmt.Println(result)
//}
//
