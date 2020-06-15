package fb

import (
	"fmt"
	"math"
)

// https://www.interviewbit.com/problems/best-time-to-buy-and-sell-stocks-i/

/*

	6,3,9,2,4,9,4,1,8
    0 1 2 3 4 5 6 7 8


// can be a loss
 */

func MaxProfit2(arr []int) (int, int, int) {
	b := 0
	s := 1

	max := math.MinInt32
	maxBuy := b
	maxSell := s
	for b < len(arr) && s < len(arr){
		// increment sell ptr until next is greater than now
		fmt.Printf("first %v %v %v \n", s, arr[s], -1)
		for s < len(arr)-1 && arr[s+1] > arr[s] {
			s = s + 1
		}
		// increment buy ptr until profit > maxProfit
		fmt.Printf("first after %v %v %v \n", s, arr[s], -1)
		fmt.Printf("second %v %v %v \n", b, arr[b], -1)
		for b < s && arr[s] - arr[b] > max {
			max = arr[s] - arr[b]
			maxBuy = b
			maxSell = s
			b = b + 1
		}
		fmt.Printf("second after %v %v %v \n", b, arr[b], -1)
		s ++
		//b ++

	}
	return max, maxBuy, maxSell
}

func MaxProfit(arr []int) (int, int, int) {
	n := len(arr) - 1
	b := 0
	s := n

	maxProfit := math.MinInt32
	maxBuy := b
	maxSell := s

	toggle := true
	for b < s {
		for s-1 > b && arr[s-1] > arr[s] {
			s--
		}

		for b+1 < s && arr[b+1] < arr[b] {
			b++
		}

		profit := arr[s] - arr[b]
		if profit > maxProfit {
			maxProfit = profit
			maxBuy = b
			maxSell = s
		}

		if toggle {
			s--
		} else {
			b++
		}
		toggle = !toggle

	}

	return maxProfit, maxBuy, maxSell

}