package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 7, 8, 10, 12, 15}

	fmt.Println(StandardBinarySearch(array, 12))

	array1 := []int{1, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4}
	fmt.Println(BinarySearchForFirst(array1, 4))
	fmt.Println(BinarySearchForLast(array1, 4))

}
