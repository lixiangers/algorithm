package array

import (
	"fmt"
	"strings"
)

type SortArray struct {
	slice    []int
	index    int
	capacity int
}

func NewSortArray(capacity int) *SortArray {
	return &SortArray{slice: make([]int, capacity, capacity), index: 0, capacity: capacity}
}

func (array *SortArray) check() bool {
	if nil == array.slice {
		panic("sort slice must be init")
		return false
	}

	return true
}

// Insert in order
func (array *SortArray) Add(item int) {
	array.check()
	if array.index < array.capacity {
		currentIndex := -1
		for i := 0; i < array.index; i++ {
			if array.slice[i] >= item {
				currentIndex = i //找到位置
				break
			}
		}

		if currentIndex == -1 {
			array.slice[array.index] = item
		} else {
			for i := array.index; i >= currentIndex && i > 0; i-- {
				array.slice[i] = array.slice[i-1]
			}
			array.slice[currentIndex] = item
		}

		array.index++
	}
}

// Removes the first occurrence of the specified element from this array, if it is present
func (array *SortArray) Remove(item int) {
	array.check()
	operateIndex := -1
	for i := 0; i < array.index; i++ {
		if array.slice[i] == item {
			operateIndex = i
			break
		}
	}

	if operateIndex > -1 {
		for i := operateIndex; i < array.index-1; i++ {
			array.slice[i] = array.slice[i+1]
		}
		array.slice[array.index-1] = 0
		array.index--
	}
}

// Output by order
func (array *SortArray) String() string {
	array.check()
	builder := strings.Builder{}

	for _, value := range array.slice {
		builder.WriteString(fmt.Sprintf("%d,", value))
	}
	builder.WriteString("end")
	return builder.String()
}
