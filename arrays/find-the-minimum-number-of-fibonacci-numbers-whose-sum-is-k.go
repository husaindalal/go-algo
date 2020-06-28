package arrays

import "fmt"

func FindMinFibonacciNumbers(k int) int {

	// generate fibonacci upto k and add to hashMap
	// take biggest number < k, next num to find is k - big
	// rinse and repeat

	//hashmap := map[int]bool{1 : true}
	fibs := []int{}
	f := 1
	s := 1
	for s <= k {
		fibs = append(fibs, s)
		num := f + s
		f = s
		s = num

	}

	result := 1

	v := fibs[len(fibs)-1]
	k = k - v
	fmt.Printf("%v %v %v \n", k, v, fibs)
	for k > 0 {

		v = findLessThan(fibs, k)
		k = k - v

		fmt.Printf("%v %v %v \n", k, v, fibs)
		result++
	}

	return result
}

func findLessThan(fibs []int, k int) int {
	st := 0
	en := len(fibs)
	mid := (en + st) / 2
	for en-st > 1 {
		if fibs[mid] == k {
			return fibs[mid]
		} else if fibs[mid] < k {
			st = mid
		} else {
			en = mid
		}
		mid = (en + st) / 2

	}
	return fibs[st]
}
