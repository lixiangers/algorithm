package array

import (
	"fmt"
	"testing"
)

func TestNewArrayList(t *testing.T) {
	list := NewArrayList()

	if 0 != list.count {
		t.Fatal("init array count is not 0")
	}
}

func TestArrayList_Add(t *testing.T) {
	list := NewArrayList()
	list.Add("lixiang")
	list.Add("yuhongbo")
	list.Add("lizhongyu")
	for index, value := range list.array {
		fmt.Printf("%d:%s\n", index, value)
	}
}

func TestArrayList_Remote(t *testing.T) {
	list := NewArrayList()
	list.Add("lixiang")
	list.Add("yuhongbo")
	list.Add("lizhongyu")

	list.Remote("yuhongbo")
	for index, value := range list.array {
		fmt.Printf("%d:%s\n", index, value)
	}
}

func TestArrayList_Get(t *testing.T) {
	list := NewArrayList()
	list.Add("lixiang")
	list.Add("yuhongbo")
	list.Add("lizhongyu")
	if list.Get(1) != "yuhongbo" {
		t.Fatal("error ")
	}
	if list.Get(3) != "" {
		t.Fatal("error ")
	}
}

func TestArrayList_Contains(t *testing.T) {
	list := NewArrayList()
	list.Add("lixiang")
	list.Add("yuhongbo")
	list.Add("lizhongyu")
	if !list.Contains("yuhongbo") {
		t.Fatal("error ")
	}
	if list.Contains("yuhongboxx") {
		t.Fatal("error ")
	}
}
