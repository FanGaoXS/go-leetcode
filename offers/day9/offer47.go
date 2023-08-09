package main

var (
	w = 0
	h = 0

	nums  [][]int
	cache [][]int
)

func maxValue(grid [][]int) int {
	w = len(grid[0])
	h = len(grid)

	nums = grid
	cache = make([][]int, 0, h)
	for i := 0; i < h; i++ {
		cache = append(cache, make([]int, w))
	}

	return dfs(0, 0)
}

// [i,j]点到终点的最大路径和
func dfs(i, j int) int {
	if i == h || j == w {
		return 0
	}
	if i == h-1 && j == w-1 {
		return nums[i][j]
	}
	if v := cache[i][j]; v != 0 {
		return v
	}

	right := dfs(i+1, j)
	down := dfs(i, j+1)
	cache[i][j] = max(right, down) + nums[i][j]

	return cache[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{
		{1, 2},
		{1, 1},
	}
	res := maxValue(grid)
	println(res)
}
