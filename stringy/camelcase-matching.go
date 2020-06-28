package stringy

import "fmt"

func CamelMatch(queries []string, pattern string) []bool {
	result := []bool{}

	for _, query := range queries {
		// extract capitals
		// i := 0
		capitals := ""
		for _, char := range query {
			if char >= 'A' && char <= 'Z' {
				capitals = capitals + string(char)

			}
		}
		fmt.Printf("%v %v \n", capitals, pattern)
		// match with pattern
		result = append(result, capitals == pattern)
	}

	return result
}

/*
fmt.Printf("result %v\n", stringy.CamelMatch([]string{"FaceBook", "FbBdWa"}, "FB"))
*/
