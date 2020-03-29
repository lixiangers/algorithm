package queue

import (
	"fmt"
	"testing"
)

func TestCircularQueue(t *testing.T) {
	queue := NewCircularQueue(5)

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Dequeue()
	queue.Enqueue(7)
	queue.Dequeue()
	queue.Dequeue()
	queue.Dequeue()
	queue.Enqueue(8)
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	fmt.Println(queue)
}
