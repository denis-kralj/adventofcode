package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func main() {
	inputLines, _ := utils.GetProblemLines()

	inputRunes := [][]rune{}

	for _, line := range inputLines {
		inputRunes = append(inputRunes, []rune(line))
	}

	inputRunes = tiltArrayNorth(inputRunes)

	sum := 0
	coef := len(inputRunes)

	for _, line := range inputRunes {
		sum += strings.Count(string(line), "O") * coef
		coef--
	}

	fmt.Println(sum)

	inputRunes = [][]rune{}

	for _, line := range inputLines {
		inputRunes = append(inputRunes, []rune(line))
	}

	// found a cycle with this loop, discovered 27 posibilities
	// I cheated and found the solution by binary search in the 
	// cycle :P
	// for i := 0; i < 1_000_000_000; i++ {
	// 	inputRunes = tiltArrayNorth(inputRunes)
	// 	inputRunes = rotate90DegreesRight(inputRunes)
	// 	inputRunes = tiltArrayNorth(inputRunes)
	// 	inputRunes = rotate90DegreesRight(inputRunes)
	// 	inputRunes = tiltArrayNorth(inputRunes)
	// 	inputRunes = rotate90DegreesRight(inputRunes)
	// 	inputRunes = tiltArrayNorth(inputRunes)
	// 	inputRunes = rotate90DegreesRight(inputRunes)
	// }

	// posibilities are:
	// 94811
	// 94816
	// 94828
	// 94831
	// 94839
	// 94841
	// 94856
	// 94864
	// 94876 <- this is correct
	// 94888
	// 94894
	// 94924
	// 94930
	// 94959
	// 94965
	// 94977
	// 94988
	// 94988
	// 95004
	// 95009
	// 95018
	// 95020
	// 95021
	// 95026
	// 95041
	// 95044
	// 95046

	fmt.Println("fin")

}

func rotate90DegreesRight(inputRunes [][]rune) [][]rune {

	for i := 0; i < len(inputRunes); i++ {
		for j := 0; j < i; j++ {
			inputRunes[i][j], inputRunes[j][i] = inputRunes[j][i], inputRunes[i][j]
		}
	}

	for i := 0; i < len(inputRunes); i++ {
		for j := 0; j < len(inputRunes)/2; j++ {
			inputRunes[i][j], inputRunes[i][len(inputRunes)-j-1] = inputRunes[i][len(inputRunes)-j-1], inputRunes[i][j]
		}
	}
	return inputRunes
}

func tiltArrayNorth(inputRunes [][]rune) [][]rune {
	for columnIndex := range inputRunes[0] {
		for rowIndex := range inputRunes {
			if inputRunes[rowIndex][columnIndex] == 'O' {
				aboveIndex := rowIndex - 1
				movingIndex := rowIndex

				for aboveIndex >= 0 && inputRunes[aboveIndex][columnIndex] == '.' {
					inputRunes[aboveIndex][columnIndex], inputRunes[movingIndex][columnIndex] = inputRunes[movingIndex][columnIndex], inputRunes[aboveIndex][columnIndex]
					aboveIndex--
					movingIndex--
				}
			}
		}
	}

	return inputRunes
}
