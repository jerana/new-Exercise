package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*
Problem Statment:
Numbers can be regarded as product of its factors. For example,

8 = 2 x 2 x 2;
  = 2 x 4.
Write a function that takes an integer n and return all possible combinations of its factors.

Note:

You may assume that n is always positive.
Factors should be greater than 1 and less than n.
Example 1:

Input: 1
Output: []
Example 2:

Input: 37
Output:[]
Example 3:

Input: 12
Output:
[
  [2, 6],
  [2, 2, 3],
  [3, 4]
]
Example 4:

Input: 32
Output:
[
  [2, 16],
  [2, 2, 8],
  [2, 2, 2, 4],
  [2, 2, 2, 2, 2],
  [2, 4, 4],
  [4, 8]
]

*/
//Solved using DFS algo
func getFactors(n int) [][]int {
	result := [][]int{}
	tmpList := []int{}
	helper(&result, tmpList, n, 2)
	return result
}
func helper(res *[][]int, list []int, n, start int) {
	if n <= 1 && len(list) > 1 {
		tmp := make([]int, len(list))
		copy(tmp, list)
		*res = append(*res, tmp)
		return
	}
	for i := start; i <= n; i++ {
		if n%i == 0 {
			list = append(list, i)
			helper(res, list, n/i, i)
			list = list[0 : len(list)-1]
		}
	}
}
