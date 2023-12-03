package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type PartNumberData struct {
	partNumber     int
	symbolLocation [2]int
	symbol         string
}

func main() {
	file := Getfile("input_3.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)

	var top = ""
	var mid = ""
	var bot = ""
	var lineNumber = 1

	var sum int = 0
	var partNumberDataCollection = []PartNumberData{}

	for sc.Scan() {
		text := sc.Text()
		top = mid
		mid = bot
		bot = text

		if len(mid) == 0 {
			continue
		}

		var numberString string = ""

		for index, character := range mid {
			if _, err := strconv.Atoi(string(character)); err == nil {
				numberString += string(character)
			} else {
				if numberString != "" {
					value, partNumberData := Getpartnumberorzero(top, mid, bot, numberString, index, lineNumber)
					sum += value
					partNumberDataCollection = append(partNumberDataCollection, partNumberData...)
					numberString = ""
				}
			}
		}

		if numberString != "" {
			value, partNumberData := Getpartnumberorzero(top, mid, bot, numberString, len(mid), lineNumber)
			sum += value
			partNumberDataCollection = append(partNumberDataCollection, partNumberData...)
		}

		lineNumber++
	}

	top = mid
	mid = bot
	bot = ""

	var numberString string = ""

	for index, character := range mid {
		if _, err := strconv.Atoi(string(character)); err == nil {
			numberString += string(character)
		} else {
			if numberString != "" {
				value, partNumberData := Getpartnumberorzero(top, mid, bot, numberString, index, lineNumber)
				sum += value
				partNumberDataCollection = append(partNumberDataCollection, partNumberData...)
				numberString = ""
			}
		}
	}

	if numberString != "" {
		value, partNumberData := Getpartnumberorzero(top, mid, bot, numberString, len(mid), lineNumber)
		sum += value
		partNumberDataCollection = append(partNumberDataCollection, partNumberData...)
	}

	fmt.Println(sum)
	fmt.Println(GetGearRatioSum(partNumberDataCollection))
}

func GetGearRatioSum(partNumberdataCollection []PartNumberData) int {
	gears := filter(partNumberdataCollection, IsGear)
	gearMap := make(map[[2]int][]PartNumberData)
	sum := 0

	for _, g := range gears {
		if val, ok := gearMap[g.symbolLocation]; ok {
			gearMap[g.symbolLocation] = append(val, g)
		} else {
			gearMap[g.symbolLocation] = []PartNumberData{g}
		}
	}

	for _, m := range gearMap {
		if len(m) == 2 {
			sum += m[0].partNumber * m[1].partNumber
		}
	}
	return sum
}

func IsGear(partNumberData PartNumberData) bool {
	return partNumberData.symbol == "*"
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func Getpartnumberorzero(top string, mid string, bot string, numberString string, indexAfterNumber int, lineNumber int) (int, []PartNumberData) {
	symbols := "*=-+&#%/@$"
	number, _ := strconv.Atoi(numberString)
	partNumberData := []PartNumberData{}

	startIndex := indexAfterNumber - len(numberString) - 1

	if startIndex < 0 {
		startIndex = 0
	}

	endIndex := indexAfterNumber + 1

	if endIndex > len(mid) {
		endIndex = len(mid)
	}

	if len(top) == len(mid) && strings.ContainsAny(top[startIndex:endIndex], symbols) {
		for index, character := range top[startIndex:endIndex] {
			if strings.ContainsAny(string(character), symbols) {
				partNumberData = append(partNumberData, PartNumberData{partNumber: number, symbolLocation: [2]int{lineNumber - 1, index + startIndex}, symbol: string(character)})
			}
		}
	}

	if strings.ContainsAny(mid[startIndex:endIndex], symbols) {
		for index, character := range mid[startIndex:endIndex] {
			if strings.ContainsAny(string(character), symbols) {
				partNumberData = append(partNumberData, PartNumberData{partNumber: number, symbolLocation: [2]int{lineNumber, index + startIndex}, symbol: string(character)})
			}
		}
	}

	if len(bot) == len(mid) && strings.ContainsAny(bot[startIndex:endIndex], symbols) {
		for index, character := range bot[startIndex:endIndex] {
			if strings.ContainsAny(string(character), symbols) {
				partNumberData = append(partNumberData, PartNumberData{partNumber: number, symbolLocation: [2]int{lineNumber + 1, index + startIndex}, symbol: string(character)})
			}
		}
	}

	if len(partNumberData) == 0 {
		return 0, partNumberData
	}
	return number, partNumberData
}

func Getfile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	return file
}
