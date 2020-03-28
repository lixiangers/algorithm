package sort

// 冒泡排序。
// 思路:每次外循环都找出最大的一个值。每次内循环就是相邻的两个比较，较大的值放在后面.
// 外循环次数 i:0 to len
// 每次内循环次数 j:0 to len-1-i  -1是因为两两相邻比较到倒数第二个就能比较完。 -i 是因为每次外循环完成后数组最后i个就是有序的且是比较大的值
// j从0是因为前面不是有序的
func BubbleSort(array []int) {
	if nil == array {
		return
	}

	for i := 0; i < len(array); i++ {
		// 优化:如果发现某次外循环过程中一次交换都没有说明已经是有序的
		hasChange := false
		for j := 0; j < len(array)-1-i; j++ {
			// 相邻两两比较，大的放到后面去
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				hasChange = true
			}
		}

		if !hasChange {
			break
		}
	}
}
