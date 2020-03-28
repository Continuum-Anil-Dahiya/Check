package tree

import "fmt"

func rightView(root *node) {

	maxLevel := 0
	level := 1
	rightViewUtil(root, &maxLevel, level)
}

func rightViewUtil(root *node, maxLevel *int, level int) {

	if root == nil {
		return
	}

	if level > *maxLevel {
		fmt.Println(root.data)
		*maxLevel = level
	}

	rightViewUtil(root.right, maxLevel, level+1)
	rightViewUtil(root.left, maxLevel, level+1)
}
