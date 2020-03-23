//合并两个有序的数组
package array

func MergeSortedArray(array1, array2 []int) []int {
	result := make([]int, 0)
	indexOfArray1 := 0
	indexOfArray2 := 0

	for indexOfArray1 < len(array1) && indexOfArray2 < len(array2) {
		if array1[indexOfArray1] < array2[indexOfArray2] {
			result = append(result, array1[indexOfArray1])
			indexOfArray1++
		} else {
			result = append(result, array2[indexOfArray2])
			indexOfArray2++
		}
	}

	if indexOfArray1 == len(array1) && indexOfArray2 < len(array2) {
		result = append(result, array2[indexOfArray2:]...)
	} else if indexOfArray2 == len(array2) && indexOfArray1 < len(array1) {
		result = append(result, array1[indexOfArray1:]...)
	}
	return result
}
