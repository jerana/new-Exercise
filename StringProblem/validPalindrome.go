package main

/*
Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.

Note: For the purpose of this problem, we define empty string as valid palindrome.

Example 1:

Input: "A man, a plan, a canal: Panama"
Output: true
Example 2:

Input: "race a car"
Output: false
*/
func isPalindrome(s string) bool {
	str := strings.ToLower(s)
	i := 0
	j := len(str) - 1

	for i < j {
		for i < j && (unicode.IsDigit(rune(str[i])) || unicode.IsLetter(rune(str[i]))) != true {

			i++
		}

		for i < j && (unicode.IsDigit(rune(str[j])) || unicode.IsLetter(rune(str[j]))) != true {

			j--
		}
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
