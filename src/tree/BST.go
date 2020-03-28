package tree

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

type binaryTree struct {
	root *node
}

func (tree *binaryTree) insert(value int) {

	if tree.root == nil {
		tree.root = &node{data: value}
	} else {
		tree.root.insert(value)
	}

}

func (n *node) insert(value int) {

	if value < n.data {
		if n.left == nil {
			n.left = &node{data: value}

		} else {
			n.left.insert(value)
		}

	} else {
		if n.right == nil {
			n.right = &node{data: value}

		} else {
			n.right.insert(value)
		}

	}
}

func inorder(root *node) {

	if root != nil {
		inorder(root.left)
		fmt.Println(root.data)
		inorder(root.right)
	}

}

func preorder(root *node) {

	if root != nil {
		fmt.Println(root.data)
		preorder(root.left)
		preorder(root.right)
	}

}

func postorder(root *node) {

	if root != nil {
		postorder(root.left)
		postorder(root.right)
		fmt.Println(root.data)
	}

}

//BST ...
func BST() {

	tree := &binaryTree{}

	tree.insert(12)
	tree.insert(10)
	tree.insert(34)
	tree.insert(24)
	tree.insert(59)
	tree.insert(70)

	// tree1 := &binaryTree{}

	// tree1.insert(12)
	// tree1.insert(10)
	// tree1.insert(34)
	// tree1.insert(24)
	// tree1.insert(59)
	// tree1.insert(71)

	//tree.root.left.left = &node{data: 9}

	// fmt.Println("Inorder  :")
	// inorder(tree.root)
	// fmt.Println("Preorder  :")
	// preorder(tree.root)
	// fmt.Println("Postorder  :")
	// postorder(tree.root)

	//leftView(tree.root)
	//rightView(tree.root)
	//CheckBST1(tree.root)
	//CheckBST2(tree.root)
	//CheckLCA(tree.root)
	//fmt.Println(checkSymmetry(tree.root, tree1.root))

	//fmt.Println(isMirrorView(tree.root, tree.root))
	//height(tree.root)
	inorder(tree.root)
	CheckElement(tree.root)
	inorder(tree.root)
}
