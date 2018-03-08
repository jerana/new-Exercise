package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func upSideDownTree(root *tree) *tree {
	if root == nil {
		return t
	}
	if root.left == nil && root.right == nil {
		return root
	}
	newRoot := upSideDownTree(root.left)
	root.left.left = root.right
	root.left.right = root
	root.left = nil
	root.right = nil
	return newRoot
}

func MaxProduct(arr []int) {
	if len(arr) == 0 {
		return 0
	}
	maxProduct := arr[0]
	minproduct := arr[0]
	maxres := maxProduct
	for i := 1; i < len(arr); i++ {
		if arr[i] >= 0 {
			maxProduct = max(maxProduct*arr[i], arr[i])
			minProduct = min(minproduct*arr[i], arr[i])
		} else {
			tmp := maxProduct
			maxProduct = max(minproduct*arr[i], arr[i])
			minProduct = min(tmp*arr[i], arr[i])
		}
		maxres = max(maxProduct, maxres)
	}
}
