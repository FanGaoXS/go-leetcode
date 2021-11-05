package main

func rotate(matrix [][]int)  {
	n := len(matrix)
	//先沿水平线交换
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-i][j]
			matrix[n-1-i][j] = temp
		}
	}
	//再沿主对角线交换
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
}
