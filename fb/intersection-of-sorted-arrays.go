package fb

import "fmt"

// https://www.interviewbit.com/problems/intersection-of-sorted-arrays/

// both arrays are sorted
func Intersection(a []int, b []int) []int {
	i := 0
	j := 0
	result := []int{}
	for i < len(a) && j < len(b) {
		fmt.Printf("start %v %v %v %v\n", i, j, a[i], b[j])
		if a[i] == b[j] {
			result = append(result, a[i])
			i++
			j++
		}
		//fmt.Printf("mid %v %v %v %v\n", i, j, a[i], b[j])
		for i < len(a) && j < len(b) && a[i] < b[j] {
			i++
		}
		//fmt.Printf("end %v %v %v %v\n", i, j, a[i], b[j])
		for i < len(a) && j < len(b) && b[j] < a[i] {
			j++
		}

	}
	return result
}

/*
result := fb.Intersection([]int{-1, -1, 0, 3,5,6,6}, []int{-5,-1,-1, 0, 3,5,9})
	fmt.Printf("factor %v %v %v\n", result, -1, -1)
*/
