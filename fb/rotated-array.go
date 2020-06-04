package fb

import (
	"fmt"
	"math"
)

// assume sorted
func GetRotatedIndex(arr []int) int {
	st := 0
	en := len(arr) - 1
	result := math.MinInt32
	for en-st > 1 { //&& st >= 0 && en <= len(arr)-1

		mid := (st + en) / 2
		fmt.Printf("st %d mid %d en %d arr %d %d %d || \n", st, mid, en, arr[st], arr[mid], arr[en])

		if arr[st] >= arr[mid] {
			en = mid
		} else if arr[mid] >= arr[en] {
			st = mid
		}

	}
	fmt.Printf("st %d en %d arr %d %d || \n", st, en, arr[st], arr[en])
	if st == en {
		return st
	}
	if en-st == 1 {
		if arr[st] <= arr[en] {
			return st
		}
		return en
	}

	return result
}

/*
	fmt.Printf("rearranged %v \n", fb.GetRotatedIndex([]int{3,0, 0,1,4}))
	fmt.Printf("rearranged %v \n", fb.GetRotatedIndex([]int{12,7,8,9,10, 11}))

	fmt.Printf("rearranged %v \n", fb.GetRotatedIndex([]int{2,3,0, 0,1,4}))
	fmt.Printf("rearranged %v \n", fb.GetRotatedIndex([]int{11,12,7,8,9,10, 11}))
	//fmt.Printf("rearranged %v \n", fb.SearchRotatedArray([]int{11, 12, 13, -3, 0, 1, 2}, 2, 3))

*/

// even
// odd
// duplicate mid

//fmt.Printf("rearranged %v ", fb.GetRotatedIndex([]int{2,3,0, 0,1,4}))
//fmt.Printf("rearranged %v ", fb.GetRotatedIndex([]int{11,12,7,8,9,10, 11}))

// split array into 2 by rotatedIndex and search in each
func SearchRotatedArray(arr []int, searchVal int, rotatedIndex int) int {
	n := len(arr) - 1

	if arr[rotatedIndex] == searchVal {
		return rotatedIndex
	}
	result := binarySearch(arr, rotatedIndex, n, searchVal)
	if result > -1 {
		return result
	}
	return binarySearch(arr, 0, rotatedIndex-1, searchVal)
}

func binarySearch(arr []int, st, en, searchVal int) int {

	for st <= en {
		mid := (en + st) / 2
		if arr[mid] == searchVal {
			return mid
		}
		if searchVal > arr[mid] {
			st = mid + 1
		} else {
			en = mid - 1
		}
	}
	return -1
}

// Returns index or -1
//func SearchRotatedArray(arr []int, searchVal int, rotatedIndex int) (int) {
//	r := rotatedIndex
//	n := len(arr) -1
//
//	st := rotatedIndex
//	en := (r -1) % n
//
//	fmt.Printf("r %d, n %d, st %d, en %d \n", r, n, st, en)
//
//	for st != en {
//		mid := (st + en) / 2
//		mid = (mid - r) % n
//
//		fmt.Printf("st %d, mid %d, en %d \n", st, mid, en)
//
//		if arr[mid] == searchVal {
//			return mid
//		}
//		if searchVal > arr[mid]  {
//			st = (mid + 1 ) % n
//		} else {
//			en = (mid - 1) % n
//		}
//	}
//
//	return -1
//}
