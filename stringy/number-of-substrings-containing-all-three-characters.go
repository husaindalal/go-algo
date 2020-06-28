package stringy

import "fmt"

func NumberOfSubstringsWithABC(s string) int {
	result := 0

	st := 0
	en := 0

	aCount := 0
	bCount := 0
	cCount := 0

	for st < len(s) {
		// not found
		for (aCount == 0 || bCount == 0 || cCount == 0) && st < len(s) {
			if s[st] == 'a' {
				aCount++
			}
			if s[st] == 'b' {
				bCount++
			}
			if s[st] == 'c' {
				cCount++
			}

			st++

			fmt.Printf(" st en %v %v \n", st, en)
			fmt.Printf(" counts %v %v %v\n", aCount, bCount, cCount)
		}
		result = result + en

		// found
		for en < st && aCount > 0 && bCount > 0 && cCount > 0 {
			result++
			if s[en] == 'a' {
				aCount--
			}
			if s[en] == 'b' {
				bCount--
			}
			if s[en] == 'c' {
				cCount--
			}
			en++
			fmt.Printf(" en st %v %v \n", st, en)
			fmt.Printf(" counts %v %v %v\n", aCount, bCount, cCount)
		}

	}

	return result
}

//
//func NumberOfSubstringsWithABC2(s string) int {
//	result := 0
//
//	en := 0
//	hashmap := map[byte]int{
//		'a':0,
//		'b':0,
//		'c':0,
//	}
//
//	for st:=0; st<len(s); st++ {
//		hashmap[s[st]] ++
//		for hashmap['a'] > 0 && hashmap['b']>0 && hashmap['c']>0 {
//
//			en++
//		}
//	}
//}
