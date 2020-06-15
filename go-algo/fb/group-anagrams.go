package fb

import (
	"sort"
	"strings"
)

func GroupAnagrams(words []string) [][]string {
	// validate TODO

	result := map[string][]string{}

	for _, word := range words {
		hash := getWordHash(word)
		result[hash] = append(result[hash], word)
	}

	return convertMapToMatrix(result)

}

func getWordHash(word string) string {
	word = strings.ToLower(word) // TODO

	bytes := []byte(word)
	ints := []int{}
	for _, b := range bytes {
		ints = append(ints, int(b))
	}
	sort.Ints(ints)

	sortedBytes := []byte{}
	for _, i := range ints {
		sortedBytes = append(sortedBytes, byte(i))
	}

	return string(sortedBytes)
}

func convertMapToMatrix(mymap map[string][]string) [][]string {
	result := [][]string{}
	for m := range mymap {
		result = append(result, mymap[m])
	}

	return result
}

/*
	fmt.Printf("anagrams %v \n", fb.GroupAnagrams([]string{"Eat", "tea", "tan", "ate", "nat", "bat"}))

*/
