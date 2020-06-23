package graph

import (
	"fmt"
	"sort"
)

type Tuple struct {
	value int
	power int
}

func GetKthPowerValue(lo int, hi int, k int) int {

	// validate, lo < hi, k < hi - lo

	// find the power of each elem from lo to hi
	// for optimization, save the found powers in map
	// sort by stable power and save in array
	// get the kth index
	tuples := []Tuple{}
	hash := map[int]int{}
	for i := lo; i <= hi; i++ {
		tuples = append(tuples, Tuple{
			value: i,
			power: getPower(i, hash),
		})
	}

	fmt.Printf("beforesort %v \n", tuples)

	sort.SliceStable(tuples, func(i, j int) bool {
		if tuples[i].power < tuples[j].power {
			return true
		} else if tuples[i].power > tuples[j].power {
			return false
		} else if tuples[i].value < tuples[j].value {
			return true
		} else {
			return false
		}
	})
	fmt.Printf("aftersort %v \n", tuples)

	return tuples[k-1].value

}

func getPower(i int, hash map[int]int) int {
	if i <= 1 {
		return 1
	}
	fmt.Printf("getPower %v, %v \n", i, hash)
	pow, present := hash[i]
	if present {
		fmt.Printf("getPower inside %v, %v \n", i, hash[i])
		return hash[i]
	}

	if i%2 == 0 {

		pow = getPower(i/2, hash)
		fmt.Printf("even %v, %v \n", i/2, pow)
	} else {

		pow = getPower(i*3+1, hash)
		fmt.Printf("odd %v, %v \n", i*3+1, pow)
	}
	hash[i] = pow + 1

	return pow + 1
}

/*

	fmt.Printf("result %v\n", graph.GetKthPowerValue(12, 15, 2))

	fmt.Printf("result %v\n", graph.GetKthPowerValue(1, 1, 1))

*/
