package tree

import "fmt"

func height(root *node) {

	max := 0

	level := 1

	findHeight(root, &max, level)

	fmt.Println("height")
	fmt.Println(max)
}

func findHeight(root *node, max *int, level int) {

	if root == nil {
		return
	}

	if level > *max {
		*max = level
	}

	findHeight(root.left, max, level+1)
	findHeight(root.right, max, level+1)

}
