package hash

import "sort"

//// An Item is something we manage in a priority queue.
//type Item struct {
//	value    string // The value of the item; arbitrary.
//	priority int    // The priority of the item in the queue.
//	// The index is needed by update and is maintained by the heap.Interface methods.
//	index int // The index of the item in the heap.
//}
//
//// A PriorityQueue implements heap.Interface and holds Items.
//type PriorityQueue []*Item
//
//func (pq PriorityQueue) Len() int { return len(pq) }
//
//func (pq PriorityQueue) Less(i, j int) bool {
//	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
//	return pq[i].priority > pq[j].priority
//}
//
//func (pq PriorityQueue) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//	pq[i].index = i
//	pq[j].index = j
//}
//
//func (pq *PriorityQueue) Push(x interface{}) {
//	n := len(*pq)
//	item := x.(*Item)
//	item.index = n
//	*pq = append(*pq, item)
//}
//
//func (pq *PriorityQueue) Pop() interface{} {
//	old := *pq
//	n := len(old)
//	item := old[n-1]
//	old[n-1] = nil  // avoid memory leak
//	item.index = -1 // for safety
//	*pq = old[0 : n-1]
//	return item
//}
//
//// update modifies the priority and value of an Item in the queue.
//func (pq *PriorityQueue) update(item *Item, value string, priority int) {
//	item.value = value
//	item.priority = priority
//	heap.Fix(pq, item.index)
//}
//
//func FindKFrequent(elements []int, k int) []string {
//	items := map[string]int{
//		"banana": 3, "apple": 2, "pear": 4, "peach": 6, "plum": 3,
//	}
//	pq := make(PriorityQueue, k)
//	i := 0
//	for value, priority := range items {
//		pq[i] = &Item{
//			value:    value,
//			priority: priority,
//			index:    i,
//		}
//		i++
//	}
//	heap.Init(&pq)
//
//	// Take the items out; they arrive in decreasing priority order.
//	result := []string{}
//
//	for pq.Len() > 0 {
//		item := heap.Pop(&pq).(*Item)
//		fmt.Printf("%.2d:%s ", item.priority, item.value)
//		result = append(result, item.value)
//
//	}
//
//	return result
//}

func FindKFrequent(elements []int, size int) []int {
	// add to freq map
	freqMap := map[int]int{} // value, freq
	for _, e := range elements {
		freqMap[e]++
	}

	// invert frequency map
	invFreqMap := map[int][]int{} // freq -> values
	for k, v := range freqMap {
		if invFreqMap[v] == nil {
			invFreqMap[v] = []int{}
		}
		invFreqMap[v] = append(invFreqMap[v], k)
	}

	// sort freq map and get top k
	freqs := []int{}
	for k := range invFreqMap {
		freqs = append(freqs, k)
	}
	sort.SliceStable(freqs, func(i, j int) bool {
		return freqs[i] > freqs[j]
	})

	// extract numbers from invFreqMap. Stop when size limit is reached
	result := []int{}
	for _, f := range freqs {
		numList := invFreqMap[f]
		for _, num := range numList {
			result = append(result, num)
			if len(result) == size {
				break
			}
		}
		if len(result) == size {
			break
		}
	}

	return result
}

/*
	result := hash.FindKFrequent([]int {
		73,74,75,71,69,72,76,73,69,69,
	}, 20)


	fmt.Printf("result %v %v %v %v\n", result, "x","y","z")

*/
