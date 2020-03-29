package tree

import "fmt"

type TreeNode struct {
	data  int
	left  *TreeNode
	right *TreeNode
}

type BinaryTree struct {
	root *TreeNode
}

func newBinaryTree(value int) *BinaryTree {
	node := new(TreeNode)
	node.data = value
	node.left = nil
	node.right = nil
	return &BinaryTree{root: node}
}

func (t *BinaryTree) Insert(value int) error {
	if t.root == nil {
		return fmt.Errorf("root is nil")
	}

	currentNode := t.root
	for currentNode != nil {
		if value == currentNode.data {
			//如果找到相同的值的节点，插入失败
			return fmt.Errorf("value is exist")
		} else if value > currentNode.data {
			//如果大于比节点值
			if currentNode.right == nil {
				//如果节点的的右节点为空，则直接插入作为右节点；
				newNode := TreeNode{data: value}
				currentNode.right = &newNode
				return nil
			} else {
				//如果右节点不为空，以这个节点为根继续寻找
				currentNode = currentNode.right
			}

		} else if value < currentNode.data {
			//如果小于比节点值
			if currentNode.left == nil {
				//如果节点的的左节点为空，则直接插入作为左节点；
				newNode := TreeNode{data: value}
				currentNode.left = &newNode
				return nil
			} else {
				//如果左节点不为空，以这个节点为根继续寻找
				currentNode = currentNode.left
			}
		}
	}

	return nil
}

// 删除节点有３中情况
// 1. 是叶子节点,直接删除就行.父节点的左右指针根据这个节点是左右孩子情况置为nil
// 2. 删除的节点只有左孩子或者右孩子。直接把父节点的左右指针根据这个节点是左右孩子情况置的删除节点的孩子节点
// 3.　删除的节点左右孩子都有。需要找到右子树中的最小值。把要删除节点的值，设置为这个最小值。然后删除这个最小值节点即可
func (t *BinaryTree) Delete(value int) error {
	if t.root == nil {
		return fmt.Errorf("root is nil")
	}

	//　要删除的节点
	p := t.root
	// p的父节点s
	var pp *TreeNode = nil

	for nil != p && p.data != value {
		if p.data < value {
			// 先找到要删除的节点
			// 继续在右子树上查找
			pp = p
			p = p.right
		} else {
			// 继续在左子树上查找
			pp = p
			p = p.left
		}
	}

	// 没有找到这个节点就返回
	if nil == p {
		return fmt.Errorf("cat not find target")
	}

	// 判断p的左右孩子情况
	if nil != p.left && nil != p.right {
		//　如果左右孩子都有.需要找出右子树上的最小值
		minp := p.right
		minpp := p

		//　最小值肯定在左子树上
		for nil != minp.left {
			minpp = minp
			minp = minp.left
		}

		//　把右字数上最小值赋值给p,同时意味着这要删除的节点成了minp
		p.data = minp.data

		p = minp
		pp = minpp
	}

	// 走到这里，要删除的节点要么只有一个子节点要么没有子节点
	// 找出children 节点
	var children *TreeNode = nil
	if nil != p.left {
		children = p.left
	} else if nil != p.right {
		children = p.right
	}

	// 判断p是pp的什么孩子
	if nil == pp {
		// 说明是根节点
		t.root = children
	} else if p.left.data == value {
		// 删除的是左子树
		pp.left = children
	} else {
		// 删除的是右字数
		pp.right = children
	}
	return nil
}

func (t *BinaryTree) Search(target int) (*TreeNode, error) {
	if t == nil || t.root == nil {
		return nil, fmt.Errorf("tree is nill  or root node is nill")
	}

	p := t.root
	for nil != p {
		if p.data == target {
			return p, nil
		} else if p.data < target {
			p = p.right
		} else {
			p = p.left
		}
	}

	return nil, nil
}

func (t *BinaryTree) FindMax() *TreeNode {
	node := t.root

	for nil != node {
		node = node.right
	}

	return node
}

func (t *BinaryTree) FindMin() *TreeNode {
	node := t.root

	for nil != node {
		node = node.left
	}

	return node
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

	fmt.Println(node.data)
	preOrder(node.left)
	preOrder(node.right)
}

func inOrder(node *TreeNode) {
	if nil == node {
		return
	}

	inOrder(node.left)
	fmt.Println(node.data)
	inOrder(node.right)
}

func postOrder(node *TreeNode) {
	if nil == node {
		return
	}

	postOrder(node.left)
	postOrder(node.right)
	fmt.Println(node.data)
}
