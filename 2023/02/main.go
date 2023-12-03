package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	gameNumber int
	redCount   int
	blueCount  int
	greenCount int
}

func main() {
	file := Getfile("input_2.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)
	limits := Game{}
	limits.blueCount = 14
	limits.greenCount = 13
	limits.redCount = 12

	var sum int = 0
	var sum2 int = 0
	for sc.Scan() {
		text := sc.Text()
		gameCandidate := Buildcandidate(text)
		sum2 += gameCandidate.blueCount * gameCandidate.greenCount * gameCandidate.redCount
		if Isvalidcandidate(gameCandidate, limits) {
			sum += gameCandidate.gameNumber
		}
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}

func Isvalidcandidate(candidate Game, limits Game) bool {
	return candidate.blueCount <= limits.blueCount &&
		candidate.redCount <= limits.redCount &&
		candidate.greenCount <= limits.greenCount
}

func Buildcandidate(input string) Game {
	candidate := Game{}
	gameResults := strings.Split(input, ":")
	gameNumber := strings.Split(gameResults[0], " ")[1]
	candidate.gameNumber = Getnumberfromstring(gameNumber)
	candidate.blueCount = Getmaxcountforcolor(gameResults[1], "blue")
	candidate.redCount = Getmaxcountforcolor(gameResults[1], "red")
	candidate.greenCount = Getmaxcountforcolor(gameResults[1], "green")

	return candidate
}

func Getmaxcountforcolor(gamesPlayed string, color string) int {
	games := strings.Split(gamesPlayed, ";")
	var highestCount int = 0
	for i := 0; i < len(games); i++ {
		draws := strings.Split(games[i], ",")
		for j := 0; j < len(draws); j++ {
			draw := strings.Split(strings.TrimSpace(draws[j]), " ")
			if draw[1] == color && Getnumberfromstring(draw[0]) > highestCount {
				highestCount = Getnumberfromstring(draw[0])
			}
		}
	}

	return highestCount
}

func Getnumberfromstring(input string) int {
	number, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	return number
}

func Getfile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	return file
}
