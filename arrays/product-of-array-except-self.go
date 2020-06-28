package arrays

import "fmt"

func ProductExceptSelf(nums []int) []int {
	// todo validate

	n := len(nums)

	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return nums
	}

	leftMultiplier := make([]int, n)
	rightMultiplier := make([]int, n)

	leftMultiplier[0] = nums[0]

	for i := 1; i < n; i++ {
		leftMultiplier[i] = leftMultiplier[i-1] * nums[i]
	}

	fmt.Printf("leftMultiplier %v\n", leftMultiplier)
	//prod = nums[n-1]
	rightMultiplier[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMultiplier[i] = rightMultiplier[i+1] * nums[i]
	}
	fmt.Printf("rightMultiplier %v\n", rightMultiplier)

	result := make([]int, n)
	result[0] = rightMultiplier[1]
	result[n-1] = leftMultiplier[n-2]
	for i := 1; i < n-1; i++ {
		result[i] = leftMultiplier[i-1] * rightMultiplier[i+1]
	}

	return result
}

/*
fmt.Printf("result %v\n", arrays.ProductExceptSelf([]int{1,2,3,4}))
*/

//func ProductExceptSelf(nums []int) []int {
//	// todo validate
//	first := nums[0]
//	//last := nums[len(nums)-1]
//	prod := first
//
//	result := make([]int, len(nums))
//	result[0] = first
//
//	for i:=1; i<len(nums); i++ {
//		prod = prod * nums[i]
//		if i+1 < len(nums) {
//			result[i+1] = prod
//		}
//	}
//
//	fmt.Printf("%v \n", result)
//
//	return result
//}
