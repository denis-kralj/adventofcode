package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	inputLines, _ := utils.GetProblemLines()

	safeSum := 0
	salvagedSum := 0

	for _, text := range inputLines {
		numberList := []int{}
		stringNumbers := strings.Split(text, " ")

		for _, stringNumber := range stringNumbers {
			numberList = append(numberList, GetNumberFromString(stringNumber))
		}

		if isSafe(numberList) {
			safeSum++
		} else if isSalvageable(numberList) {
			salvagedSum++
		}
	}

	fmt.Println("Total safe count: ", safeSum)
	fmt.Println("Total safe + salvaged count: ", safeSum+salvagedSum)
}

func isSalvageable(sequence []int) bool {

	for i := 0; i < len(sequence); i++ {
		slice := []int{}
		slice = append(slice, sequence[:i]...)
		slice = append(slice, sequence[i+1:]...)
		if isSafe(slice) {
			return true
		}
	}
	return false
}

func isSafe(sequence []int) bool {
	if len(sequence) < 2 {
		return true
	}

	if sequence[0] == sequence[1] {
		return false
	}

	isRising := false

	if sequence[0] < sequence[1] {
		isRising = true
	}

	for i := 1; i < len(sequence); i++ {
		if isRising {
			if sequence[i] <= sequence[i-1] {
				return false
			}
		} else {
			if sequence[i] >= sequence[i-1] {
				return false
			}
		}
	}

	for i := 1; i < len(sequence); i++ {
		if isRising {
			if sequence[i]-sequence[i-1] > 3 {
				return false
			}
		} else {
			if sequence[i-1]-sequence[i] > 3 {
				return false
			}
		}
	}

	return true
}

func GetNumberFromString(input string) int {
	number, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	return number
}
