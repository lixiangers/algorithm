package highlevel

import (
	"fmt"
	"testing"
)

func TestKnapsackDP(t *testing.T) {
	fmt.Println(KnapsackDP([]int{2, 2, 4, 6, 3}, 5, 9))
}
