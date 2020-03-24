package list

import "testing"

func TestIsHuiWenSingList(t *testing.T) {
	singleList := New()
	//singleList.PushBack("l")
	//singleList.PushBack("e")
	//singleList.PushBack("v")
	//singleList.PushBack("e")
	//singleList.PushBack("l")

	//singleList.PushBack("n")
	//singleList.PushBack("o")
	//singleList.PushBack("o")
	//singleList.PushBack("o")
	//singleList.PushBack("o")
	//singleList.PushBack("o")
	//singleList.PushBack("o")
	//singleList.PushBack("n")

	singleList.PushBack("1")
	singleList.PushBack("2")
	singleList.PushBack("3")
	singleList.PushBack("4")
	singleList.PushBack("5")
	singleList.PushBack("6")

	IsHuiWenSingList(singleList)
}
