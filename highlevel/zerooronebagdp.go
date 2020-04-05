package highlevel

// 0/1背包问题．有几个误判他们的重量分别是[]int{w1,w2,w3,w4,w5}.我的背包能装的总重量是w.需要我们计算出在不超过背包总重量的前提下
// 我们最大可以装入的重量.
// 题解:对于每个物品来说，都有两种选择，装进背包或者不装进背包。动态规划的思路是每一步都计算好当前的状态．下一步依赖上一步的状态，动态
// 推导出这一步的状态，一直到最后．　结合这个例子．针对每一个物品，我们可以放和不放的情况，然后结合上一步中的重量，推导出这一步的所有
// 可能的重量．不包括于背包总重量．这样的话到最后一个物品的时候我们就能知道所有的重量组合情况．然后选择最接近背包总重量的．
// 动态规划解决0/1问题
//@items 物品各自的重量
//@items 物品的个数
//@maxWeight 背包支持的最大重量
func KnapsackDP(items []int, n int, maxWeight int) int {
	// 保存每一阶段重量的可能性.下一步根据上一步和组合计算这一步的组合可能.背包总重量是w,那么不超过w的情况有0-w 即w+1种情况
	states := make([]bool, maxWeight+1)

	// 第一阶段的需要初始化
	states[0] = true //第一个物品不装入背包
	if items[0] <= maxWeight {
		// 如果第一个物品的重量不超过总重量，第一个物品重量的列也是一种状态
		states[items[0]] = true
	}

	// 从第二物品开始到最后一个物品，依赖上一步的状态，推到出这一步的状态
	for i := 1; i < n; i++ {
		// 从后往前推到出每个状态的可能性
		for j := maxWeight; j >= 0; j-- {
			if states[j] == true && j+items[i] <= maxWeight {
				// 在原来可能性上再加上当前物品重量.相当于是放入物品．对于不放入物品的情况，其实就是加上0.也就是保存上一步的状态即可．
				// 这里我们只需要在上一步是true的列中加上改物品重量，把相应列置为true.需要对超过weight的情况做出判断
				states[j+items[i]] = true
			}
		}
	}

	// 完成推导后，找到最解决maxWeight的情况
	for i := maxWeight; i >= 0; i++ {
		if states[i] == true {
			return i
		}
	}

	return 0
}

// 在0/1背包问题的前提下，我们引入物品对于的价值value.求在不超过背包重量的前提下，价值最大是多少
// 把整个求解过程分为 n 个阶段，每个阶段会决策一个物品是否放到背包中。每个阶段决策完之后，背包中的物品的总重量以及总价值.
// 我们使用state[n][weight+1]来保存每层的不同状态．不过这里我们存储的是对于的最大总价值.
// 每层放入的情况还是分weight+1种，但是它对应的值是最大的值.
//@items 物品各自的重量
//@values 物品各自的价值
//@items 物品的个数
//@maxWeight 背包支持的最大重量
func KnapsackDP2(items []int, values []int, n int, maxWeight int) int {
	states := make([][]int, n)

	// 初始化状态
	for i := 0; i < n; i++ {
		intslice := make([]int, maxWeight+1)
		for j := 0; j < maxWeight+1; j++ {
			intslice[j] = -1
		}

		states[i] = intslice
	}

	// 初始化第一个物品
	states[0][0] = 0 //第一个物品不放入
	if items[0] <= maxWeight {
		//第一个物品放入
		states[0][items[0]] = values[0]
	}

	// 依次推导每层的状态
	for i := 1; i < n; i++ {
		for j := 0; j < maxWeight+1; j++ {
			// 不放入物品
			if states[i-1][j] >= 0 {
				states[i][j] = states[i-1][j]
			}
		}

		for j := 0; j < maxWeight+1; j++ {
			// 放入物品.需要判断不要超过maxWeight
			if states[i-1][j] >= 0 && j+items[i] <= maxWeight {
				currValue := states[i-1][j] + values[i]
				if currValue > states[i][j+items[i]] {
					// 保存最大值
					states[i][j+items[i]] = currValue
				}
			}
		}
	}

	//找到最大值
	maxValue := 0
	for i := 0; i <= maxWeight; i++ {
		if states[n-1][i] > maxValue {
			maxValue = states[n-1][i]
		}
	}

	return maxValue
}
