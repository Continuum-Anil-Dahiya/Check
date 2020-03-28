package tree

// Mirror view of Binary tree is symmertical or not
// https://www.youtube.com/watch?v=9jH2L2Ysxko

func isMirrorView(root1 *node, root2 *node) bool {

	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.data != root2.data {
		return false
	}

	return isMirrorView(root1.left, root2.right) && isMirrorView(root1.right, root2.left)

}
