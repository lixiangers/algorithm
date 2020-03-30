package highlevel

import "math"

// 查找数组中的有序第几个元素
func FindTopMax(array []int, top int) int {
	if nil == array {
		return -1
	}

	return find(array, 0, len(array)-1, top)

}

func find(array []int, begin int, end int, top int) int {
	if begin > end {
		return -1
	}

	piovtIndex := partition(array, begin, end)
	if piovtIndex+1 == top {
		return array[piovtIndex]
	} else if piovtIndex+1 > top {
		// 在左边找
		return find(array, begin, piovtIndex-1, top)
	} else {
		// 在右边找
		return find(array, piovtIndex+1, end, top)
	}
}

func partition(array []int, begin int, end int) int {
	item := array[end]
	i := begin //有序区域指针

	for j := begin; j < end; j++ {
		if array[j] < item {
			array[j], array[i] = array[i], array[j]
			i++
		}
	}
	array[i], array[end] = array[end], array[i]
	return i
}

// 查找第二个元素.通过两个数字记录
func FindTop2(array []int) int {
	if nil == array || len(array) < 2 {
		return -1
	}

	min := math.MaxInt64
	secondMin := math.MaxInt64

	for _, value := range array {
		if value < min {
			secondMin = min
			min = value
		} else if value < secondMin {
			secondMin = value
		}
	}

	return secondMin
}
