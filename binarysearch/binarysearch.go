package binarysearch

// 标准的二分查找. array假设是有序的
func StandardBinarySearch(array []int, target int) int {
	result := -1

	if nil == array {
		return result
	}

	lowIndex := 0
	highIndex := len(array) - 1
	// 注意是 low<=high，而不是 low<high。
	for lowIndex <= highIndex {
		// 等同(highIndex-lowIndex)/2 位移更快
		middle := (highIndex - lowIndex) >> 1
		//这个地方一定要注意加上lowIndex
		middle += lowIndex
		if array[middle] == target {
			result = middle
			break
		} else if array[middle] > target {
			//说明在middle的左边
			highIndex = middle - 1
		} else {
			//说明在middle的右边
			lowIndex = middle + 1
		}
	}

	return result
}

// 假设数组是从小到大有序的。查找第一个等于给定的值
func BinarySearchForFirst(array []int, target int) int {
	result := -1
	if nil == array {
		return result
	}

	lowIndex := 0
	highIndex := len(array) - 1

	for lowIndex <= highIndex {
		middle := lowIndex + (highIndex-lowIndex)>>1
		if array[middle] > target {
			// 在middle的左边
			highIndex = middle - 1
		} else if array[middle] < target {
			//　在middle的右边
			lowIndex = middle + 1
		} else {
			//相等的情况
			// 已经是最左边了或者middle左边的值不是target。则当前这个就是第一个.否则继续二分
			if middle == 0 || array[middle-1] != target {
				result = middle
				break
			} else {
				highIndex = middle - 1
			}
		}
	}

	return result
}

// 假设从小到达排序。查找最后一个等于给target的index
func BinarySearchForLast(array []int, target int) int {
	result := 1
	if nil == array {
		return result
	}

	lowIndex := 0
	highIndex := len(array) - 1

	for lowIndex <= highIndex {
		middle := lowIndex + (highIndex-lowIndex)>>1
		if array[middle] > target {
			//说明在middle的左边
			highIndex = middle - 1
		} else if array[middle] < target {
			// 说明在middle的右边
			lowIndex = middle + 1
		} else {
			// 相同
			// 如果已经是数组的最后一个。或者middle右边的值不等于给定的值说明middle就是最后一个给定的值.
			// 否则继续查找
			if middle == len(array)-1 || array[middle+1] != target {
				result = middle
				break
			} else {
				lowIndex = middle + 1
			}
		}
	}
	return result
}
