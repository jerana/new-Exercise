package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
 1. Iterative and push left Subtree Node till its left subtree is null
2. record  nodes
3. if nodes right subtree is not nil, then push nodes  right subtree nodes till its left subtree is null
*/
func inorderTraversalIterative(root *TreeNode) []int {
	var inorder []int
	var stack []*TreeNode
	if root == nil {
		return root
	}
	//First track first node
	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for len(stack) != 0 { //Till Stack is empty
		p := stack[len(stack)-1]     //Last Node Insert(i.e top node in Stack)
		stack = stack[:len(stack)-1] //Pop node
		inorder = append(inorder, p.Val)
		if p.Right != nil {
			l := p.Right //Right SUbtree
			for l != nil {
				stack = append(stack, l) //Push till reach left most node in this tree
				l = l.Left
			}
		}

	}
	return list
}
