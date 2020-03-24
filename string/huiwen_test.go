package string

import (
	"fmt"
	"testing"
)

func TestIsHuiWen(t *testing.T) {
	fmt.Println("result:", IsHuiWen("level"))
	fmt.Println("result:", IsHuiWen("nooovooon"))
	fmt.Println("result:", IsHuiWen("noooooon"))
	fmt.Println("result:", IsHuiWen(""))
	fmt.Println("result:", IsHuiWen("levels"))
}
