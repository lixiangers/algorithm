package list

// 反转单链表
func ReverseSingleList(l *SingleList) {
	if nil == l.root {
		return
	}

	node := l.root
	var pre *SingleListNode
	var after *SingleListNode
	for nil != node {
		after = node.next
		node.next = pre

		pre = node
		node = after
	}

	// node最后是nil,不能指向root,而是pre.指向原来的最后一个
	l.root = pre
}
