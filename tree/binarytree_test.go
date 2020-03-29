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

func TestBinaryTree(t *testing.T) {
	bst := assembleBinaryTree()
	fmt.Println(bst)
}

func assembleBinaryTree() *BinaryTree {
	binaryTree := newBinaryTree(5)
	binaryTree.Insert(4)
	binaryTree.Insert(2)
	binaryTree.Insert(3)
	binaryTree.Insert(1)
	binaryTree.Insert(8)
	binaryTree.Insert(7)
	binaryTree.Insert(9)
	return binaryTree
}
