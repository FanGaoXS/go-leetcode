package main

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof

func main() {}

// 暴力解法
func findNumberIn2DArray(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}

	return false
}

// 搜索算法：从左下角起
func findNumberIn2DArray2(matrix [][]int, target int) bool {
	// [i, j] is the left-bottom
	i := len(matrix) - 1
	j := 0
	for i >= 0 && j < len(matrix[0]) {
		current := matrix[i][j]
		if target > current {
			j++
			continue
		}
		if target < current {
			i--
			continue
		}
		if target == current {
			return true
		}
	}
	return false
}
