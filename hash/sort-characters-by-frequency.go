package hash

import "sort"

type touple struct {
	char  byte
	count int
}

func SortCharByFreq(str string) string {
	counts := map[byte]int{} //

	for i := 0; i < len(str); i++ {
		counts[str[i]]++ // TODO make case insensitive
	}
	touples := []touple{}
	for k, v := range counts {
		touples = append(touples, touple{
			char:  k,
			count: v,
		})
	}

	// sort by count
	sort.SliceStable(touples, func(i, j int) bool {
		return touples[i].count > touples[j].count
	})

	res := []byte{}
	for i := 0; i < len(touples); i++ {
		for j := 0; j < touples[i].count; j++ {
			res = append(res, touples[i].char)
		}
	}
	return string(res)
}

/*
	result := hash.SortCharByFreq("constantinople")
	fmt.Printf("result %v %v %v %v\n", result, "x","y","z")

*/
