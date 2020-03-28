package tree

import "fmt"

//Longest Common Ansestor for BST
//Longest Common Ansestor for binary Tree

//CheckLCA ...
func CheckLCA(root *node) {
	fmt.Println(lcaofBST(root, 24, 70).data)
	fmt.Println(lcaofBinaryTree(root, 10, 70).data)
}

func lcaofBST(root *node, x int, y int) *node {

	if root == nil {
		return nil
	}

	//
	// if root.data == x || root.data == y {
	// 	return root
	// }

	if root.data > x && root.data > y {
		return lcaofBST(root.left, x, y)
	}

	if root.data < x && root.data < y {
		return lcaofBST(root.right, x, y)

	}
	return root
}

func lcaofBinaryTree(root *node, x int, y int) *node {

	if root == nil {
		return nil
	}

	if root.data == x || root.data == y {
		return root
	}

	left := lcaofBST(root.left, x, y)
	right := lcaofBST(root.right, x, y)

	if left != nil && right != nil {
		return root

	}

	if left != nil {
		return left
	}

	return right

}
