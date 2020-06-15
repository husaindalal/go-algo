package hash

import (
	"fmt"
	"math"
)

/*
	use an array of temperatures for lookup
	go reverse and save the latest index in lookup with the current temp
	search for min index for all temps > current from lookup
 */
func DailyTemperature(temps []int) []int {
	//n := len(temps)
	result := make([]int, len(temps))

	// index is temperature. value is position
	lookup := make([]int, 100)

	for i:=len(temps)-1; i>=0; i-- {
		result[i] = findClosestAbove(temps, i, lookup)
		lookup[temps[i]] = i
	}

	return result
}

func findClosestAbove(temps []int, tindex int, lookup []int) int {
	minIndex := math.MaxInt8
	//fmt.Printf("temp %v, tindex %v, minIndex %v, lookup %v\n", temps[tindex], tindex, minIndex, lookup)
	for i:= temps[tindex]; i<len(lookup); i++ {
		if lookup[i] != 0 && lookup[i] < minIndex {
			minIndex = lookup[i]
		}
	}
	fmt.Printf("temp %v, tindex %v, minIndex %v, lookup %v\n", temps[tindex], tindex, minIndex, lookup)
	if minIndex == math.MaxInt8 {
		return 0
	}
	return minIndex - tindex
}

/*
	result := hash.DailyTemperature([]int{73,74,75,71,69,72})
	fmt.Printf("result %v %v %v %v\n", result, "x","y","z")

 */