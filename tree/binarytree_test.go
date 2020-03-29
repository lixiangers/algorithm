package tree

import (
	"fmt"
	"testing"
)

func TestPreOrder(t *testing.T) {
	fmt.Println("PreOrderTraversal")
	binaryTree := assembleBinaryTree()

	binaryTree.PreOrderTraversal()

	binaryTree.PreOrderNoRecursion()
}

func TestInOrder(t *testing.T) {
	fmt.Println("InOrder")
	binaryTree := assembleBinaryTree()

	binaryTree.InOrderTraversal()
	binaryTree.InOrderTraversal()
}

func TestPostOrder(t *testing.T) {
	fmt.Println("PostOrder")
	binaryTree := assembleBinaryTree()

	binaryTree.PostOrderTraversal()
	binaryTree.PostOrderNoRecursion()
}

func TestBinaryTree(t *testing.T) {
	bst := assembleBinaryTree()
	fmt.Println(bst)
}

func TestBinaryTree_Invert(t *testing.T) {
	bst := assembleBinaryTree()

	bst.Invert()
}

func assembleBinaryTree() *BinaryTree {
	binaryTree := NewBinaryTree(5)
	binaryTree.Insert(4)
	binaryTree.Insert(2)
	binaryTree.Insert(3)
	binaryTree.Insert(1)
	binaryTree.Insert(8)
	binaryTree.Insert(7)
	binaryTree.Insert(9)
	return binaryTree
}
