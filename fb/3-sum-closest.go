package fb

import (
	"math"
	"sort"
)

// https://www.interviewbit.com/problems/3-sum/

/*
  -4, 3, 0, 6, 11, -3, 3, 8
   0  1  2  3   4   5  6  7

  6 (3, 3, 0)
  -4 (-4, 3, -3) / (4, -4, 8)

  -4 -> 0
  3 -> 1


*/

func ThreeSum(arr []int, s int) (bool, int, int, int) {
	//lookup := map[int]int{}

	//for i, a := range arr {
	//	lookup[a] = i
	//}

	x := math.MinInt32
	y := math.MinInt32
	z := math.MinInt32
	found := false
	for i, a := range arr {
		found, x, y = twoSum(arr, i, s-a)
		if found {
			z = arr[i]
			break
		}
	}
	return found, x, y, z
}

func twoSum(arr []int, i int, sum int) (bool, int, int) {
	set := map[int]bool{}
	for j := i + 1; j < len(arr); j++ {
		if set[sum-arr[j]] {
			return true, arr[j], sum - arr[j]
		}
		set[j] = true
	}
	return false, math.MinInt32, math.MinInt32
}

// with duplicates
// after sorting two pointers are front and back.
// move only one pointer at a time depending on < sum or >
// dont forget absolute compare
func ThreeSumClosest(arr []int, sum int) (int, int, int, int) {

	// sort
	sort.Ints(arr)

	n := len(arr) - 1
	x := math.MinInt32
	y := math.MinInt32
	z := math.MinInt32
	closest := math.MaxInt8
	// iterate through first loop n-2 times
	for i := 0; i < n-2; i++ {
		j := i + 1
		k := n
		for j < k {
			diff := sum - arr[i] - arr[j] - arr[k]
			if diff == 0 {
				return closest, arr[i], arr[j], arr[k]
			}
			if int(math.Abs(float64(diff))) < int(math.Abs(float64(closest))) {
				closest = diff
				x = arr[i]
				y = arr[j]
				z = arr[k]
			}
			if diff < 0 {
				k--
			} else {
				j++
			}
		}
	}
	// two pointers
	return closest, x, y, z
}

/*
	result, x, y, z := fb.ThreeSumClosest([]int{-1, -1, 0, 3, 3, 3, 3}, -17)
	fmt.Printf("factor %v %v %v %v\n", result, x, y, z)

*/
