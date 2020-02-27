package array

import "fmt"

// 这是只是演练，完全没有必要.slice就是一个动态的数组类型
type ArrayList struct {
	count int
	array []string
}

func NewArrayList() *ArrayList {
	return &ArrayList{count: 0, array: make([]string, 0)}
}

// Appends the specified element to the end of this list
func (list *ArrayList) Add(item string) {
	list.array = append(list.array, item)
	list.count += 1
}

// Removes the first occurrence of the specified element from this list, if it is present
func (list *ArrayList) Remote(item string) {
	startIndex := -1
	for index, value := range list.array {
		if item == value {
			startIndex = index
			list.count -= 1
			break
		}
	}

	if startIndex >= 0 {
		for i := startIndex; i < list.count; i++ {
			list.array[startIndex] = list.array[startIndex+1]
		}

		list.array[list.count] = ""
	}
}

// Returns the element at the specified position in this list
func (list *ArrayList) Get(index int) string {
	result := ""
	if index < 0 || index >= list.count {
		fmt.Println("invalid index ", index)
	} else {
		result = list.array[index]
	}
	return result
}

// Returns true if this list Contains the specified element.
func (list *ArrayList) Contains(item string) bool {
	result := false
	for _, value := range list.array {
		if item == value {
			result = true
			break
		}
	}
	return result
}
