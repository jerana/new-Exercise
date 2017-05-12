/*Algorithm :
Do Inorder traversal and detect first and second swap nodes

*/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var firstNode, secondNode, prevNode *TreeNode

func recoverBST(root *TreeNode) {
	prevNode = &TreeNode{Val: -(int(1<<30) - 1)}
	helperFunc(root)
	firstNode.Val, secondNode.Val = secondNode.Val, firstNode.Val
	return
}
func helperFunc(root) {
	if root == nil {
		return
	}
	helperFunc(root.Left)
	if firstNode == nil && prevNode.Val > root {
		firstNode = prevNode
	}
	if firstNode != nil && prevNode.Val > root {
		secondNode = root
	}
	helperFunc(root.Right)
}
