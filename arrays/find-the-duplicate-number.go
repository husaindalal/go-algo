package arrays

func findDuplicate(nums []int) int {
	slow := 0
	fast := nums[slow]

	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for slow != fast {
		fast = nums[fast]
		slow = nums[slow]
	}
	return fast

}
