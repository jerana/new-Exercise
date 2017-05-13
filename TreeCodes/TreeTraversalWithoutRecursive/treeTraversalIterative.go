package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

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
	return inorder
}

/*Preorder Traversal*/
/*
1.Visit Node : push into Stack
2.Pop Nodes : add Right Subtree Node then left tree Nodes
*/
func preOrderTraversalIterative(root *TreeNode) []int {
	var preorder []int
	var stack []*TreeNode
	if root == nil {
		return nil
	}
	stack = append(stack, root)
	for len(stack) != 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1] //pop Nodes
		preorder = append(preorder, p.Val)
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}
	return preorder
}

/*Post Order Traversal*/
/*Do same logic as preOder expect instead change nodes push order i.e root, right,left
after that reverse the list
*/
func postOrderTraversalIterative(root *TreeNode) []int {
	var postorder []int
	var stack []*TreeNode
	if root == nil {
		return nil
	}
	stack = append(stack, root)
	for len(stack) != 0 {
		p := stack[len(stack)-1]     //Last In nodes
		stack = stack[:len(stack)-1] //pop nodes
		postorder = append(postorder, p.Val)
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
		if p.Right != nil {
			stack = append(stack, p.Right)
		}
	}
	reverseList(postorder)
	return postorder
}
func reverseList(list []int) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}
