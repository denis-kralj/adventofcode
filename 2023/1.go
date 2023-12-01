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
	numbers := [18]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	file, err := os.OpenFile("input_1.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	sum := 0
	for sc.Scan() {
		text := sc.Text()
		var lowestIndex int = 1000
		var lowestIndexDigit string = ""
		var highestIndex int = -1
		var highestIndexDigit string = ""
		var index int = -1
		for i := 0; i < len(numbers); i++ {
			index = strings.Index(text, numbers[i])
			if index != -1 && index < lowestIndex {
				lowestIndex = index
				lowestIndexDigit = numbers[i]
			}

			index = strings.LastIndex(text, numbers[i])
			if index != -1 && index > highestIndex {
				highestIndex = index
				highestIndexDigit = numbers[i]
			}
		}

		number := Todigit(lowestIndexDigit) + Todigit(highestIndexDigit)
		i, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		sum += i
	}

	fmt.Println(sum)
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
