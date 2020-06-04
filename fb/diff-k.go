package fb

import "math"

// https://www.interviewbit.com/problems/diffk/

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
