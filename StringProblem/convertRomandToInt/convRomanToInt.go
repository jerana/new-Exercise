package main

import "fmt"

func main() {
	s := "LIX"
	fmt.Println(romanToInt(s))
}

/*aA
Convert Input string which is represened by ROmand Char, and return integer
corresponding to symbol
Algo : travers string from right to left and check if right char is less than its immediate left char , then
substract its value from left char values
*/
var romanKey = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	l := len(s) - 1
	sum := romanKey[rune(s[l])] //Trace right most char
	for i := len(s) - 2; i >= 0; i-- {
		if s[i] < s[i+1] {
			sum -= romanKey[rune(s[i])]
		} else {
			sum += romanKey[rune(s[i])]
		}
	}
	return sum

}
