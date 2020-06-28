package stringy

import "fmt"

func IntToRoman(num int) string {
	vals := []int{10, 9, 5, 4, 1}
	roman := []string{"X", "IX", "V", "IV", "I"}

	result := ""
	//for num > 0 {
	for i, val := range vals {
		for num >= val {
			fmt.Printf("%v %v %v\n", i, val, num)
			num = num - val
			result = result + roman[i]

		}
	}
	//}

	return result
}

//fmt.Printf("result %v\n", stringy.IntToRoman(12))

/*
	// validate negative

	lookup := map[int]string{
		1 : "I",
		4 : "IV",
		5: "V",
		9: "IV",
		10: "X",
	}

	divisor := 10

	result := ""

	for num > 0 {
		val := int (num / divisor)
		val = val * divisor
		fmt.Printf("val %v %v %v \n", val, num, divisor)
		result = result + findRoman(val, lookup)

		num = num % divisor
		divisor = divisor / 10

	}
	return result
}

func findRoman(val int, lookup map[int]string) string {
	if lookup[val] != "" {
		return lookup[val]
	}
	key := 0
	for k := range lookup {
		if k < val {
			key = k
		}
	}

	for i:=0; i< val; i++ {

	}
*/
