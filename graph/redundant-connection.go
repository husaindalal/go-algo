package graph

// https://leetcode.com/problems/redundant-connection/

func FindRedundantConnection(edges [][]int) []int {
	uf := NewUnionFind(len(edges) * len(edges[0]))

	result := []int{}
	for i := 0; i < len(edges); i++ {
		newEdge := uf.Union(edges[i][0], edges[i][1])
		if !newEdge {
			result = append(result, edges[i][0], edges[i][1])
		}
	}

	return result
}

type UnionFind struct {
	root []int
	size []int
}

func NewUnionFind(size int) *UnionFind {
	uf := &UnionFind{
		root: make([]int, size),
		size: make([]int, size),
	}
	for i := 0; i < size; i++ {
		uf.root[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) GetRoot(p int) int {
	r := p
	for uf.root[r] != r {
		//collapse
		uf.root[r] = uf.root[uf.root[r]]
		r = uf.root[r]
	}
	return r
}

func (uf *UnionFind) Union(u int, v int) bool {
	ru := uf.GetRoot(u)
	rv := uf.GetRoot(v)

	if ru == rv {
		return false
	}

	if uf.size[ru] > uf.size[rv] {
		uf.root[rv] = ru
		uf.size[ru] = uf.size[ru] + uf.size[rv]
	} else {
		uf.root[ru] = rv
		uf.size[rv] = uf.size[rv] + uf.size[ru]
	}
	return true
}

func (uf *UnionFind) Find(p int) int {
	return uf.GetRoot(p)
}

func (uf *UnionFind) Connected(u int, v int) bool {
	return uf.GetRoot(u) == uf.GetRoot(v)
}

/*

	fmt.Printf("result %v\n", graph.FindRedundantConnection([][]int {
		{1, 2},
		{1,3},
		{3,4},
		{1,4},
		{1,5},
	}))
*/
