package main

/*
ProlemStatment:
Given an array of strings, group anagrams together.

Example:

Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
Output:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
Note:

All inputs will be in lowercase.
The order of your output does not matter.
*/
func groupAnagrams(strs []string) [][]string {
	hMap := make(map[string][]string, 0)
	for _, str := range strs {
		v := hashVal(str)
		hMap[v] = append(hMap[v], str)
	}
	result := [][]string{}
	for _, res := range hMap {
		result = append(result, res)
	}
	return result
}
func hashVal(str string) string {
	val := [26]rune{}
	for _, c := range str {
		val[c-'a']++
	}
	s := []rune{}
	for i := 0; i < 26; i++ {
		s = append(s, rune(val[i]))
	}

	return string(s)
}
