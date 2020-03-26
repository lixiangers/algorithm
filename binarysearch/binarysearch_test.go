package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	array := []int{1, 2, 3, 4, 7, 8, 10, 12, 15}

	fmt.Println(StandardBinarySearch(array, 12))
}
