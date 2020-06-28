package stringy

func EntityParser(text string) string {
	result := ""

	i := 0
	for i < len(text) {
		if text[i] == '&' {
			wlen, word := gatherWord(text, i)
			result += word
			i += wlen
		} else {
			result = result + string(text[i])
			i++

		}

	}

	return result
}

func gatherWord(text string, i int) (int, string) {
	lookup := map[string]string{
		"&quot;": "\"",
		"&apos;": "'",
		"&amp;":  "&",
	}

	word := ""
	for i < len(text) {
		word += string(text[i])
		//fmt.Printf("%v %v %v \n", word, i, string(text[i]))
		if text[i] == ';' {
			break
		}
		i++
	}
	//fmt.Printf("%v %v \n", word, lookup[word])
	if lookup[word] != "" {
		return len(word), lookup[word]
	}
	return len(word), word
}
