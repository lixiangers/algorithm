package highlevel

import "fmt"

var total = 0

func Get8QueenResult() {
	result := make([]int, 8, 8)
	cal8Queen(0, result)
}

func cal8Queen(row int, result []int) {
	if row == 8 {
		printQueens(result)
		total++
		fmt.Printf("第%d种\n", total)
	}
	//每一行都有8种放法.都回去尝试放
	for column := 0; column < 8; column++ {
		if isOk(row, column, result) {
			//第row行的第column占位
			result[row] = column
			cal8Queen(row+1, result)
		}
	}
}
func isOk(row int, column int, result []int) bool {
	left := column - 1
	right := column + 1
	for i := row - 1; i >= 0; i-- {
		//判断第i行的colmum上有棋子
		if result[i] == column {
			return false
		}

		if left >= 0 {
			//判断第i上左对角线上是否有棋子
			if result[i] == left {
				return false
			}
		}

		if right < 8 {
			//判断第i上右对角线上是否有棋子
			if result[i] == right {
				return false
			}
		}
		//每次向上一行，对角线位置需要改变
		left--
		right++
	}

	return true
}

func printQueens(result []int) {
	for row := 0; row < 8; row++ {
		for column := 0; column < 8; column++ {
			if result[row] == column {
				fmt.Print("Q ")
			} else {
				fmt.Print("* ")
			}
		}
		fmt.Println()
	}

	fmt.Println("===============")
}
