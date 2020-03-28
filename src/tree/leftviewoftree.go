package tree

import "fmt"

//Link
//https://www.geeksforgeeks.org/print-left-view-binary-tree/
//TimeComplixity:O(n) as we simply reaching to every node
//SpaceComplexity:O(n)
func leftView(root *node) {

	// Set intial maxlevel=0 and level=1
	maxLevel := 0
	level := 1
	// pass maxLevel as a pointer because we need to update it on eveylevel
	leftViewUtil(root, &maxLevel, level)
}

func leftViewUtil(root *node, maxLevel *int, level int) {

	// Base condition
	if root == nil {
		return
	}

	if level > *maxLevel {
		fmt.Println(root.data)
		*maxLevel = level
	}

	leftViewUtil(root.left, maxLevel, level+1)
	leftViewUtil(root.right, maxLevel, level+1)

}
