package tree

import (
	"fmt"
	"math"
)

//https://www.geeksforgeeks.org/a-program-to-check-if-a-binary-tree-is-bst-or-not/
//Efficient Method Time Complexity:O(n) Space Complexity:O(1)

//CheckBST2 ...
func CheckBST2(root *node) {

	min := math.MinInt64
	max := math.MaxInt64

	if isbstUtil(root, min, max) {
		fmt.Println("Is BST")
	} else {
		fmt.Println("Not BST")
	}
}

func isbstUtil(root *node, min int, max int) bool {

	if root == nil {
		return true
	}

	if root.data < min || root.data > max {
		return false
	}

	if !isbstUtil(root.left, min, root.data-1) || !isbstUtil(root.right, root.data+1, max) {
		return false
	}

	return true
}
