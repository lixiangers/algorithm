package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	array := []int{2, 3, 1, 4, 5, 5, 2, 6, 8}
	QuickSort(array)
	fmt.Println(array)
}
