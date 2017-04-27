package main

import "math"

/*
Given a binary tree, return the tilt of the whole tree.

The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the sum of all right subtree node values. Null node has tilt 0.

The tilt of the whole tree is defined as the sum of all nodes' tilt.

Example:
Input:
         1
	        /   \
		      2     3
		      Output: 1
		      Explanation:
		      Tilt of node 2 : 0
		      Tilt of node 3 : 0
		      Tilt of node 1 : |2-3| = 1
		      Tilt of binary tree : 0 + 0 + 1 = 1
		      Note:

		      The sum of node values in any subtree won't exceed the range of 32-bit integer.
		      All the tilt values won't exceed the range of 32-bit integer.

*/
type tree struct {
	val   uint32
	left  *tree
	right *tree
}

func findGivenBtreeTilt(root *tree) uint32 {
	var tilt uint32
	if root != nil {
		travers(root, &tilt)
	}
	return tilt
}
func travers(root *tree, tilt *uint32) uint32 {
	if root == nil {
		return 0
	}
	lsum := travers(root.left, tilt)  //Check left Subtree
	rsum := travers(root.right, tilt) //Check right Subtree
	//Node tilt = ABS(left subtree sum -right sumtree sum)
	ntilt := uint32(math.Abs(float64(lsum - rsum)))
	*tilt += ntilt
	return lsum + root.val + rsum // return Subtree SUm Value to caller
}
