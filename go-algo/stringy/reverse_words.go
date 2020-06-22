package stringy

import "fmt"

// to err is human
// ot rre si namuh

/*
to err is human
  i

*/
// assumption space is the only character
func ReverseWords(str string) string {
	result := ""
	wordStart := 0
	for i, chr := range str {
		if chr == ' ' {
			fmt.Printf("chr %v wordstart %v, i %v result %v \n", chr, wordStart, i, result)
			result = result + " " + reverse(str, wordStart, i-1)
			wordStart = i + 1
		}
	}

	if wordStart > 0 {
		result = result + " " + reverse(str, wordStart, len(str)-1)
	}

	return result
}

func reverse(str string, st, en int) string {
	result := ""
	for i := en; i >= st; i-- {
		result = result + string(str[i])
	}

	return result
}

// to err is human
// human is err to
/*
to err is human
  i         i
*/

func ReverseSentence(str string) string {
	wordStart := 0
	result := ""
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			result = str[wordStart:i] + " " + result
			wordStart = i + 1
		}
	}

	result = str[wordStart:len(str)] + " " + result

	return result
}

//func main() {
//	//stringy.ListPermutations([]string{"A", "B", "C", "D"})
//	//fmt.Printf("ReverseWords %v", stringy.ReverseWords("to err is human"))
//	//fmt.Printf("ReverseWords %v", stringy.ReverseWords(" to err is human "))
//
//	fmt.Printf("ReverseSentence %v \n", stringy.ReverseSentence("to err is human"))
//	fmt.Printf("ReverseSentence %v \n", stringy.ReverseSentence(" to err is human "))
//}
