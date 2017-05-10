package main

import "fmt"

func findPath(root, n *TreeNode, nList []*TreeNode) bool {
	if root == nil {
		return false
	}
	nList = append(nList, root)
	if root.Val == n.Val {
		return true
	}
	if findPath(root.Left, n, nList) || findPath(root.Right, n, nList) {
		return true
	}
	nList = nList[len(nList)-1]
	return false
}

func findLCA(root, n1, n2 *TreeNode) *TreeNode {
	var nList1, nList2 []*TreeNode
	if !findPath(root, nList1, n1) || !findPath(root, nList2, n2) {
		return nil
	}
	for i := 0; i < len(nList1) && i < len(nList2); i++ {
		if nList1[i] == nList2[i] {
			break
		}
	}
	return nList1[i-1]
}
