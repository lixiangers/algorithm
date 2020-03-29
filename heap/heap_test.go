package heap

import (
	"fmt"
	"testing"
)

func TestBuildHeap(t *testing.T) {
	array := []int{7, 5, 19, 8, 4, 1, 20, 13, 16}
	heap := BuildHeap(array, 11)
	fmt.Println(heap)

	heap.Insert(21)
	fmt.Println(heap)

	heap.RemoveRoot()
	fmt.Println(heap)

	array1 := []int{7, 5, 19, 8, 4, 1, 20, 13, 16}
	fmt.Println(Sort(array1))
}
