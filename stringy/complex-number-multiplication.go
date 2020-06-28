package stringy

import (
	"strconv"
	"strings"
)

// (a+ib)×(x+iy)=ax+i^2 by+i(bx+ay)=ax−by+i(bx+ay)

func ComplexNumberMultiply(a string, b string) string {
	// validate a and b

	r1, _ := strconv.Atoi(strings.Split(a, "+")[0])
	i1str := strings.Split(a, "+")[1]
	i1, _ := strconv.Atoi(strings.Split(i1str, "i")[0])

	r2, _ := strconv.Atoi(strings.Split(b, "+")[0])
	i2str := strings.Split(b, "+")[1]
	i2, _ := strconv.Atoi(strings.Split(i2str, "i")[0])

	result := strconv.Itoa(r1*r2-i1*i2) + "+" + strconv.Itoa(r1*i2+r2*i1) + "i"

	return result

}

/*
fmt.Printf("result %v\n", stringy.ComplexNumberMultiply("1+-1i", "1+-1i"))
*/
