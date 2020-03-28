package sort

import "fmt"

// 选择一个pivot,让小于pivot的数字移到pivot左边，大于等于pivot的移动它的后边
// 我们通过游标 i 把 A[p…r-1]分成两部分。A[p…i-1]的元素都是小于 pivot 的，
// 我们暂且叫它“已处理区间”，A[i…r-1]是“未处理区间”。
// 我们每次都从未处理的区间 A[i…r-1]中取一个元素 A[j]，与 pivot 对比，如果小于 pivot，则将其加入到已处理区间的尾部，也就是 A[i]的位置。
func QuickSort(array []int) {
	if nil == array {
		return
	}

	quickSort(array, 0, len(array)-1)
}

func quickSort(array []int, start int, end int) {
	if start >= end {
		return
	}

	partition := partition(array, start, end)

	fmt.Println("start:", start, " end:", end, "p:", partition)
	quickSort(array, start, partition-1)
	quickSort(array, partition+1, end)
}

// 1. 取end的值为pivot.
// 2. i 就是分割点.j是循环指针
// 2. 把[start,end] 分成以处理区[start ,i-1] ，未处理区[i+1,end]。i就是我们需要最终的值，分区值.
// 3. 从未处理依次取值和pivot比较。小于pivot的值就放在已处理区的尾部，已处理区域指针+1。否则不交换。 循环一直到end-1
func partition(array []int, begin int, end int) int {
	// i是已处理区域的指针.默认是start
	i := begin
	pivot := array[end]

	for j := begin; j < end; j++ {
		if array[j] < pivot {
			//发现比pivot小，需要把当前item移动到有序的尾部
			array[i], array[j] = array[j], array[i]
			//并且要后移已处理区域的指针
			i++
		}
	}

	//最后要把pivot移动到已处理区域指针位置
	array[i], array[end] = array[end], array[i]

	return i
}
