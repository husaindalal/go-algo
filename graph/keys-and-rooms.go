package graph

/*
 start with room 0
 mark room0 visited
 for each key {
	mark the room visited.
	add the keys to the queue
 }
*/

func CanVisitAllRooms(rooms [][]int) bool {
	queue := []int{0}
	visited := make([]bool, len(rooms))

	for len(queue) > 0 {
		key := queue[0]
		queue = queue[1:]
		visited[key] = true

		// only add if visited == false
		for _, room := range rooms[key] {
			if !visited[room] {
				queue = append(queue, room)
			}
		}

	}

	allVisited := true
	for _, v := range visited {
		if v == false {
			allVisited = false
			break
		}
	}

	return allVisited
}

/*



	fmt.Printf("result %v\n", graph.CanVisitAllRooms([][]int{
		{1},
		{2},
		{3},
		{},
	}))

	fmt.Printf("result %v\n", graph.CanVisitAllRooms([][]int{
		{1, 3},
		{3, 0, 1},
		{2},
		{0},
	}))

*/
