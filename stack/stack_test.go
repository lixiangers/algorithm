package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(7)

	fmt.Println(stack.Len())
	fmt.Println(stack.Top())
	fmt.Println(stack.Len())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Len())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Len())
}
