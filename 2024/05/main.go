package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func main() {
	inputLines, _ := utils.GetProblemLines()

	noFollowList := make(map[int][]int)
	printSequences := [][]int{}
	reachedPrintSequences := false

	for _, text := range inputLines {

		if text == "" {
			reachedPrintSequences = true
			continue
		}

		if reachedPrintSequences {
			printSequences = append(printSequences, extractPrintSequence(text))
		} else {
			left, right := extractSequenceRule(text)
			if collection, ok := noFollowList[right]; ok {
				noFollowList[right] = append(collection, left)
			} else {
				noFollowList[right] = []int{left}
			}
		}
	}

	correctSequenceSumOfMedians := 0
	fixedSequenceSumOfMedians := 0
	for _, sequence := range printSequences {
		if returnSequence, wasFixed := FixSequenceIfNeeded(sequence, noFollowList); wasFixed {
			fixedSequenceSumOfMedians += returnSequence[len(sequence)/2]
		} else {
			correctSequenceSumOfMedians += returnSequence[len(sequence)/2]
		}
	}

	fmt.Println("Sum of all medians in correct print sequences: ", correctSequenceSumOfMedians)
	fmt.Println("Sum of all medians in corrected print sequences: ", fixedSequenceSumOfMedians)
}

func extractSequenceRule(text string) (int, int) {
	stringNumbers := strings.Split(text, "|")
	return utils.GetNumberFromString(stringNumbers[0]), utils.GetNumberFromString(stringNumbers[1])
}

func extractPrintSequence(text string) []int {
	stringNumbers := strings.Split(text, ",")
	numbers := []int{}
	for _, number := range stringNumbers {
		numbers = append(numbers, utils.GetNumberFromString(number))
	}
	return numbers
}

func FixSequenceIfNeeded(sequence []int, noFollowList map[int][]int) ([]int, bool) {
	ptrLeft, ptrRight := 0, 1
	wasFixed := false
	for ptrRight < len(sequence) {
		if utils.Contains(noFollowList[sequence[ptrLeft]], sequence[ptrRight]) {
			sequence[ptrLeft], sequence[ptrRight] = sequence[ptrRight], sequence[ptrLeft]
			wasFixed = true
			ptrRight = ptrLeft + 1
		} else {
			ptrRight++
		}

		if ptrRight == len(sequence) {
			ptrLeft++
			ptrRight = ptrLeft + 1
		}
	}

	return sequence, wasFixed
}
