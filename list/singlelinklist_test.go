package list

import (
	"fmt"
	"testing"
)

func TestSingleList(t *testing.T) {
	singleList := New()
	singleList.PushBack(1)
	singleList.PushBack(2)
	singleList.PushBack(3)
	singleList.PushBack(4)
	singleList.PushBack(5)
	singleList.PushBack(6)

	fmt.Println(singleList.len)

	singleList.RemoveFront()
	fmt.Println(singleList.len)
	singleList.RemoveBack()
	fmt.Println(singleList.len)
}
