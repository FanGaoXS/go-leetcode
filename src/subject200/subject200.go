package main

import "fmt"

var (
	m int
	n int
)

func numIslands(grid [][]byte) int {
	m = len(grid)
	n = len(grid[0])
	var result int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j]!='0' {
				result++
				dfs(grid,i,j)
			}
		}
	}
	return result
}

func dfs(grid [][]byte,i,j int)  {
	if i<0 || i>=m || j<0 || j>=n || grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(grid,i+1,j)
	dfs(grid,i-1,j)
	dfs(grid,i,j+1)
	dfs(grid,i,j-1)
}

func main() {
	grid := [][]byte{
		{'1','1','1','1','0'},
		{'1','1','0','1','0'},
		{'1','1','0','0','0'},
		{'0','0','0','0','0'},
	}
	num := numIslands(grid)
	fmt.Println("num =",num)
}