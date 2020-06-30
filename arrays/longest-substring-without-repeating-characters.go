package arrays

/*
two pointers
st and en

c = []int, 26
maxlen = 0

for st < n && en < n

	if c[en-'a'] < 1 {
		c[en-'a']++
		en++
	}
	if c[en-'a'] >= 1 {
		c[en-'a']++?
		if maxlen < en-st {
			maxlen = en-st
		}
		en++
		for c[en-'a'] >= 1 {
			c[st-'a'] --
			st++
		}
	}


*/

func LengthOfLongestSubstring(s string) int {
	return 0
}
