package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*
Given an array of integers and an integer k, you need to find the total number of continuous subarrays whose sum equals to k.

Example 1:
Input:nums = [1,1,1], k = 2
Output: 2
Note:
The length of the array is in range [1, 20,000].
The range of numbers in the array is [-1000, 1000] and the range of the integer k is [-1e7, 1e7].

*/
//First Approach : O(n2)
func SumArray(arr []int, k int) {
	var count int
	for start := 0; start < len(arr); start++ {
		sum := 0
		for end := start + 1; end < len(arr); end++ {
			sum += arr[end]
			if sum == k {
				count++
			}

		}

	}

}

//Second Approcach :
func SumContinousArray(arr []int, k int) {
	for start := 0; start < len(arr); start++ {
		sum = 0
		for end := start + 1; end < len(arr); end++ {
			sum += arr[end]
			if sum == k {
				count++
			}
		}
	}
}

//Use HashMap  :O(n)
func SumArrayUsingHashmap(arr []int, k int) {
	//Calculate Cumulative Sum
	hashMap := make(map[int]int, 0)
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		e, ok := hashMap[sum-k]
		if ok {
			count += hashMap[sum-k]
		}
		e, ok := hashMap[sum]
		if ok {
			hashMap[sum]++
		}
		hashMap[sum]++
	}

}
