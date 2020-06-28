package stringy

import "fmt"

func QueryString(S string, N int) bool {

	results := []string{}
	for i := 0; i < len(S); i++ {
		results = genSubstrings(S, "", i, results)
	}

	fmt.Printf("results %v \n", results)

	return false
}

func genSubstrings(S string, tmp string, start int, results []string) []string {
	if start == len(S) {
		return results
	}

	for i := start + 1; i < len(S); i++ {
		tmp = tmp + string(S[start])
		results = append(results, tmp)
		results = genSubstrings(S, tmp, i, results)
		if len(tmp) > 0 {
			tmp = tmp[0 : len(tmp)-1]
		}
	}

	return results
}
