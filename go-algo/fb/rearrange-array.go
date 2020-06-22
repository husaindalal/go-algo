package fb

//https://www.interviewbit.com/problems/rearrange-array/
//https://www.quora.com/How-can-I-rearrange-a-given-array-so-that-Arr-I-becomes-Arr-Arr-I-with-O-1-extra-space

func RearrangeArray(arr []int) []int {
	// trust but verify TODO
	currentIndex := 0
	currentValue := arr[0]
	zeroVal := arr[0]
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			arr[currentIndex] = zeroVal
		} else {
			arr[currentIndex] = arr[currentValue]
		}
		currentIndex = currentValue
		currentValue = arr[currentIndex]
	}
	return arr
}

// 	fmt.Printf("rearranged %v ", fb.RearrangeArray([]int{2,3,1,0,4}))
