package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	//array := []int{3, 1, 4,5}
	array := []int{3, 1, 4, 5, 5, 2, 6, 8}
	MergeSort(array)
	fmt.Println(array)
}
