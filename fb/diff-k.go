package fb

import "math"

// https://www.interviewbit.com/problems/diffk/

// two pointers same direction
// assumes sorted
// make sure to check pointers are not same location
// and check absolute difference is K
func DiffK2(arr []int, k int) (int, int) {
	st := 0
	en := 1
	n := len(arr) - 1

	for st < n && en < n {
		if st != en && (arr[en]-arr[st] == k || arr[st]-arr[en] == k) {
			return st, en
		}

		if arr[en]-arr[st] < k {
			en++
		} else {
			st++
		}
	}

	return -1, -1
}

/*
	st, en := fb.DiffK2([]int{6,3,9,2,4,7,4,1,8}, 2)
	fmt.Printf("factor %v %v %v\n", st, en, -1)
*/

// binary search
func DiffKBinary(arr []int, k int) int {
	for i, val := range arr {
		find := k + val
		pair := search(arr, i+1, find)
		if pair != math.MinInt32 {
			return pair
		}
	}
	return 0
}

func search(arr []int, start int, find int) int {
	// binary search
	return math.MinInt32
}

// incorrect solution
func DiffK(arr []int, k int) (int, int) {
	st := 0
	en := len(arr) - 1

	// check for odd even
	for st < en {
		if arr[en]-arr[st] == k {
			return st, en
		}
		if arr[en]-arr[st] > k {
			st = st + 1
		} else {
			// arr[en] - arr[st] > k
			en = en - 1
		}
	}
	return 0, 0
}
