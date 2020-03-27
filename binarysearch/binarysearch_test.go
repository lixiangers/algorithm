package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 7, 8, 10, 12, 15}

	fmt.Println(StandardBinarySearch(array, 12))

	array1 := []int{1, 2, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 6, 6, 7}
	fmt.Println(BinarySearchForFirst(array1, 4))
	fmt.Println(BinarySearchForLast(array1, 4))
	fmt.Println(BinarySearchForFirstGreaterThanOrEqual(array1, 5))
	fmt.Println(BinarySearchForFirstGreaterThanOrEqual(array1, 3))
	fmt.Println(BinarySearchForLastLessThanOrEqual(array1, 6))
	fmt.Println(BinarySearchForFirst(array1, 4))

	array2 := []int{6, 5, 4, 3, 2, 1, 2, 3, 4, 5, 7, 8}
	fmt.Println(BinarySearchFindSplitPoint(array1))
	fmt.Println(BinarySearchFindSplitPoint(array2))

}

func ExampleBinarySearchForFirstGreaterThanOrEqual() {

}
