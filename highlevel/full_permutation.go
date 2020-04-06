package highlevel

import (
	"fmt"
	"strings"
)

var permutationArray []int
var selectedArray []bool
var permutationNumber = 0
var result [][]int

func GetFullPermutation(originData []int)[][]int {
	permutationArray = make([]int, len(originData))
	selectedArray = make([]bool, len(originData))

	assemble(0, originData)
	fmt.Println(result)
	return result
}

func assemble(index int, originData []int) {
	if index == len(originData) {
		permutationNumber++
		temp:=make([]int,len(permutationArray))
		copy(temp,permutationArray)
		result=append(result,temp)
		fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(permutationArray), " ", ",", -1), "[]"))
		return
	}

	for i, data := range originData {
		if !selectedArray[i] {
			//如果没有选中则选中
			permutationArray[index] = data
			selectedArray[i] = true
			assemble(index+1, originData)
			//index排列后,设置非选中状态，自动循环到下一个
			selectedArray[i] = false
		}
	}

}
