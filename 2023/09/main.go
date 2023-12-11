package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputLines, _ := utils.GetProblemLines()

	sum1 := 0
	sum2 := 0

	for _, line := range inputLines {
		inputNumbers := GetInputNumbers(line)
		v1, v2 := ExtrapolateNextValue(inputNumbers)
		sum1 += v1
		sum2 += v2
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}

func GetInputNumbers(line string) []int {
	numberStrings := strings.Split(line, " ")
	numbers := []int{}

	for _, value := range numberStrings {
		number, _ := strconv.Atoi(value)
		numbers = append(numbers, number)
	}

	return numbers
}

func ExtrapolateNextValue(inputs []int) (int, int) {
	dataRows := [][]int{}

	current := inputs
	dataRows = append(dataRows, current)

	for !areAllZero(current) {
		current = calculateDifferences(current)
		dataRows = append(dataRows, current)
	}

	sum1 := 0
	sum2 := 0

	for i := len(dataRows)-1; i >= 0; i-- {
		sum1 += dataRows[i][len(dataRows[i])-1]
		sum2 = dataRows[i][0] - sum2
	}
	return sum1, sum2
}

func calculateDifferences(current []int) []int {
	output := []int{}

	for i := 0; i < len(current)-1; i++ {
		output = append(output, current[i+1]-current[i])
	}

	return output
}

func areAllZero(collection []int) bool {
	for _, num := range collection {
		if num != 0 {
			return false
		}
	}

	return true
}
