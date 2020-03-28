package tree

import (
	"fmt"
	"math"
)

//CheckBST1 ...
func CheckBST1(root *node) {

	if isCheckBST(root) {
		fmt.Println("IS BST")
	} else {
		fmt.Println("IS NOT BST")
	}
}

func isCheckBST(root *node) bool {

	if root == nil {
		return true
	}
	if root.left != nil && max(root.left) > root.data {
		return false
	}

	if root.right != nil && min(root.right) < root.data {
		return false
	}

	if !isCheckBST(root.left) || !isCheckBST(root.right) {
		return false
	}

	return true
}

//////////////////////////////////////////////////////
//Find max and min value in a binay tree
func max(root *node) int {
	max := math.MinInt64
	maxUtil(root, &max)
	return max
}

func maxUtil(root *node, max *int) {

	if root == nil {
		return
	}

	if root.data > *max {
		*max = root.data
	}
	maxUtil(root.left, max)
	maxUtil(root.right, max)
}

func min(root *node) int {
	min := math.MaxInt64
	minUtil(root, &min)
	return min
}

func minUtil(root *node, min *int) {

	if root == nil {
		return
	}

	if root.data < *min {
		*min = root.data
	}
	minUtil(root.left, min)
	minUtil(root.right, min)
}
