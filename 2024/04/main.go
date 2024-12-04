package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	inputLines, _ := utils.GetProblemLines()

	xmasMatrix := [][]rune{}

	for _, text := range inputLines {
		xmasMatrix = append(xmasMatrix, []rune(text))
	}

	sum := 0
	otherSum := 0

	for i := 0; i < len(xmasMatrix); i++ {
		for j := 0; j < len(xmasMatrix[0]); j++ {
			sum += GetCountXmas(xmasMatrix, i, j)
			otherSum += GetCountOtherXmas(xmasMatrix, i, j)
		}
	}

	fmt.Println("sum: ", sum)
	fmt.Println("otherSum: ", otherSum)
}

func GetCountOtherXmas(matrix [][]rune, i int, j int) int {
	sum := 0

	if matrix[i][j] == 'A' && j+1 < len(matrix[i]) && i+1 < len(matrix) && i-1 >= 0 && j-1 >= 0 {
		if matrix[i-1][j-1] == 'M' && matrix[i+1][j+1] == 'S' && matrix[i-1][j+1] == 'M' && matrix[i+1][j-1] == 'S' {
			sum++
		}

		if matrix[i-1][j-1] == 'S' && matrix[i+1][j+1] == 'M' && matrix[i-1][j+1] == 'M' && matrix[i+1][j-1] == 'S' {
			sum++
		}

		if matrix[i-1][j-1] == 'M' && matrix[i+1][j+1] == 'S' && matrix[i-1][j+1] == 'S' && matrix[i+1][j-1] == 'M' {
			sum++
		}

		if matrix[i-1][j-1] == 'S' && matrix[i+1][j+1] == 'M' && matrix[i-1][j+1] == 'S' && matrix[i+1][j-1] == 'M' {
			sum++
		}
	}

	return sum
}

func GetCountXmas(matrix [][]rune, i int, j int) int {
	sum := 0

	if matrix[i][j] == 'X' {
		if j+3 < len(matrix[i]) && matrix[i][j+1] == 'M' && matrix[i][j+2] == 'A' && matrix[i][j+3] == 'S' {
			sum++
		}

		if j-3 >= 0 && matrix[i][j-1] == 'M' && matrix[i][j-2] == 'A' && matrix[i][j-3] == 'S' {
			sum++
		}

		if i+3 < len(matrix) && matrix[i+1][j] == 'M' && matrix[i+2][j] == 'A' && matrix[i+3][j] == 'S' {
			sum++
		}

		if i-3 >= 0 && matrix[i-1][j] == 'M' && matrix[i-2][j] == 'A' && matrix[i-3][j] == 'S' {
			sum++
		}

		if j+3 < len(matrix[i]) && i+3 < len(matrix) && matrix[i+1][j+1] == 'M' && matrix[i+2][j+2] == 'A' && matrix[i+3][j+3] == 'S' {
			sum++
		}

		if j-3 >= 0 && i-3 >= 0 && matrix[i-1][j-1] == 'M' && matrix[i-2][j-2] == 'A' && matrix[i-3][j-3] == 'S' {
			sum++
		}

		if j+3 < len(matrix[i]) && i-3 >= 0 && matrix[i-1][j+1] == 'M' && matrix[i-2][j+2] == 'A' && matrix[i-3][j+3] == 'S' {
			sum++
		}

		if j-3 >= 0 && i+3 < len(matrix) && matrix[i+1][j-1] == 'M' && matrix[i+2][j-2] == 'A' && matrix[i+3][j-3] == 'S' {
			sum++
		}
	}

	return sum
}
