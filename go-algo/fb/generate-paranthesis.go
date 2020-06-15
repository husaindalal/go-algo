package fb

// https://www.interviewbit.com/problems/generate-all-parentheses-ii/

func GenerateParenthesis(n int) map[string]bool {

	result := map[string]bool{
		"()": true,
	}

	for i := 1; i<n; i++ {
		tmpResult := map[string]bool{}
		for r := range result {
			tmpResult["()"+r] = true
			tmpResult[r+"()"] = true
			tmpResult["("+r+")"] = true
		}
		result = tmpResult
	}
	return result
}

/*
result := fb.GenerateParenthesis(4)
	fmt.Printf("factor %v\n", result)
 */

//func paraHelper(n int, tmp string) string {
//	//if n == 1 {
//	//	return "()"
//	//}
//	result := []string{}
//	for i := n; i>1; i-- {
//		 tmp = "("+ tmp+ ")"
//	}
//}