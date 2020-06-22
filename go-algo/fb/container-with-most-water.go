package fb

import "math"

// two pointers
// increment left if
// returns area, i, j
func ContainerWithMostWater(arr []int) (int, int,int) {
	maxArea := math.MinInt32
	mx := -1
	my := -1
	i := 0
	j := len(arr)-1
	for i < j {
		area := min(arr[i], arr[j]) * (j - i)
		if area > maxArea {
			maxArea = area
			mx = i
			my = j
		}

		if arr[i+1] > arr[i] { //boundary check
			i++
		} else if arr[j-1] > arr[j] { // boundary
			j--
		} else {
			i++
			j--
		}


	}
	return maxArea, mx, my
}

func min(x , y int) int {
	if x < y {
		return x
	}
	return y
}