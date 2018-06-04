package main

/*
Problem Statment:
Given an input string , reverse the string word by word.

Example:

Input:  ["t","h","e"," ","s","k","y"," ","i","s"," ","b","l","u","e"]
Output: ["b","l","u","e"," ","i","s"," ","s","k","y"," ","t","h","e"]
Note:

A word is defined as a sequence of non-space characters.
The input string does not contain leading or trailing spaces.
The words are always separated by a single space.
Follow up: Could you do it in-place without allocating extra space?
*/
func reverseWords(str []byte) {
	if len(str) != 0 {
		reverse(str, 0, len(str)-1) //Reverse first whole string
		i, j := 0, 0
		for j < len(str)-1 {
			if str[j] == ' ' {
				reverse(str, i, j-1)
				j++
				i = j
			} else {
				j++
			}
		}
		reverse(str, i, j)
	}
}
func reverse(str []byte, start, end int) {
	for start <= end {
		str[start], str[end] = str[end], str[start]
		start++
		end--
	}
}
