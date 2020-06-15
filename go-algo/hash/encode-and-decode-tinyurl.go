package hash

import "math/rand"

type TinyURL struct {
	lookup map[string]string
	size int
}

func NewTinyURL() TinyURL {
	return TinyURL{
		lookup: map[string]string{},
		size: 7,
	}
}

const values = "0123456789abcdefABCDEF"

func (t *TinyURL) Encode(long string) string {
	hash := ""
	for i := 0; i < t.size; i ++ {
		index := rand.Int() % len(values)
		hash = hash + string(values[index])
	}

	return hash
}

func (t *TinyURL) Decode(short string) string {
	return t.lookup[short]
}