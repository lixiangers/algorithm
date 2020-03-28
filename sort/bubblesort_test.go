package sort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	array := []int{4, 3, 2, 1, 5, 2, 8, 7}
	BubbleSort(array)
	fmt.Println(array)
}
