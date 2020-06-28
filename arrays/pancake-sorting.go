package arrays

import (
	"fmt"
	"math"
)

func PancakeSort(A []int) []int {
	sortedIndex := len(A) - 1
	result := []int{}

	for sortedIndex > 0 {
		maxIndex := findMaxIndex(A, sortedIndex)
		fmt.Printf("%v %v %v\n", sortedIndex, maxIndex, A)
		if maxIndex == sortedIndex {
			sortedIndex--
			continue
		}
		flip(A, maxIndex)
		//fmt.Printf("%v %v\n", maxIndex, A)
		flip(A, sortedIndex)
		//fmt.Printf("%v %v\n", sortedIndex, A)
		result = append(result, maxIndex+1, sortedIndex+1)
		sortedIndex--
	}

	return result
}

func findMaxIndex(A []int, uptoIndex int) int {
	maxVal := math.MinInt32
	maxIndex := uptoIndex
	for i := uptoIndex; i >= 0; i-- {
		if A[i] > maxVal {
			maxIndex = i
			maxVal = A[i]
		}
	}
	return maxIndex
}

func flip(A []int, uptoIndex int) {
	for i := 0; i <= uptoIndex/2; i++ {
		A[i], A[uptoIndex-i] = A[uptoIndex-i], A[i]
	}
}

/*
	fmt.Printf("result %v\n", arrays.PancakeSort([]int{
		3,2,4,1,
	}))

*/
