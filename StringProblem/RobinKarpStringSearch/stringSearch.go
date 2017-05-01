package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "CGC"
	t := "GACCGCGCCACGC"
	result := findSeachStringIndexList(s, t)
	fmt.Println(result)

}

func findSeachStringIndexList(s, t string) []int {

	if len(s) > len(t) {
		return nil
	}
	Kbase := 26 //Base number for rolling hash
	power_s := 1
	//Search string hash and first target substring hash
	s_hash, t_hash := 0, 0
	i := 0

	for i = 0; i < len(s); i++ {
		if i > 0 {
			power_s *= Kbase
		}
		s_hash = s_hash*Kbase + int(s[i])
		t_hash = t_hash*Kbase + int(t[i])
	}
	var match_idx []int
	for i = len(s); i < len(t); i++ {
		if t_hash == s_hash && 0 == strings.Compare(s, t[(i-len(s)):i]) {
			match_idx = append(match_idx, (i - len(s)))
			fmt.Println("Match index :", i-len(s))
		}
		t_hash -= power_s * int(t[(i-len(s))]) //remove left most char from window
		t_hash = t_hash*Kbase + int(t[i])
	}
	//Case for last occurance
	if t_hash == s_hash && 0 == strings.Compare(s, t[i-len(s):]) {
		match_idx = append(match_idx, (i - len(s)))
	}
	return match_idx
}
