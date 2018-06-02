package main

/*
Problem Statment :
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

Input:
[
 [ 1, 2, 3 ],
  [ 4, 5, 6 ],
   [ 7, 8, 9 ]
   ]
   Output: [1,2,3,6,9,8,7,4,5]
   Example 2:

   Input:
   [
     [1, 2, 3, 4],
       [5, 6, 7, 8],
         [9,10,11,12]
	 ]
	 Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	ans := []int{}
	r1 := 0
	c1 := 0
	r2 := len(matrix) - 1
	c2 := len(matrix[0]) - 1

	for r1 <= r2 && c1 <= c2 {
		for c := c1; c <= c2; c++ {
			ans = append(ans, matrix[r1][c])
		}
		for r := r1 + 1; r <= r2; r++ {
			ans = append(ans, matrix[r][c2])
		}
		if r1 < r2 && c1 < c2 {
			for c := c2 - 1; c > c1; c-- {
				ans = append(ans, matrix[r2][c])
			}
			for r := r2; r > r1; r-- {
				ans = append(ans, matrix[r][c1])
			}
		}
		r1++
		c1++
		r2--
		c2--
	}
	return ans
}
