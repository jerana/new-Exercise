package main

/*Find LCA of Binary Tree*/
/*Algo 1 : Recursive*/

func findLCABTree(root, n1, n2 *TreeNode) *TreeNode {
	//If root is nil or one of node match with root
	if !root || root == n1 || root == n2 {
		return root
	}
	//Case 2 : Find if any of given nodes exist in left subtree and right subtree of given tree
	lSubtree := findLCABTree(root.Left, n1, n2)
	rSubtree := findLCABTree(root.Right, n1, n2)
	//If both are non-nil means given nodes fall into different tree
	if lSubtree && rSubtree {
		return root
	}
	//If both are nil means nodes doesn't belong to tree
	if !rSubtree && !lSubtree {
		return nil
	}
	if rSubtree != nil {
		return rSubtree
	}
	return lSubtree
}

/*Algo 2 : Path trace solution */
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
