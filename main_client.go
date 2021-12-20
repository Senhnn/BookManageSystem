package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0 || len(matrix[len(matrix)-1]) == 0 {
		return false
	}
	hig := len(matrix) - 1
	wid := len(matrix[0]) - 1
	i := hig
	j := 0

	for true {
		if target > matrix[i][j] {
			j++
		} else if target < matrix[i][j] {
			i--
		} else if target == matrix[i][j] {
			return true
		}
		if j > wid || i < 0 {
			return false
		}
	}
	return false
}
