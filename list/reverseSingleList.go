package list

// 反转单链表
func ReverseSingleList(l *SingleList) {
	if nil == l.root {
		return
	}

	root := l.root
	var pre *SingleListNode
	var after *SingleListNode
	for nil != root {
		after = root.next
		root.next = pre

		pre = root
		root = after
	}

	// root最后是nil,不能指向root,而是pre
	l.root = pre
}
