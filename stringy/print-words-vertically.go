package stringy

import "strings"

func PrintVertically(s string) []string {
	strs := strings.Split(s, " ")
	maxLen := 0
	for _, str := range strs {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}

	result := make([]string, len(strs))

	for _, str := range strs {
		for j := 0; j < maxLen; j++ {
			if j < len(str) {
				result[j] = result[j] + string(str[j])
			} else {
				result[j] = result[j] + " "
			}

		}
	}

	return result
}

/*
fmt.Printf("result %v\n", stringy.PrintVertically("A BC DEF GH"))
*/
