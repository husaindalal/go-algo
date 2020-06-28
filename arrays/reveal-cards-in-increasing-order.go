package arrays

// Incorrect solution
func DeckRevealedIncreasing(deck []int) []int {
	tmp := make([]int, len(deck))
	//for i, d := range deck {
	//	tmp[i] = d
	//}

	//sort.Slice(deck, func(i, j int) bool {
	//	return deck[i] < deck[j]
	//})
	//
	//fmt.Printf(" %v\n", deck)
	//
	//i:= 0
	//j := len(deck)-1
	//k:= 0
	//for i<= j {
	//	tmp[k] = deck[i]
	//	i++
	//	k++
	//
	//	if k == len(deck) || i>j{
	//		continue
	//	}
	//	tmp[k] = deck[j]
	//	j--
	//	k++
	//}

	return tmp

}
