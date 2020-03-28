package sort

import (
	"fmt"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	array := []int{4, 5, 7, 3, 2, 1, 5, 2, 8, 7}
	InsertionSort(array)
	fmt.Println(array)
}
