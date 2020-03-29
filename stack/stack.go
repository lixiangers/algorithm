// 链表实现stack
package stack

import list2 "container/list"

type Stack struct {
	list list2.List
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(v interface{}) {
	s.list.PushBack(v)
}

func (s *Stack) Pop() interface{} {
	element := s.list.Back()
	if nil != element {
		s.list.Remove(element)
		return element.Value
	}
	return nil
}
func (s *Stack) Top() interface{} {
	element := s.list.Back()
	if nil != element {
		return element.Value
	}
	return nil
}
func (s *Stack) Len() int {
	return s.list.Len()
}
