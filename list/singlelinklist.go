//单向链表
package list

import (
	"fmt"
	"strings"
)

type SingleList struct {
	len  int
	root *SingleListNode
}

func New() *SingleList {
	return &SingleList{len: 0, root: nil}
}

func (l *SingleList) Len() int {
	return l.len
}
func (l *SingleList) PushBack(v interface{}) *SingleListNode {
	node := &SingleListNode{list: l, value: v, next: nil}
	if l.root == nil {
		l.root = node
	} else {
		//找出最后一个Node
		beforeNode := l.root
		for beforeNode.next != nil {
			beforeNode = beforeNode.next
		}
		beforeNode.next = node
	}
	l.len++
	return node
}

func (l *SingleList) PushFront(v interface{}) *SingleListNode {
	node := &SingleListNode{list: l, value: v, next: nil}
	if l.root == nil {
		l.root = node
	} else {
		//插入到第一个
		originRoot := l.root
		node.next = originRoot
		l.root = node
	}
	l.len++
	return node
}
func (l *SingleList) InsertAfter(node *SingleListNode, v interface{}) *SingleListNode {
	if node.list != l {
		fmt.Println("not the node of this list")
		return nil
	}

	nextNode := node.next
	newNode := &SingleListNode{list: l, value: v, next: nextNode}
	node.next = newNode

	l.len++
	return newNode
}

func (l *SingleList) Remove(v interface{}) bool {
	if l.root == nil {
		return false
	}

	node := l.root
	var beforeNode *SingleListNode = nil

	for node.value != v && node.next != nil {
		node = node.next
		beforeNode = node
	}

	if node.value == v && beforeNode != nil {
		newNextNode := node.next
		beforeNode.next = newNextNode
		l.len--
		return true
	}

	return false
}

func (l *SingleList) RemoveBack() interface{} {
	if l.root == nil {
		return nil
	}
	node := l.root
	var beforeNode *SingleListNode = nil

	for node.next != nil {
		beforeNode = node
		node = node.next
	}

	if beforeNode != nil {
		//删除最后一个节点前面的一个节点对最后一个节点的引用
		beforeNode.next = nil
	}

	l.len--
	return node.value
}

func (l *SingleList) RemoveFront() interface{} {
	if l.root == nil {
		return nil
	}

	rootValue := l.root.value
	nextNode := l.root.next
	l.root = nextNode

	l.len--
	return rootValue
}

func (l *SingleList) Front() interface{} {
	if l.root == nil {
		return nil
	}

	return l.root.value
}

func (l *SingleList) Back() interface{} {
	if l.root == nil {
		return nil
	} else {
		//找出最后一个Node
		beforeNode := l.root
		for beforeNode.next != nil {
			beforeNode = beforeNode.next
		}

		return beforeNode.value
	}
}

func (l *SingleList) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("len:%d\n", l.len))
	root := l.root
	for nil != root {
		builder.WriteString(fmt.Sprintf("%s\n", root.value))
		root = root.next
	}

	return builder.String()
}

type SingleListNode struct {
	next *SingleListNode

	value interface{}

	list *SingleList
}
