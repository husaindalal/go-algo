package arrays

/*
 Quicksort in descending order

take K-1 as the pivot
note sorting should be in descending order
pivot value = nums[k
find element less than pivot from left = increment i till nums[i] > pivotVal
find element greater than pivot from right = decrement j till nums[j] < pivotval

*/
func FindKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, k-1)
}

func quickSort(nums []int, st int, en int, k int) int {

	pivotVal := nums[k]
	i := st
	j := en
	for i < j {
		for nums[i] > pivotVal {
			i++
		}
		for nums[j] < pivotVal {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	if i == k {
		return nums[k]
	}
	if i > k {
		return quickSort(nums, st, k, k)
	}
	return quickSort(nums, k, en, k)

}
