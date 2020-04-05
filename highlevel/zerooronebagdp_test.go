package highlevel

import (
	"fmt"
	"testing"
)

func TestKnapsackDP(t *testing.T) {
	fmt.Println(KnapsackDP([]int{2, 2, 4, 6, 3}, 5, 9))
}

func TestKnapsackDP2(t *testing.T) {
	fmt.Println(KnapsackDP2([]int{2, 2, 4, 6, 3}, []int{3, 4, 8, 9, 6}, 5, 9))
}
