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

// 寻找第一个大于等于给定的值
func BinarySearchForFirstGreaterThanOrEqual(array []int, target int) int {
	result := -1
	if nil == array {
		return result
	}

	lowIndex := 0
	highIndex := len(array) - 1

	for lowIndex <= highIndex {
		middle := lowIndex + (highIndex-lowIndex)>>1

		if array[middle] >= target {
			// 发现一个大于等于target的值，需要判断是否已经是第一个符合的条件。
			// 如果已经是数组额左边边，或者middle-1小于target，就说明是的第一个符合的条件
			if middle == 0 || array[middle-1] < target {
				result = middle
				break
			} else {
				// 说明满足条件的值还在middle左边
				highIndex = middle - 1
			}
		} else {
			lowIndex = middle + 1
		}
	}

	return result
}

// 查找最后一个小于等于给定target的值
func BinarySearchForLastLessThanOrEqual(array []int, target int) int {
	result := -1
	if nil == array {
		return result
	}

	lowIndex := 0
	highIndex := len(array) - 1

	for lowIndex <= highIndex {
		middle := lowIndex + (highIndex-lowIndex)>>1

		if array[middle] <= target {
			// 发现一个小于等于target的值，需要判断是否已经是最后个符合的条件。
			// 如果已经是数组最右边，或者middle-1大于target，就说明是的最后一个符合的条件
			if middle == len(array)-1 || array[middle+1] > target {
				result = middle
				break
			} else {
				// 说明满足条件的值还在middle右边
				lowIndex = middle + 1
			}
		} else {
			highIndex = middle - 1
		}
	}

	return result
}

// 对于一个V行数组，找出最低端的值 [6,5,4,3,4,7,8,10]  3
// 左边大于它，右边也大于它的值.
func BinarySearchFindSplitPoint(array []int) int {
	result := -1
	if nil == array {
		return result
	}

	lowIndex := 0
	highIndex := len(array) - 1

	for lowIndex <= highIndex {
		middle := lowIndex + (highIndex-lowIndex)>>1
		if middle == 0 || middle == len(array)-1 {
			//当已经是左边或最右边的表示没有这样的值
			break
		} else if array[middle-1] > array[middle] && array[middle+1] > array[middle] {
			// 说明找到做了这个值
			result = middle
			break
		} else if array[middle+1] > array[middle] {
			// 说明是在右边的递增区域,想要的值在middle的左边
			highIndex = middle - 1
		} else {
			// 说明是在左边的递减区域,想要的值在middle的右边
			lowIndex = middle + 1
		}
	}

	return result
}
