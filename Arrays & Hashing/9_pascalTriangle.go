package main

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		row := make([]int, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				row[j] = 1
			} else {
				val := res[i-1][j-1] + res[i-1][j]
				row[j] = val
			}
		}
		res[i] = append(res[i], row...)
	}
	return res
}
