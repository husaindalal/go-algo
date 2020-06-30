package arrays

import (
	"fmt"
	"math"
)

func ColumnWith1(arr [][]int) int {
	first1 := math.MaxInt8

	for i := 0; i < len(arr); i++ {
		oneIdx := binarySearch(arr[i], 0, first1)
		fmt.Printf("oneIdx %v %v %v\n", i, oneIdx, first1)
		if oneIdx < first1 {
			first1 = oneIdx
		}
	}

	return first1
}

func binarySearch(arr []int, st int, en int) int {
	if en == math.MaxInt8 {
		en = len(arr) - 1
	}
	if arr[en] == 0 {
		return math.MaxInt8
	}

	oneIdx := math.MaxInt8
	for st < en {
		mid := (st + en) / 2
		if arr[mid] == 1 {
			oneIdx = mid
			en = mid - 1
		} else {
			st = mid + 1
		}
	}

	return oneIdx
}

/*
	if st >= en {
		if st == 1 {
			return st
		} else {
			return math.MaxInt8
		}
	}

	//oneIdx := math.MaxInt8
	mid := (st + en) / 2

	//if mid > len(arr) {
	//	return
	//}

	if arr[mid] == 1 {
		mid = binarySearch(arr, st, mid-1)
	} else {
		mid = binarySearch(arr, mid+1, en)
	}

	return mid
*/
