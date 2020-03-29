package heap

import (
	"bytes"
	"fmt"
	"strconv"
)

type Heap struct {
	capacity int
	count    int
	array    []int
}

func NewHeap(capacity int) *Heap {
	//这里实现一个大顶堆
	heap := Heap{}
	heap.capacity = capacity
	heap.count = 0
	//从下表1开始,capacity要+1
	heap.array = make([]int, capacity+1, capacity+1)
	return &heap
}

func BuildHeap(array []int, count int) *Heap {
	heap := NewHeap(count)
	//把数组转换成从1开始
	for i := 0; i < count; i++ {
		heap.array[i+1] = array[i]
	}

	//从非叶子节点开始(叶子节点向下堆化只能和自己比较没有意义)，从上向下堆化。直到根节点(包含)
	//堆是完全二叉树，所以count / 2之后的节点都是叶子节点
	for i := count / 2; i >= 1; i-- {
		heapify(heap.array, count, i)
	}
	return heap
}

func (heap *Heap) Insert(item int) error {
	if heap.count >= heap.capacity {
		return fmt.Errorf("heap is full")
	}
	heap.count++
	//1. 放到最后
	heap.array[heap.count] = item
	i := heap.count
	//从下向上堆化
	//i/2:父节点 直到小于等于父节点,或者到了根节点
	for i/2 > 0 && heap.array[i] > heap.array[i/2] {
		Swap(heap.array, i, i/2)
		i = i / 2
	}
	return nil
}

func (heap *Heap) ToString() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("heap item :")
	for _, num := range heap.array {
		buffer.WriteString(strconv.Itoa(num))
	}

	return buffer.String()
}

func (heap *Heap) RemoveRoot() error {
	if heap.count <= 0 {
		return fmt.Errorf("heap is empty")
	}

	//1.把最后一个节点赋值个根节点
	heap.array[1] = heap.array[heap.count]
	heap.array[heap.count] = 0
	heap.count--
	heapify(heap.array, heap.count, 1)
	return nil
}

/**
  从上向下堆化
*/
func heapify(slice []int, count int, startIndex int) {
	for {
		maxIndex := startIndex
		//节点和自己的左右节点比较,找出最大的
		if startIndex*2 <= count && slice[maxIndex] < slice[startIndex*2] {
			maxIndex = startIndex * 2
		}

		if startIndex*2+1 <= count && slice[maxIndex] < slice[startIndex*2+1] {
			maxIndex = startIndex*2 + 1
		}

		if maxIndex == startIndex {
			//发现节点已经比自己左右节点大了，堆化完成
			break
		}
		//交互节点数据
		Swap(slice, startIndex, maxIndex)
		//继续向下堆化
		startIndex = maxIndex
	}
}

func Swap(slice []int, index1 int, index2 int) {
	temp := slice[index1]
	slice[index1] = slice[index2]
	slice[index2] = temp
}

func Sort(slice []int, count int) []int {
	//1. 建堆
	heap := BuildHeap(slice, count)
	for k := count; k > 1; {
		//2. 每次都把堆顶元素移动到最后
		Swap(heap.array, 1, k)
		k--
		//3. 堆化剩余的元素，直到剩下一个元素
		heapify(heap.array, k, 1)
	}

	return heap.array
}
