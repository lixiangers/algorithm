package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)

	fmt.Println(queue.Len())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())

	fmt.Println(queue.Len())
}
