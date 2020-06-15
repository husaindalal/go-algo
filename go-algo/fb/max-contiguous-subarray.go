package fb

// Note the sum can be negative
// https://www.interviewbit.com/problems/max-sum-contiguous-subarray/
// https://www.geeksforgeeks.org/maximum-subarray-sum-using-divide-and-conquer-algorithm/

func MaxContiguousSubarray(array []int) int {

	return maxSubarray(array, 0, len(array)-1)

}

func maxSubarray(array []int, st, en int) int {
	if st == en {
		return array[st]
	}

	// note the formula
	mid := (en + st) / 2

	return max3(maxSubarray(array, st, mid), maxSubarray(array, mid+1, en), maxSubarrayMiddle(array, st, mid, en))
}

func maxSubarrayMiddle(array []int, st, mid, en int) int {
	// leftSum := math.MinInt64
	// rightSum := max right
	// total = leftSum + rightSum || max(leftSum, rightSum)
	return 0
}

func max3(x, y, z int) int {
	return max(max(x, y), z)
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*


	result := fb.MaxContiguousSubarray([]int {
		-5, 3, -2, 2, 5, -2, -6,
	})

	fmt.Printf("Result %v \n", result)
*/
