package hash

import "strings"

func ReplaceWords(dict []string, sentence string) string {
	set := map[string]bool{}
	for _, d := range dict {
		set[d] = true
	}

	splits := strings.Split(sentence, " ")
	newSentence := ""
	for _, word := range splits {

		for i :=2; i< len(word); i++ {
			prefix := word[0:i]
			if set[prefix] {
				word = prefix
				break
			}
		}
		newSentence = newSentence + " " + word
	}
	return newSentence[1:]
}

/*

	result := hash.ReplaceWords([]string{"car", "cat", "rat", "bat"}, "the cattle was rattled by the battery")

	fmt.Printf("result %v %v %v %v\n", result, "x","y","z")

 */
