package queue

import (
	"fmt"
)

// 实现循环队列
type CircularQueue struct {
	cap   int // 容量
	len   int //当前长度
	head  int //头指针
	tail  int //尾指针
	array []interface{}
}

func NewCircularQueue(capacity int) CircularQueue {
	return CircularQueue{
		cap:   capacity,
		len:   0,
		head:  0,
		tail:  0,
		array: make([]interface{}, capacity)}
}

func (c *CircularQueue) Enqueue(v interface{}) {
	if c.cap == c.len {
		// 队列满了
		fmt.Println("queue is full")
		return
	}

	c.array[c.tail] = v
	if c.tail == c.cap-1 {
		// 跳动头部
		c.tail = 0
	} else {
		c.tail++
	}

	c.len++
}

func (c *CircularQueue) Dequeue() interface{} {
	if c.len == 0 {
		//　队列是空
		fmt.Println("queue is empty")
		return nil
	}

	item := c.array[c.head]
	c.array[c.head] = 0 //清除为0,也可以不做
	if c.head == c.cap-1 {
		// 跳到头部
		c.head = 0
	} else {
		c.head++
	}
	c.len--
	return item
}

func (c *CircularQueue) Len(v interface{}) int {
	return c.len
}
