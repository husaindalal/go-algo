package arrays

import "fmt"

func NumOfSubarraysGreaterThanAvg(arr []int, k int, threshold int) int {
	// two pointers
	// initialize window with avg from 0 to k
	//

	// validate k is positive && k >= len(arr)

	st := 0
	en := k - 1
	avg := 0
	result := 0
	for i := st; i <= en; i++ {
		avg = avg + arr[i]
	}
	avg = avg / k

	fmt.Printf("%v %v %v %v \n", st, en, avg, result)
	if avg > threshold {
		result++
	}

	st++
	en++
	for en < len(arr) { // off by 1 issue
		//fmt.Printf("%v %v %v %v \n", arr[st-1], arr[en], avg*k, en-st)
		avg = (avg*k - arr[st-1] + arr[en]) / k // overflow
		if avg > threshold {
			result++
		}
		fmt.Printf("%v %v %v %v \n", st, en, avg, result)
		st++
		en++
	}
	return result
}

//fmt.Printf("result %v\n", arrays.NumOfSubarraysGreaterThanAvg([]int{3,2,4,1}, 2, 1))
//fmt.Printf("result %v\n", arrays.NumOfSubarraysGreaterThanAvg([]int{1,1,-1,-1}, 2, -2))
//fmt.Printf("result %v\n", arrays.NumOfSubarraysGreaterThanAvg([]int{}, 2, 1))
//fmt.Printf("result %v\n", arrays.NumOfSubarraysGreaterThanAvg([]int{3,2,4,1}, 2, 1))
