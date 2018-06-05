package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*
ProblemStatment:
A message containing letters from A-Z is being encoded to numbers using the following mapping:

'A' -> 1
'B' -> 2
...
'Z' -> 26
Given a non-empty string containing only digits, determine the total number of ways to decode it.

Example 1:

Input: "12"
Output: 2
Explanation: It could be decoded as "AB" (1 2) or "L" (12).
Example 2:

Input: "226"
Output: 3
Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2 6).
*/
func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	} else {
		dp[1] = 0
	}
	for i := 2; i <= len(s); i++ {
		first, _ := strconv.Atoi(string(s[i-1 : i]))
		second, _ := strconv.Atoi(string(s[i-2 : i]))
		if first >= 1 && first <= 9 {
			dp[i] += dp[i-1]
		}
		if second >= 10 && second <= 26 {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
