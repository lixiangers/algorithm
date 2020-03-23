package array

import "testing"

func TestMergeSortedArray(t *testing.T) {
	array1 := []int{2, 5, 6, 9}
	array2 := []int{1, 3, 5, 7}

	result := MergeSortedArray(array1, array2)
	t.Log(result)
}
