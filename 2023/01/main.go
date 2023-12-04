package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

var digitsNumeric []string = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9",
}

var digitsString []string = []string{
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func main() {
	inputLines, _ := utils.GetProblemLines()

	sum := 0
	sum2 := 0

	for _, text := range inputLines {
		firstDigit, secondDigit := GetDigitStrings(text, digitsNumeric)
		sum += BuildNumber(firstDigit, secondDigit)
		firstDigit, secondDigit = GetDigitStrings(text, append(digitsNumeric, digitsString...))
		sum2 += BuildNumber(firstDigit, secondDigit)
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}

func GetDigitStrings(input string, possibleDigits []string) (string, string) {
	var lowestIndex int = 1000 // known that no index will be this high
	var lowestIndexDigit string = ""
	var highestIndex int = -1
	var highestIndexDigit string = ""

	var index int = -1
	for i := 0; i < len(possibleDigits); i++ {
		index = strings.Index(input, possibleDigits[i])
		if index != -1 && index < lowestIndex {
			lowestIndex = index
			lowestIndexDigit = possibleDigits[i]
		}

		index = strings.LastIndex(input, possibleDigits[i])
		if index != -1 && index > highestIndex {
			highestIndex = index
			highestIndexDigit = possibleDigits[i]
		}
	}

	return lowestIndexDigit, highestIndexDigit
}

func BuildNumber(firstDigit string, secondDigit string) int {
	numberString := ToDigit(firstDigit) + ToDigit(secondDigit)

	number, err := strconv.Atoi(numberString)

	if err != nil {
		log.Fatal(err)
	}

	return number
}

func ToDigit(number string) string {
	switch number {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return number
	}
}
