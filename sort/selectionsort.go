package sort

// 选择排序
// 思路:也是把数组分成有序区域和无序区域。每次从无序区域选择最小的放到有序区的尾部。这里有序开始是一个都没有的从0开始
// 外循环是循环无序区域i: 0 to len .一开始都是无序的
// 内循环是从当前无序区域的开始到数组的最后j:  i to len
func SelectionSort(array []int) {
	if nil == array {
		return
	}

	for i := 0; i < len(array); i++ {
		minIndex := i
		min := array[i]

		// 循环找出当前无序区域的最小值的index
		for j := i; j < len(array); j++ {
			if array[j] < min {
				min = array[j]
				minIndex = j
			}
		}

		//把这次循环的最小值放在有序区域的尾部
		array[i], array[minIndex] = array[minIndex], array[i]
	}
}
