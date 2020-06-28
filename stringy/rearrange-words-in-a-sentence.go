package stringy

import (
	"sort"
	"strings"
)

func ArrangeWords(text string) string {
	result := ""

	splits := strings.Split(strings.ToLower(text), " ")
	sort.SliceStable(splits, func(i, j int) bool {
		return len(splits[i]) < len(splits[j])
	})
	for i, str := range splits {
		if i == 0 {
			str = strings.ToUpper(string(str[0])) + str[1:]
		}
		result = result + str
		if i < len(splits)-1 {
			result = result + " "
		}
	}

	return result
}

// fmt.Printf("result %v\n", stringy.ArrangeWords("&amp; is an HTML entity but &ambassador; is not"))
