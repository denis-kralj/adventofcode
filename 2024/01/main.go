package main

import (
	"aoc/utils"
	"fmt"
	"slices"
	"strings"
)

func main() {
	inputLines, _ := utils.GetProblemLines()

	firstList := []int{}
	secondList := []int{}
	simMap := make(map[int][2]int)

	for _, text := range inputLines {
		stringNumberPair := strings.Split(text, "   ")

		firstList = append(firstList, utils.GetNumberFromString(stringNumberPair[0]))
		secondList = append(secondList, utils.GetNumberFromString(stringNumberPair[1]))
	}

	slices.Sort(firstList)
	slices.Sort(secondList)

	sum := 0

	for i := 0; i < len(firstList); i++ {
		if firstList[i] > secondList[i] {
			sum += firstList[i] - secondList[i]
		} else {
			sum += secondList[i] - firstList[i]
		}
		if value, ok := simMap[firstList[i]]; ok {
			simMap[firstList[i]] = [2]int{value[0] + 1, value[1]}
		} else {
			simMap[firstList[i]] = [2]int{1, 0}
		}

		if value, ok := simMap[secondList[i]]; ok {
			simMap[secondList[i]] = [2]int{value[0], value[1] + 1}
		} else {
			simMap[secondList[i]] = [2]int{0, 1}
		}
	}

	simScore := 0

	for key, value := range simMap {
		simScore += key * value[0] * value[1]
	}

	fmt.Println("Total distance between lists: ", sum)
	fmt.Println("Similarity score: ", simScore)
}
