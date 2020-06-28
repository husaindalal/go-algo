package arrays

import "fmt"

// 2,1,3,5,4
func NumTimesAllBlue(light []int) int {
	maxBlue := -1
	maxOn := -1
	result := 0

	bulbs := make([]int, len(light))

	for _, l := range light {
		l = l - 1
		bulbs[l] = 1

		fmt.Printf("first %v %v\n", l, maxBlue)

		if l == maxBlue+1 {
			//maxBlue = l

			for i := l; i < len(light); i++ {
				if bulbs[i] == 1 {
					maxBlue++
				} else {
					break
				}
			}
			fmt.Printf("inside %v %v\n", l, maxBlue)
		}

		fmt.Printf("maxOn %v %v\n", l, maxOn)
		if l > maxOn {
			maxOn = l
		}

		if maxOn == maxBlue {
			result++
		}

	}

	return result
}

//fmt.Printf("result %v\n", arrays.NumTimesAllBlue([]int{2,1,3,5,4}))

//fmt.Printf("result %v\n", arrays.NumTimesAllBlue([]int{3,2,4,1,5}))
//fmt.Printf("result %v\n", arrays.NumTimesAllBlue([]int{2,1,4,3,6,5}))
