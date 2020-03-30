package highlevel

import (
	"fmt"
	"testing"
)

func TestFindTopMax(t *testing.T) {
	array := []int{4, 6, 2, 7, 1}
	topMax := FindTopMax(array, 4)
	fmt.Println(topMax)
}

func TestFindTop2(t *testing.T) {
	array := []int{4, 6, 2, 7, 3, 1, 0}
	topMax := FindTop2(array)
	fmt.Println(topMax)
}
