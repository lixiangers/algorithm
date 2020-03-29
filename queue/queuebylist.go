package queue

import list2 "container/list"

// list实现队列
type Queue struct {
	list list2.List
}

func NewQueue() Queue {
	return Queue{}
}

func (q *Queue) Enqueue(v interface{}) {
	q.list.PushBack(v)
}

func (q *Queue) Dequeue() interface{} {
	element := q.list.Front()
	if nil == element {
		return nil
	}

	q.list.Remove(element)
	return element.Value
}

func (q *Queue) Len() int {
	return q.list.Len()
}
