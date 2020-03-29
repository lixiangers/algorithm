package heap

import (
	"bytes"
	"fmt"
	"strconv"
)

// 堆是个完全二叉树
// 节点要大于等于子树上所有节点(大顶堆)或者小于等于子树上所有节点(小顶堆)
// 使用数组存储是最方便的。对于index=i的于元素。其左节点是2*i+1 右节点是2*i+2 父节点是(i-1)/2
// 对于完全二叉树，非叶子节点是0-((count-1)/2-1)，后面的都是叶子节点．在建堆的时候我们只需要从非叶子节点开始堆化。
// 叶子节点没有节点可以比较，进行堆化没有意义

type Heap struct {
	capacity int
	count    int
	array    []int
}

//这里实现一个大顶堆
func NewHeap(capacity int) *Heap {
	heap := Heap{}
	heap.capacity = capacity
	heap.count = 0
	heap.array = make([]int, capacity)
	return &heap
}

func BuildHeap(array []int, count int) *Heap {
	heap := NewHeap(count)

	for i := 0; i < len(array); i++ {
		heap.array[i] = array[i]
	}

	heap.count = len(array)

	// 从非叶子节点开始(叶子节点向下堆化只能和自己比较没有意义)，从上向下堆化。直到根节点(包含)
	// 堆是完全二叉树，所以(count-1)/2 - 1之后的节点都是叶子节点
	for i := (count-1)/2 - 1; i >= 0; i-- {
		heapify(heap.array, i, count-1)
	}
	return heap
}

func (heap *Heap) Insert(item int) error {
	if heap.count >= heap.capacity {
		return fmt.Errorf("heap is full")
	}
	//1. 放到最后
	heap.array[heap.count] = item
	i := heap.count

	// 从下向上堆化
	// 父节点是(i-1)/2 直到小于等于父节点,或者到了根节点
	for (i-1)/2 >= 0 && heap.array[i] > heap.array[(i-1)/2] {
		// 和父节点交换
		parent := (i - 1) / 2
		swap(heap.array, i, parent)
		i = parent
	}
	heap.count++
	return nil
}

func (heap *Heap) String() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("heap item :")
	for _, num := range heap.array {
		buffer.WriteString(strconv.Itoa(num))
		buffer.WriteString(" ")
	}

	return buffer.String()
}

func (heap *Heap) RemoveRoot() error {
	if heap.count <= 0 {
		return fmt.Errorf("heap is empty")
	}

	//1.把最后一个节点赋值个根节点.最后一个节点赋值为0
	heap.array[0] = heap.array[heap.count-1]
	heap.array[heap.count-1] = 0
	heap.count--

	// 从上往下堆化
	heapify(heap.array, 0, heap.count-1)
	return nil
}

// 从上向下堆化
// @slice 堆底层对于的slice
// @startIndex 堆化的节点index
// @endIndex 堆化的结束节点index
func heapify(slice []int, startIndex int, endIndex int) {
	for {
		maxIndex := startIndex
		//节点和左节点比较,找出较大的
		leftIndex := startIndex*2 + 1
		if leftIndex <= endIndex && slice[maxIndex] < slice[leftIndex] {
			maxIndex = leftIndex
		}

		// 节点和右节点比较，找出较大的
		rightIndex := startIndex*2 + 2
		if rightIndex <= endIndex && slice[maxIndex] < slice[rightIndex] {
			maxIndex = rightIndex
		}

		if maxIndex == startIndex {
			// 建堆完成后，堆已经满足堆的特性。我们堆化只是一个节点发生变化。当遇到这个情况是可以直接break的。
			// 建堆过程中我们是从最后的节点开始堆化的，这样保证我们在堆化某一个节点时候，后面的已经满足了堆的特性
			// 发现节点已经比自己左右节点大了，也可能是遇到了叶子节点。堆化完成
			break
		}
		//交互节点数据
		swap(slice, startIndex, maxIndex)
		//继续向下堆化.从上往下
		startIndex = maxIndex
	}
}

func swap(slice []int, index1 int, index2 int) {
	slice[index1], slice[index2] = slice[index2], slice[index1]
}

func Sort(slice []int) []int {
	//1. 建堆
	heap := BuildHeap(slice, len(slice))
	for k := len(slice) - 1; k > 0; {
		//2. 每次都把堆顶元素（最大元素）移动到最后一位
		swap(heap.array, 0, k)
		k--
		//3. 堆化剩余的元素，直到剩下一个元素
		heapify(heap.array, 0, k)
	}

	return heap.array
}
