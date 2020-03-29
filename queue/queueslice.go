package queue

// slice 实现队列
type Queue2 struct {
	array []interface{}
}

func NewQueue2() Queue2 {
	return Queue2{}
}

func (q *Queue2) Enqueue(v interface{}) {
	q.array = append(q.array, v)
}

func (q *Queue2) Dequeue() interface{} {
	if len(q.array) <= 0 {
		return nil
	}

	element := q.array[0]
	q.array = q.array[1:len(q.array)]
	return element
}

func (q *Queue2) len() int {
	return len(q.array)
}
