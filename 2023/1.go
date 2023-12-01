package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	digits := [18]string{
		"1", "2", "3", "4", "5", "6", "7", "8", "9",
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
	}

	file := Getfile("input_1.txt")
	defer file.Close()

	sum := 0
	sc := bufio.NewScanner(file)

	for sc.Scan() {
		text := sc.Text()
		firstDigit, secondDigit := Getdigitstrings(text, digits)
		sum += Buildnumber(firstDigit, secondDigit)
	}

	fmt.Println(sum)
}

func Getfile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func Getdigitstrings(input string, possibleDigits [18]string) (string, string) {
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

func Buildnumber(firstDigit string, secondDigit string) int {
	numberString := Todigit(firstDigit) + Todigit(secondDigit)

	number, err := strconv.Atoi(numberString)

	if err != nil {
		log.Fatal(err)
	}

	return number
}

func Todigit(number string) string {
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
