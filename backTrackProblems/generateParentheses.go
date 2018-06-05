package main

/*
ProblemStatment:
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

*/
func generateParenthesis(n int) []string {
	result := []string{}
	tmp := ""
	helper(&result, tmp, 0, 0, n)
	return result
}
func helper(res *[]string, tStr string, openP, closeP, max int) {
	if len(tStr) == 2*max && isValid(tStr) {
		tmp := make([]byte, len(tStr))
		copy(tmp, tStr)
		*res = append(*res, string(tmp))
		return
	} else {
		if openP < max {
			helper(res, tStr+"(", openP+1, closeP, max)
		}
		if closeP <= openP {
			helper(res, tStr+")", openP, closeP+1, max)
		}
	}
}
func isValid(s string) bool {
	count := 0
	for _, c := range s {
		if c == '(' {
			count++
		} else {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}
