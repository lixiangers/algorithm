package binarysearch

// 标准的二分查找. array假设是有序的
func StandardBinarySearch(array []int, target int) int {
	result := -1

	if nil == array {
		return result
	}

	lowIndex := 0
	highIndex := len(array) - 1
	for lowIndex <= highIndex {
		middle := (highIndex - lowIndex) / 2
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
