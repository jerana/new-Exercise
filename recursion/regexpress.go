package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("Input :", text)

		fmt.Println(matches("m?e?t", text))
	}
}

func matches(regex string, test string) bool {
	if len(regex) == 0 {
		return len(test) == 0
	}
	return matchHelper(regex, "", test, 0)
}
func matchHelper(pattern, nStr, text string, sIndex int) bool {
	if sIndex == len(pattern) {
		return nStr == text
	}
	if pattern[sIndex] != '?' {
		m := nStr + string(pattern[sIndex])
		return matchHelper(pattern, m, text, sIndex+1)
	}
	return matchHelper(pattern, nStr[0:len(nStr)-1], text, sIndex+1) ||
		matchHelper(pattern, nStr, text, sIndex+1)
}
