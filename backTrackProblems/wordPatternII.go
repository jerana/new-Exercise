package main

/*ProblemStatment:
Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty substring in str.

Example 1:

Input: pattern = "abab", str = "redblueredblue"
Output: true
Example 2:

Input: pattern = pattern = "aaaa", str = "asdasdasdasd"
Output: true
Example 3:

Input: pattern = "aabb", str = "xyzabcxzyabc"
Output: false
Notes:
You may assume both pattern and str contains only lowercase letters.
*/
func wordPatternMatch(pattern string, str string) bool {
	//hashMap for Patter tracker
	pMap := make(map[rune]string, 0)
	//hashMap for string trakcer
	sMap := make(map[string]bool, 0)

	return match(str, pattern, 0, 0, pMap, sMap)
}
func match(str, pattern string, pIndex, sIndex int, pMap map[rune]string, sMap map[string]bool) bool {
	//if Pattern search and string reach at end ; means mapping possible
	if pIndex == len(pattern) && sIndex == len(str) {
		return true
	}
	if pIndex == len(pattern) || sIndex == len(str) {
		return false
	}

	//Get Pattern character
	c := pattern[pIndex]

	if s, ok := pMap[rune(c)]; ok {
		//Check if this subString is exist which start from str[sIndex:]
		if !strings.HasPrefix(string(str[sIndex:]), s) {
			return false
		}
		//If Substring match ; go for ramining string
		return match(str, pattern, pIndex+1, sIndex+len(s), pMap, sMap)
	}
	//Case Pattern character does't match
	for i := sIndex; i < len(str); i++ {
		subStr := str[sIndex : i+1]
		//Reverse check for found Substring, if substring already has been mapped then continue
		//for next substring match ; this logic avoid having multiple pattern having same substring
		//mapped
		if _, ok := sMap[subStr]; ok {
			continue
		}
		pMap[rune(c)] = subStr
		sMap[subStr] = true
		if match(str, pattern, pIndex+1, i+1, pMap, sMap) {
			return true
		}
		delete(pMap, rune(c))
		delete(sMap, subStr)
	}
	return false
}
