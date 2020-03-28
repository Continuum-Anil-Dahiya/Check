package tree

import (
	"fmt"
	"math"
)

//CheckElement ...
func CheckElement(root *node) {

	max := math.MinInt64

	checkElementUtil(root, &max)
}

//checkElementUtil
func checkElementUtil(root *node, max *int) {

	if root == nil {
		return
	}

	if *max < root.data {
		*max = root.data
	}

	fmt.Println("*max")
	fmt.Println(*max)

	checkElementUtil(root.right, max)
	root.data = *max + root.data
	checkElementUtil(root.left, max)

}
