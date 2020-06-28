package basic

import (
	"fmt"
)

func ListPermutations(arr []string) [][]string {
	for i := range arr {
		permutationUtil(arr, i)
	}

	return nil
}

func permutationUtil(arr []string, i int) {
	if i == len(arr)-1 {
		fmt.Printf("%v", arr)
		return
	}
	for j := i + 1; j < len(arr); j++ {
		// permute
		arr[i], arr[j] = arr[j], arr[i]

		permutationUtil(arr, j)
		// unpermute
		arr[j], arr[i] = arr[i], arr[j]
	}
}
