package tree

import "fmt"

type BinaryTree struct {
	root *TreeNode
}

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	value interface{}
}

// 前序打印.先先打印改节点，然后分别打印左右节点
func (tree *BinaryTree) preOrder() {
	preOrder(tree.root)
}

// 中序打印.先先打印改节点的左节点，然后打印改节点,最后打印改节点的右节点
func (tree *BinaryTree) inOrder() {
	inOrder(tree.root)
}

// 后序打印.先打印改节点的左节点，然后打印改节点的右节点,最后打印改节点
func (tree *BinaryTree) postOrder() {
	postOrder(tree.root)
}

func preOrder(node *TreeNode) {
	if nil == node {
		return
	}

	fmt.Println(node.value)
	preOrder(node.left)
	preOrder(node.right)
}

func inOrder(node *TreeNode) {
	if nil == node {
		return
	}

	inOrder(node.left)
	fmt.Println(node.value)
	inOrder(node.right)
}

func postOrder(node *TreeNode) {
	if nil == node {
		return
	}

	postOrder(node.left)
	postOrder(node.right)
	fmt.Println(node.value)
}
