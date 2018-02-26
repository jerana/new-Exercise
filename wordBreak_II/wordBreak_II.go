package main

import "fmt"

func main() {
	s := "catsanddog"
	dict := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak(s, dict))
}

//Problem:
//Given a non-empty string s and a dictionary wordDict containing a list of
//non-empty words, add spaces in s to construct a sentence where each word is a
//valid dictionary word. You may assume the dictionary does not contain
//duplicate words.

//Return all such possible sentences.

//For example, given
//s = "catsanddog",
//dict = ["cat", "cats", "and", "sand", "dog"].

//A solution is ["cats and dog", "cat sand dog"].
func wordBreak(s string, wordDict []string) []string {
	rMap := make(map[int][]string, 0)
	dMap := make(map[string]bool, 0)
	for _, w := range wordDict {
		dMap[w] = true
	}
	return wordBreakHelper(s, dMap, 0, rMap)
}
func wordBreakHelper(s string, wordDict map[string]bool, start int, rMap map[int][]string) []string {
	var rList []string
	if _, ok := rMap[start]; ok {
		return rMap[start]
	}
	if start == len(s) {
		rList = append(rList, "")
		return rList
	}
	for end := start + 1; end <= len(s); end++ {
		w := s[start:end]
		if _, ok := wordDict[w]; ok {
			list := wordBreakHelper(s, wordDict, end, rMap)
			for _, l := range list {
				if l == "" {
					rList = append(rList, w+l)
				} else {
					rList = append(rList, w+" "+l)
				}

			}
		}

	}
	rMap[start] = rList
	return rList
}
