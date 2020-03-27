package tree

import (
	"fmt"
	"testing"
)

func TestPreOrder(t *testing.T) {
	fmt.Println("preOrder")
	binaryTree := assembleBinaryTree()

	binaryTree.preOrder()
}

func TestInOrder(t *testing.T) {
	fmt.Println("InOrder")
	binaryTree := assembleBinaryTree()

	binaryTree.inOrder()
}

func TestPostOrder(t *testing.T) {
	fmt.Println("PostOrder")
	binaryTree := assembleBinaryTree()

	binaryTree.postOrder()
}

func assembleBinaryTree() BinaryTree {
	node6 := TreeNode{value: 6}
	node7 := TreeNode{value: 7}
	node3 := TreeNode{value: 3, left: &node6, right: &node7}

	node4 := TreeNode{value: 4}
	node5 := TreeNode{value: 5}
	node2 := TreeNode{value: 2, left: &node4, right: &node5}

	node1 := TreeNode{value: 1, left: &node2, right: &node3}

	binaryTree := BinaryTree{root: &node1}
	return binaryTree
}
