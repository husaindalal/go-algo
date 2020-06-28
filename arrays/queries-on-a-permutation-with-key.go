package arrays

import "container/list"

func ProcessQueries(queries []int, m int) []int {
	permutation := list.New()
	for i := 1; i <= m; i++ {
		permutation.PushBack(i)
	}
	result := []int{}

	for _, q := range queries {
		res := move(permutation, q)
		result = append(result, res)
	}

	return result
}

func move(permutation *list.List, pos int) int {
	// find the index of pos
	i := 0
	node := &list.Element{}
	result := i
	for e := permutation.Front(); e != nil; e = e.Next() {
		// do something with e.Value

		if e.Value == pos {
			result = i
			node = e
		}
		i++
	}

	// result is the index

	// remove the node in index and move to front
	permutation.MoveToFront(node)

	return result
}

/*
	fmt.Printf("result %v\n", arrays.ProcessQueries([]int{4,1,2,2}, 5))
*/
