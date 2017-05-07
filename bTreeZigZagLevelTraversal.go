/*
Given a binary tree, return the zigzag level order traversal of its nodes' values. (ie, from left to right, then right to left for the next level and alternate between).

For example:
Given binary tree [3,9,20,null,null,15,7],
*/

package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

/*Algo:  Use 2 Stacks , on level traversal
Step1: Pop Elem from first Stack and Push it into second stack
Step2 : Do same as above steps till both Stack get empty

*/
//defining Stack using Slice
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

//custom Stack implementation Using Slice

type tstack []*TreeNode

func (s tstack) Push(n *TreeNode) tstack {
	return append(s, n)
}
func (s tstack) Pop() tstack {
	if s.isEmpty() {
		return nil
	}
	l := len(s)
	return (s[:l-1])
}
func (s tstack) Top() *TreeNode {
	l := len(s)
	return s[l-1]
}
func (s tstack) isEmpty() bool {
	if len(s) == 0 {
		return true
	}
	return false
}

var llist [][]int

func bTreeZigZagTraversal(root *TreeNode) [][]int {
	var s1, s2 tstack
	if root == nil {
		return nil
	}
	llist = make([][]int, 0) //Build 2 d array
	level := -1
	//push root into Stack 1
	s1 = s1.Push(root)
	var n *TreeNode
	for !s1.isEmpty() || !s2.isEmpty() {
		level++
		if !s1.isEmpty() {
			//Add Row on per level
			llist = append(llist, make([]int, 0))
		}
		for !s1.isEmpty() { //Pop all elements from s1 and push into s2
			n = s1.Top()
			s1 = s1.Pop()
			//level order list
			l := llist[level]
			l = append(l, n.val)
			if n.right != nil {
				s2 = s2.Push(n.right)
			}
			if n.left != nil {
				s2 = s2.Push(n.left)
			}
		}
		level++
		if !s2.isEmpty() {
			//Add Row on per level
			llist = append(llist, make([]int, 0))
		}
		for !s2.isEmpty() {
			n = s2.Top()
			s2 = s2.Pop()
			//level order list
			l := llist[level]
			l = append(l, n.val)
			if n.right != nil {
				s1 = s1.Push(n.right)
			}
			if n.left != nil {
				s1 = s1.Push(n.left)
			}
		}

	}
	fmt.Println(llist)
	return llist
}
