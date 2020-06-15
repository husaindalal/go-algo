package fb

//https://www.interviewbit.com/problems/sort-by-color/

// only values are -1, 0, 1
func SortByColor(arr []int) []int {
	// validate TODO
	n := 0 // negative end
	p := len(arr) -1 // positive start

	for i := n; i <= p; i++ {
		if arr[i] == -1 {
			arr[i], arr[n] = arr[n], arr[i]
			n++
		} else if arr[i] == 1 {
			arr[i], arr[p] = arr[p], arr[i]
			p--
		}
	}

	return arr
}

/*
result := fb.SortByColor([]int{1, 1, 1, 0, 0, 0, -1, -1, -1})
	fmt.Printf("factor %v %v %v %v\n", result, 1,1,1)
 */