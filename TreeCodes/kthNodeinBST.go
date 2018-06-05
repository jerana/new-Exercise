package main

/*
Problem Statment:
Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

Note:
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.

Example 1:

Input: root = [3,1,4,null,2], k = 1
Output: 1
Example 2:

Input: root = [5,3,6,2,4,null,null,1], k = 3
Output: 3

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var result int
	kthSmallestHelper(root, &k, &result)
	return result
}
func kthSmallestHelper(root *TreeNode, k *int, res *int) {
	if root == nil {
		return
	}
	kthSmallestHelper(root.Left, k, res)
	*k -= 1
	if *k == 0 {
		*res = root.Val
		return
	}
	kthSmallestHelper(root.Right, k, res)
}
