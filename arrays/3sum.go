package arrays

import (
	"fmt"
	"sort"
)

func ThreeSum(nums []int) [][]int {
	result := [][]int{}

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			fmt.Printf("top %v %v %v %v %v %v \n", i, j, k, nums[i], nums[j], nums[k])
			if nums[i]+nums[j]+nums[k] == 0 {
				fmt.Printf("inside ==\n")
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			}
			if nums[i]+nums[j]+nums[k] > 0 {
				fmt.Printf("inside >\n")
				k--
				for j < k && nums[k] == nums[k-1] {
					fmt.Printf("inside >\n")
					k--
				}
			} else {
				j++
				for j < k && nums[j] == nums[j+1] {
					fmt.Printf("inside <")
					j++
				}

			}
		}
	}

	return result
}
