package main

import "fmt"

func main() {
	num := "222333"
	list := mnemohicsNumber(num)
	fmt.Println(list)
}

/*
Write a Program which takes as input a phone number , specified as a string of digits
and returns all possible character sequences that corresponding to the phone number
the cell phone keypad is specified by a mapping that takes a digit and returns the
corresponding set of characters.
*/

var keyMap = map[int]string{
	0: "0",
	1: "1",
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

func mnemohicsNumber(number string) []string {
	var list []string
	partial_number := ""
	helper(number, 0, partial_number, &list)
	return list
}
func helper(num string, idx int, partial_num string, list *[]string) {
	if idx == len(num) {
		//reach at end of string
		*list = append(*list, partial_num)
	} else {
		c := num[idx] - '0'   //Get interger value
		tmp := keyMap[int(c)] //Mapped string
		for _, c := range tmp {
			partial_str := partial_num + string(c)
			helper(num, idx+1, partial_str, list)
		}

	}
}
