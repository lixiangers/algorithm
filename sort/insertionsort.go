package sort

import "fmt"

// 插入排序
// 思路:把数组前后分成有序的和无序的。默认第一个元素是有序的。每次从无序中选择一个，和有序部分从后往前依次比较找到位置。
//　外循环次数就是数组后面无序元素的次数i: 1 to len   1:是因为第一个元素认为是有序的
//  内循环次数就是和前面有序元素依次比较的次数j: i-1 to 0  i-1:是因为i是第一个无序元素的index i-1开始的元素才开始是有序的。
func InsertionSort(array []int) {
	if nil == array {
		return
	}

	// i就是第一个无序部分的index.i-1及以前是有序的
	for i := 1; i < len(array); i++ {
		item := array[i]
		j := i - 1
		for ; j >= 0; j-- {
			if array[j] > item {
				// 如果发现有比当前元素大的值.就把当前元素后移一位
				array[j+1] = array[j]
			} else {
				fmt.Println("before break:", j)
				// 已经找到位置
				break
			}
		}

		fmt.Println("done:", j)
		// j+1　就是要插入的位置
		array[j+1] = item
	}
}
