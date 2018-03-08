package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}
func sumFromStringDigit(s string) int {
	if len(s) <= 0 {
		return 0
	}
	num := new([]byte)
	for x := range s {
		if x >= '0' && x <= '9' {
			new = append(new, x)
		} else {
			if len(new) > 0 {
				sum += strconv.Atoi(new)
				new = new([]byte)
			}
		}
	}
	x := s[len(s)-1]
	if x >= '0' && x <= '9' && len(num) > 1 {
		sum += strconv.Atoi(new)
	}
	return sum

}
