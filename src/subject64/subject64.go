package subject64

import (
	"math"
)

var (
	m int
	n int
	nums [][]int
	cache [][]int
)

func minPathSum(grid [][]int) int {
	//初始化全局变量，否则会出错，leetcode的bug
	m = 0
	n = 0
	nums = [][]int{}
	cache = [][]int{}

	m = len(grid)
	n = len(grid[0])
	nums = grid

	for i := 0; i < m; i++ {
		dp := make([]int,n)
		for j := 0; j < n; j++ {
			dp[j] = 0
		}
		cache = append(cache,dp)
	}//初始化缓存数组的所有元素为0

	return dfs(0,0)
}

//返回[i,j]点到终点的最短路径
func dfs(i,j int) int {
	//到达边界
	if i == m || j == n {
		return math.MaxInt32
	}
	//到达终点
	if i == m-1 && j == n-1 {
		return nums[i][j]
	}
	//尝试读取缓存
	if cache[i][j] != 0 {
		return cache[i][j]
	}
	right := dfs(i,j+1)
	down := dfs(i+1,j)
	nextMinPath := 0
	//取右边和下边的较小一个
	if right <= down {
		nextMinPath = right
	} else {
		nextMinPath = down
	}
	//当前[i,j]点的最短路径是下一最短路径和当前路径和
	nowMinPath := nextMinPath + nums[i][j]
	//存入缓存
	cache[i][j] = nowMinPath
	return nowMinPath
}