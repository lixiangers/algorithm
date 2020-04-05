package highlevel

import "math"

// 矩阵从0,0 到 n-1,n-1的最短路径
// 状态转义表法
// 走到 n-1,j-1的走法只有两种(n-1-1,j-1)，(n-1,j-1-1)走过来．
// distance(n-1,j-1)=matrix[i-1][j-1]+min(distance(n-1-1,j-1),distance(n-1,j-1-))
// 如果我们知道计算每一行每一列的对于的minDistance.最后直接取distance(n-1,j-1)的值就行
func MinDistance(matrix [][]int, n int) int {
	// 我们要推到每一行的状态．需要一些基础的数据.对于第一行和第一列的值我们是知道的．因为他们的走法只有一种．
	//就能直接算出．作为我们推到状态的的基础数据

	minDistanceState := make([][]int, n)
	for i, _ := range minDistanceState {
		minDistanceState[i] = make([]int, n)
	}

	// 初始化第一行和第一列的数据
	minDistanceState[0][0] = matrix[0][0]
	for i := 1; i < n; i++ {
		// 第一行
		minDistanceState[0][i] = matrix[0][i] + minDistanceState[0][i-1]
		// 第一列
		minDistanceState[i][0] = matrix[i][0] + minDistanceState[i-1][0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			minDistanceState[i][j] = matrix[i][j] + min(minDistanceState[i-1][j], minDistanceState[i][j-1])
		}
	}

	return minDistanceState[n-1][n-1]
}

// 状态转义方程法
// 状态转义方程: MinDistance(i,j)=matrix[i][j]+min(MinDistance(i-1,j),MinDistance(i,j-1))
func MinDistance2(matrix [][]int, n int) int {
	states := make([][]int, n)
	for i := 0; i < n; i++ {
		states[i] = make([]int, n)
	}
	return distance(matrix, states, n-1, n-1)
}

//@state 存放临时(i,j)的最小路径.这里可以避免重复的计算,优化算法
func distance(matrix [][]int, state [][]int, i, j int) int {
	if i == 0 && j == 0 {
		return matrix[0][0]
	}

	if state[i][j] > 0 {
		return state[i][j]
	}

	// 分别计算左边的最小值和右边的最小值
	minLeft, minUp := math.MaxInt64, math.MaxInt64
	if i-1 >= 0 {
		minLeft = distance(matrix, state, i-1, j)
	}

	if j-1 >= 0 {
		minUp = distance(matrix, state, i, j-1)
	}

	currMin := min(minLeft, minUp) + matrix[i][j]
	state[i][j] = currMin

	return currMin
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
