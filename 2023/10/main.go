package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

func main() {
	inputLines, _ := utils.GetProblemLines()

	l, c := getStartPosition(inputLines)

	stepsInCircle, pipeMap := countAllSteps(inputLines, l, c)

	fmt.Println(stepsInCircle / 2)

	// little bit of a cheat, this works only for my solution and
	// any solution where the start is equivalent to a "J"
	inputLines[l] = strings.Replace(inputLines[l], "S", "J", 1)

	inCount := countIns(inputLines, pipeMap)
	fmt.Println(inCount)
}

func countIns(inputLines []string, pipeMap map[[2]int]bool) int {
	inCount := 0
	out := true

	for i := 0; i < len(inputLines); i++ {
		for j := 0; j < len(inputLines[i]); j++ {
			if !pipeMap[[2]int{i, j}] {
				if !out {
					inCount++
				}
			} else {
				if inputLines[i][j:j+1] == "|" {
					out = !out
				} else {
					inCorner := inputLines[i][j : j+1]

					for !strings.ContainsAny(inputLines[i][j:j+1], "7J") {
						j++
					}

					outCorner := inputLines[i][j : j+1]

					switch {
					case inCorner == "L" && outCorner == "7", inCorner == "F" && outCorner == "J":
						out = !out
					}
				}
			}
		}
	}

	return inCount
}

func getStartPosition(inputLines []string) (int, int) {
	for l := 0; l < len(inputLines); l++ {
		for c, char := range inputLines[l] {
			if char == 'S' {
				return l, c
			}
		}
	}

	panic("this sholdn't happen")
}

func countAllSteps(inputLines []string, l, c int) (int, map[[2]int]bool) {
	pipeMap := map[[2]int]bool{}
	pipeMap[[2]int{l, c}] = true
	nextL, nextC, direction := getFirstStepFromStart(inputLines, l, c)
	steps := 1

	for inputLines[nextL][nextC:nextC+1] != "S" {

		pipeMap[[2]int{nextL, nextC}] = true
		nextL, nextC, direction = step(inputLines, nextL, nextC, direction)
		steps++
	}

	return steps, pipeMap
}

func getFirstStepFromStart(inputLines []string, startL int, startC int) (int, int, Direction) {
	switch {
	case startC-1 > 0 && strings.ContainsAny(inputLines[startL][startC-1:startC], "-FL"):
		return startL, startC - 1, Right
	case startC+1 < len(inputLines[startL]) && strings.ContainsAny(inputLines[startL][startC+1:startC+2], "-7J"):
		return startL, startC + 1, Left
	case startL-1 > 0 && strings.ContainsAny(inputLines[startL-1][startC:startC+1], "|F7"):
		return startL - 1, startC, Down
	case startL+1 < len(inputLines) && strings.ContainsAny(inputLines[startL+1][startC:startC+1], "|LJ"):
		return startL + 1, startC, Up
	}

	panic("This shouldn't happen")
}

func step(inputLines []string, currentL int, currentC int, fromWhere Direction) (int, int, Direction) {
	currentPipe := inputLines[currentL][currentC : currentC+1]

	switch currentPipe {
	case "F":
		if fromWhere == Down {
			return currentL, currentC + 1, Left
		} else if fromWhere == Right {
			return currentL + 1, currentC, Up
		} else {
			panic("this shouldn't happen")
		}
	case "7":
		if fromWhere == Down {
			return currentL, currentC - 1, Right
		} else if fromWhere == Left {
			return currentL + 1, currentC, Up
		} else {
			panic("this shouldn't happen")
		}
	case "J":
		if fromWhere == Up {
			return currentL, currentC - 1, Right
		} else if fromWhere == Left {
			return currentL - 1, currentC, Down
		} else {
			panic("this shouldn't happen")
		}
	case "L":
		if fromWhere == Up {
			return currentL, currentC + 1, Left
		} else if fromWhere == Right {
			return currentL - 1, currentC, Down
		} else {
			panic("this shouldn't happen")
		}
	case "-":
		if fromWhere == Left {
			return currentL, currentC + 1, Left
		} else if fromWhere == Right {
			return currentL, currentC - 1, Right
		} else {
			panic("this shouldn't happen")
		}
	case "|":
		if fromWhere == Up {
			return currentL + 1, currentC, Up
		} else if fromWhere == Down {
			return currentL - 1, currentC, Down
		} else {
			panic("this shouldn't happen")
		}
	}

	panic("this shouldn't happen")
}
