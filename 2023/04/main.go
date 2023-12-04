package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type Scratchcard struct {
	count   int
	winners []string
	draws   []string
}

func main() {
	inputLines, _ := utils.GetProblemLines()
	sum := 0

	scratchcards := []Scratchcard{}

	for _, text := range inputLines {
		points, scratchcard := GetPoints(text)
		sum += points
		scratchcards = append(scratchcards, scratchcard)
	}

	fmt.Println(sum)
	sum2 := 0

	for index, sc := range scratchcards {
		count := GetWinnerCount(sc)
		for innerIndex := index + 1; innerIndex < len(scratchcards) && count > 0; innerIndex, count = innerIndex+1, count-1 {
			scratchcards[innerIndex].count += sc.count
		}

		sum2 += sc.count
	}

	fmt.Println(sum2)
}

func GetWinnerCount(scratchcard Scratchcard) int {
	count := 0

	for _, w := range scratchcard.winners {
		if contains(scratchcard.draws, w) {
			count++
		}
	}

	return count
}

func GetPoints(input string) (int, Scratchcard) {
	points := 0

	numberStrings := strings.Split(strings.Split(input, ":")[1], "|")

	winners := filter(strings.Split(strings.TrimSpace(numberStrings[0]), " "), IsNotEmpty)
	draws := filter(strings.Split(strings.TrimSpace(numberStrings[1]), " "), IsNotEmpty)

	for _, w := range winners {
		if contains(draws, w) {
			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}
	}

	return points, Scratchcard{count: 1, winners: winners, draws: draws}
}

func IsNotEmpty(input string) bool {
	return input != ""
}
func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}

func contains(ss []string, ele string) bool {
	for _, s := range ss {
		if s == ele {
			return true
		}
	}
	return false
}
