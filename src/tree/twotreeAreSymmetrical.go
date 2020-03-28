package tree

func checkSymmetry(root1 *node, root2 *node) bool {

	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.data != root2.data {
		return false
	}

	return checkSymmetry(root1.left, root2.left) && checkSymmetry(root1.right, root2.right)

}
