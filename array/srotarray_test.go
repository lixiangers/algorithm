package array

import (
	"fmt"
	"testing"
)

func TestNewSortArray(t *testing.T) {
	defer func() {
		if err := recover(); err == nil { //没有产生panic异常
			fmt.Println(err)
			t.Fatal("should panic")
		}
	}()
	array := &SortArray{}
	array.Add(1)
}

func TestNewSortArray2(t *testing.T) {
	array := NewSortArray(10)
	if array.index != 0 && array.capacity != 10 {
		t.Error("init error")
	}
}

func TestSortArray_Add(t *testing.T) {
	array := NewSortArray(5)
	array.Add(3)
	array.Add(2)
	array.Add(2)
	array.Add(5)
	array.Add(4)
	array.Add(7)
	if array.String() != "2,2,3,4,5,end" {
		t.Error("add error")
	}
	fmt.Println(array.String())
}

func TestSortArray_Remove(t *testing.T) {
	array := NewSortArray(5)
	array.Add(3)
	array.Add(2)
	array.Add(1)
	array.Add(5)
	array.Add(4)

	array.Remove(3)
	if array.String() != "1,2,4,5,0,end" {
		fmt.Println(array.String())
		t.Error("remove error")
	}
}
