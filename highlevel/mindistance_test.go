package highlevel

import (
	"fmt"
	"testing"
)

func TestMinDistance(t *testing.T) {
	matrix := make([][]int, 4)
	matrix[0] = []int{1, 3, 5, 9}
	matrix[1] = []int{2, 1, 3, 4}
	matrix[2] = []int{5, 2, 6, 7}
	matrix[3] = []int{6, 8, 4, 3}

	fmt.Println(MinDistance(matrix, 4))
}

func TestMinDistance2(t *testing.T) {
	matrix := make([][]int, 4)
	matrix[0] = []int{1, 3, 5, 9}
	matrix[1] = []int{2, 1, 3, 4}
	matrix[2] = []int{5, 2, 6, 7}
	matrix[3] = []int{6, 8, 4, 3}

	fmt.Println(MinDistance2(matrix, 4))
}
