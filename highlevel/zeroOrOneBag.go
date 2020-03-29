package highlevel

import "fmt"

var MaxWeight = 0
var zeroOrOneBagNumber = 0

//
func GetZeroOrOneBagResult(itemList []int, totalWeight int) {
	//assembleItemSimple(0, 0, itemList, totalWeight)
	assembleItemOptimize(0, 0, itemList, totalWeight)
}

//全部可能都装入包中，找出所有值中满足条件的最大值
func assembleItemSimple(index int, currentWeight int, itemList []int, totalWeight int) {
	if index == len(itemList) {
		if currentWeight > MaxWeight && currentWeight <= totalWeight {
			MaxWeight = currentWeight
		}
		zeroOrOneBagNumber++
		fmt.Println("weight:", currentWeight)
		fmt.Println("number:", zeroOrOneBagNumber)
		return
	}

	//不装入当前Item
	assembleItemSimple(index+1, currentWeight, itemList, totalWeight)
	//装入当前Item
	assembleItemSimple(index+1, currentWeight+itemList[index], itemList, totalWeight)
}

//优化:当前中间超过包的重量的就中止当前选择，回溯到前面
//
func assembleItemOptimize(index int, currentWeight int, itemList []int, totalWeight int) {
	if index == len(itemList) {
		if currentWeight > MaxWeight {
			MaxWeight = currentWeight
		}
		zeroOrOneBagNumber++
		fmt.Println("weight:", currentWeight)
		fmt.Println("number:", zeroOrOneBagNumber)
		return
	}

	//不装入当前Item
	assembleItemOptimize(index+1, currentWeight, itemList, totalWeight)
	//如果加上当前item的重量没有超过总重量，就加入包中 (剪枝操作)
	if currentWeight+itemList[index] <= totalWeight {
		assembleItemOptimize(index+1, currentWeight+itemList[index], itemList, totalWeight)
	}
}
