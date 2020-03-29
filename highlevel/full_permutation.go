package highlevel

import (
	"fmt"
	"strings"
)

var permutationArray []int
var selectedArray []bool
var permutationNumber = 0

func GetFullPermutation(originData []int) {
	permutationArray = make([]int, len(originData), len(originData))
	selectedArray = make([]bool, len(originData), len(originData))

	assemble(0, originData)
}

func assemble(index int, originData []int) {
	if index == len(originData) {
		permutationNumber++
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
