package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var k, n int
	rand.Seed(time.Now().Unix())
	arr := []int{5, 7, 3, 11}
	fmt.Println(arr)
	sample := randomSampling(arr, 3)
	fmt.Println("Sample set:", sample)
	fmt.Println("Let's start Sampling:")
	fmt.Println("Sampling sz:")
	fmt.Scan(&k)
	fmt.Println("first Sample from :")
	fmt.Scan(&n)

	readRandomSampleData(k, n)
}

/*
Implementing an Algo that takes an input an array of distinct elements and a size
and returns a subset of the given size of the array elements . All subsets should be
equally likely . Return the result in input array itself.
*/
func randomSampling(arr []int, k int) []int {
	var rSubset []int
	for i := 0; i < k; i++ {
		r := random(i, len(arr)) //Random number gentor in range
		rSubset = append(rSubset, arr[r])
		//Swap the random selected sumber i position num
		arr[i], arr[r] = arr[r], arr[i]
	}
	return rSubset
}

//Generate Random number in range
func random(min, max int) int {
	return rand.Intn(max-min) + min
}

/*
Idea is from packet capture tools, which collect packet at random with given size ,
Problem: Design algo which takes size k as a input and return sample data from continously dyanic
stream
Approch 1: Save all data in disk and Call Above function to get a random subset from size n data
This approch will cause huge space problem

Aprroch 2: Dynamic calculation of Random subset, first get size k subset from first n sequence data
then findout the propibilty of n+1 could fall into given subset or not by using random sequence generator
if given n+1 probility fall in given sequence then place it by replacing this read into given sequence
*/

func readRandomSampleData(k, n int) {
	var x int
	var count int
	var data []int
	var sample []int
	for {
		r, err := fmt.Scan(&x)
		if err != nil {
			fmt.Println("Stop input sequence")
			return
		}
		if r < 1 {
			fmt.Println("Invalid input")
		}
		fmt.Println("Read Data:", x)
		//first read N  sequence data
		if count < n {
			data = append(data, r)
			count++
			continue
		} else if count == n {
			sample = randomSampling(data, k)
			fmt.Println(sample)
		} else { //Case n+1 sequence
			r := rand.Intn(n)
			if r < k {
				//Include into Sample and replcase read sequence
				sample[r] = x
				fmt.Println(sample)
			}

		}
		count++

	}
}
