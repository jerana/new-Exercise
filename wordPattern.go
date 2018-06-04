package main

/*
Problem Statment:
Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

Example 1:

Input: pattern = "abba", str = "dog cat cat dog"
Output: true
Example 2:

Input:pattern = "abba", str = "dog cat cat fish"
Output: false
Example 3:

Input: pattern = "aaaa", str = "dog cat cat dog"
Output: false
Example 4:

Input: pattern = "abba", str = "dog dog dog dog"
Output: false
Notes:
You may assume pattern contains only lowercase letters, and str contains lowercase letters separated by a single space.
*/
import "strings"

func wordPattern(pattern string, str string) bool {
	pMap := make(map[rune]string, 0)
	sMap := make(map[string]rune, 0)
	strSlice := strings.Split(str, " ")

	if len(pattern) != len(strSlice) {
		return false
	}
	for i, p := range pattern {
		if _, ok := pMap[p]; ok && pMap[p] != strSlice[i] {
			return false
		}
		if _, ok := sMap[strSlice[i]]; ok && sMap[strSlice[i]] != p {
			return false
		}
		pMap[p] = strSlice[i]
		sMap[strSlice[i]] = p
	}
	return true
}
