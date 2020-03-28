package sort

import "fmt"

func MergeSort(array []int) {
	mergeSort(array, 0, len(array)-1)
}

// 冒泡排序:
// 思路: 递归执行.每次都从中间一分为二.直到只有一个元素。然后合并.

func mergeSort(array []int, start int, end int) {
	if start >= end {
		return
	}
	fmt.Println(fmt.Printf("start:%d end:%d\n", start, end))

	middle := (end - start) / 2
	middle += start
	mergeSort(array, start, middle)
	mergeSort(array, middle+1, end)
	merge(array, start, array[start:middle+1], array[middle+1:end+1])
}

func merge(array []int, start int, ints []int, ints2 []int) {
	temp := make([]int, 0)

	//需要两个指针x,y分别指向正两个已经排好序的数组指针
	x := 0
	y := 0

	for x < len(ints) && y < len(ints2) {
		if ints[x] < ints2[y] {
			temp = append(temp, ints[x])
			x++
		} else {
			temp = append(temp, ints2[y])
			y++
		}
	}

	// 判断哪个还有数据，合并到后面
	if x == len(ints) {
		// ints2 还有数据
		temp = append(temp, ints2[y:]...)
	} else {
		// ints 还有数据
		temp = append(temp, ints[x:]...)
	}

	// 拷贝到原始数组中
	for i := 0; i < len(temp); i++ {
		array[start+i] = temp[i]
	}

}

// 合并。这里我们需要新开一个[]int.存放合并后的数据。最后赋值给原理数据
//func merge(array []int, start int, middle int, end int) {
//	// 临时的temp，用于存放合并的数据
//	temp := make([]int, end-start+1)
//
//	//需要两个指针，分别指向前一个
//}
