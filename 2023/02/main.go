package main

import (
	"aoc/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Game struct {
	gameNumber int
	redCount   int
	blueCount  int
	greenCount int
}

var limits = Game{blueCount: 14, greenCount: 13, redCount: 12}

func main() {
	inputLines, _ := utils.GetProblemLines()

	var sum int = 0
	var sum2 int = 0
	for _, text := range inputLines {
		gameCandidate := BuildCandidate(text)
		sum2 += gameCandidate.blueCount * gameCandidate.greenCount * gameCandidate.redCount
		if IsValidCandidate(gameCandidate) {
			sum += gameCandidate.gameNumber
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}

func IsValidCandidate(candidate Game) bool {
	return candidate.blueCount <= limits.blueCount &&
		candidate.redCount <= limits.redCount &&
		candidate.greenCount <= limits.greenCount
}

func BuildCandidate(input string) Game {
	candidate := Game{}
	gameResults := strings.Split(input, ":")
	gameNumber := strings.Split(gameResults[0], " ")[1]
	candidate.gameNumber = GetNumberFromString(gameNumber)
	candidate.blueCount = GetMaxCountForColor(gameResults[1], "blue")
	candidate.redCount = GetMaxCountForColor(gameResults[1], "red")
	candidate.greenCount = GetMaxCountForColor(gameResults[1], "green")

	return candidate
}

func GetMaxCountForColor(gamesPlayed string, color string) int {
	games := strings.Split(gamesPlayed, ";")
	var highestCount int = 0
	for i := 0; i < len(games); i++ {
		draws := strings.Split(games[i], ",")
		for j := 0; j < len(draws); j++ {
			draw := strings.Split(strings.TrimSpace(draws[j]), " ")
			if draw[1] == color && GetNumberFromString(draw[0]) > highestCount {
				highestCount = GetNumberFromString(draw[0])
			}
		}
	}

	return highestCount
}

func GetNumberFromString(input string) int {
	number, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	return number
}
