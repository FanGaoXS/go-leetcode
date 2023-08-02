package subject240

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}

	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	i := len(matrix) - 1
	j := 0
	for i >= 0 && j < len(matrix[0]) {
		if target == matrix[i][j] {
			return true
		}
		if target > matrix[i][j] {
			j++
			continue
		}
		if target < matrix[i][j] {
			i--
			continue
		}
	}

	return false
}
