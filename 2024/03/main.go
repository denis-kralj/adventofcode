package main

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	inputLines, _ := utils.GetProblemLines()

	sum := 0
	optimizedSum := 0
	mulEnabled := true
	for _, text := range inputLines {
		expression := regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)
		for _, match := range expression.FindAllString(text, -1) {
			if match == "don't()" {
				mulEnabled = false
				continue
			} else if match == "do()" {
				mulEnabled = true
				continue
			}
			sum += multiplyExpression(match)

			if mulEnabled {
				optimizedSum += multiplyExpression(match)
			}
		}
	}

	fmt.Println("Sum: ", sum)
	fmt.Println("Optimized Sum: ", optimizedSum)
}

func multiplyExpression(expression string) int {
	expression = strings.TrimPrefix(expression, "mul(")
	expression = strings.TrimSuffix(expression, ")")

	numbers := strings.Split(expression, ",")

	return utils.GetNumberFromString(numbers[0]) * utils.GetNumberFromString(numbers[1])
}
