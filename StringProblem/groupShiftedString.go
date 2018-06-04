package main

/*
Problem Statement:
Given a string, we can "shift" each of its letter to its successive letter, for example: "abc" -> "bcd". We can keep "shifting" which forms the sequence:

"abc" -> "bcd" -> ... -> "xyz"
Given a list of strings which contains only lowercase alphabets, group all strings that belong to the same shifting sequence.

Example:

Input: ["abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"],
Output:
[
  ["abc","bcd","xyz"],
  ["az","ba"],
  ["acef"],
  ["a","z"]
]
*/
//Use Shift value of each string and record them into hashMap
func groupStrings(strings []string) [][]string {
	hMap := make(map[byte][]string, 0)
	result := [][]string{}
	for _, s := range strings {
		v := shift(s)
		hMap[v] = append(hMap[v], s)
	}
	for _, w := range hMap {
		sort.Strings(w)
		result = append(result, w)
	}
	return result
}

//Function to cal string shift value
func shift(str string) byte {
	var t string
	for i := 1; i < len(str); i++ {
		diff := int(str[i]) - int(str[i-1])
		if diff < 0 {
			diff += 26
		}
		t += strconv.Itoa(diff) + ","
	}
	return t
}
