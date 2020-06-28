package stringy

// https://golangcookbook.com/chapters/strings/processing/

func FindAndReplacePattern(words []string, pattern string) []string {
	results := []string{}
	// validate pattern len 0, word len 0.
	// normalize chars to lowercase

	for _, word := range words {
		pwmap := map[byte]byte{}
		wpmap := map[byte]byte{}
		patternMatch := true
		for i := 0; i < len(word); i++ {
			pw, present := pwmap[pattern[i]]
			if present {
				if pw != word[i] {
					patternMatch = false
					break
				}
			} else {
				pwmap[pattern[i]] = word[i]
			}

			wp, present := wpmap[word[i]]
			if present {
				if wp != pattern[i] {
					patternMatch = false
					break
				}
			} else {
				wpmap[word[i]] = pattern[i]
			}

		}
		if patternMatch {
			results = append(results, word)
		}
	}

	return results
}

/*
fmt.Printf("result %v\n", stringy.FindAndReplacePattern([]string{"abc","deq","mee","aqq","dkd","ccc"}, "abb"))
*/
