package hash

// https://leetcode.com/problems/find-smallest-common-element-in-all-rows/

/*
Note nested maps are requited for duplicates
map[value]m	ap[rownum]bool
*/
func FindSmallestCommon(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return -1
	}
	hash := map[int]map[int]bool{} // map[value]map[rownum]bool
	for j := 0; j < len(matrix[0]); j++ {
		for i := 0; i < len(matrix); i++ {
			value := matrix[i][j]
			if hash[value] == nil {
				hash[value] = map[int]bool{}
			}
			hash[value][i] = true

			if len(hash[value]) == len(matrix) {
				return value
			}
		}

	}
	return -1
}

/*
	smal := hash.FindSmallestCommon([][]int {
		{1,2,3,4},
		{3,2,5,6},
		{2,3,4,5},
		{2,2,2,2},
	})
	fmt.Printf("result %v %v %v %v\n", smal, "x","y","z")
*/
