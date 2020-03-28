package sort

import (
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	array := []int{4, 3, 2, 1, 5, 2, 8, 7}
	SelectionSort(array)
	fmt.Println(array)
}
