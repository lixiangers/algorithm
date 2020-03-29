package highlevel

import "fmt"

const number = 3

func SolveSudoku(originData [number][number]int) {
	var resultData [number][number]int
	if calSoduku(0, 0, originData, &resultData) {
		printSudoKuResult(resultData)
	} else {
		fmt.Println("error")
	}
}

func calSoduku(rowIndex int, columnIndex int, originData [number][number]int, result *[number][number]int) bool {

	if rowIndex == (number-1) && columnIndex == number {
		return true
	}

	if columnIndex == number {
		//换一行
		rowIndex++
		columnIndex = 0
	}

	if rowIndex >= number {
		//超过最后一行
		return false
	}

	if originData[rowIndex][columnIndex] != 0 {
		result[rowIndex][columnIndex] = originData[rowIndex][columnIndex]
		if isValid(result, rowIndex, columnIndex) {
			isOk := calSoduku(rowIndex, columnIndex+1, originData, result)
			if isOk {
				return true
			}
		}
	} else {
		for k := 1; k < (number + 1); k++ {
			result[rowIndex][columnIndex] = k
			if isValid(result, rowIndex, columnIndex) {
				isOk := calSoduku(rowIndex, columnIndex+1, originData, result)
				if isOk {
					return true
				}
			}
		}
		return false
	}
	return false
}

func printSudoKuResult(ints [number][number]int) {
	for i := 0; i < number; i++ {
		for j := 0; j < number; j++ {
			fmt.Print(ints[i][j], " ")
		}
		fmt.Println()
	}
}

func isValid(sudokuData *[number][number]int, row int, column int) bool {
	var rowData []int
	var columnData []int
	//var gongGe9Data []int

	for i := 0; i <= column; i++ {
		rowData = append(rowData, sudokuData[row][i])
	}
	for i := 0; i <= row; i++ {
		columnData = append(columnData, sudokuData[i][column])
	}

	//1.判断当前行是否有重复数据
	if haveDuplicateElement(rowData) {
		return false
	}

	//2.判断当前列是否有重复数据
	if haveDuplicateElement(columnData) {
		return false
	}

	//3.判断当前9宫格内是否有重复的数据
	//计算当前9宫格的左上坐标
	//leftRow := row / 3 * 3
	//topColumn := column / 3 * 3
	//for i := leftRow; i <= row; i++ {
	//	for j := topColumn; j <= column; j++ {
	//		if sudokuData[leftRow][topColumn] != 0 {
	//			gongGe9Data = append(gongGe9Data, sudokuData[i][j])
	//		}
	//	}
	//}
	//
	//if haveDuplicateElement(gongGe9Data) {
	//	return false
	//}

	return true
}

func haveDuplicateElement(data []int) bool {
	tempMap := map[int]struct{}{}
	for _, element := range data {
		tempMap[element] = struct{}{}
	}

	return len(tempMap) < len(data)
}
