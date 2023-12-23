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

type LightPath struct {
	source    [2]int
	direction Direction
}

type DoneType struct {
	x int
	y int
	d Direction
}

func main() {
	inputLines, _ := utils.GetProblemLines()

	startX, startY := 0, 0
	startDirection := Right

	var sum = getEnergizeLevelForStartingPositionAndDirection(inputLines, startX, startY, startDirection)

	fmt.Println(sum)

	sum2 := 0

	for i := 0; i < len(inputLines); i++ {
		var energizeLevel = getEnergizeLevelForStartingPositionAndDirection(inputLines, i, 0, Right)
		if energizeLevel > sum2 {
			sum2 = energizeLevel
		}
	}

	for i := 0; i < len(inputLines); i++ {
		var energizeLevel = getEnergizeLevelForStartingPositionAndDirection(inputLines, i, len(inputLines[0])-1, Left)
		if energizeLevel > sum2 {
			sum2 = energizeLevel
		}
	}

	for i := 0; i < len(inputLines[0]); i++ {
		var energizeLevel = getEnergizeLevelForStartingPositionAndDirection(inputLines, 0, i, Down)
		if energizeLevel > sum2 {
			sum2 = energizeLevel
		}
	}

	for i := 0; i < len(inputLines[0]); i++ {
		var energizeLevel = getEnergizeLevelForStartingPositionAndDirection(inputLines, len(inputLines)-1, i, Up)
		if energizeLevel > sum2 {
			sum2 = energizeLevel
		}
	}

	fmt.Println(sum2)
}

func getEnergizeLevelForStartingPositionAndDirection(inputLines []string, startX int, startY int, startDirection Direction) int {
	energizedMap := [][]bool{}

	for i, line := range inputLines {
		energizedMap = append(energizedMap, []bool{})
		for j := 0; j < len(line); j++ {
			energizedMap[i] = append(energizedMap[i], false)
		}
	}

	firstDirections := getFirstLightDirections(rune(inputLines[startX][startY]), startDirection)

	lightPaths := []LightPath{}
	for _, direction := range firstDirections {
		lightPaths = append(lightPaths, LightPath{source: [2]int{startX, startY}, direction: direction})
	}

	doneMap := map[DoneType]bool{}

	for len(lightPaths) != 0 {
		path := lightPaths[0]

		energizedMap[path.source[0]][path.source[1]] = true

		for canMove(path, inputLines) {
			path = move(path)

			energizedMap[path.source[0]][path.source[1]] = true
		}

		lightPaths = lightPaths[1:]

		nextPaths := getNextPaths(path, inputLines, doneMap)

		for _, newPath := range nextPaths {
			doneMap[DoneType{x: newPath.source[0], y: newPath.source[1], d: newPath.direction}] = true
		}

		lightPaths = append(lightPaths, nextPaths...)
	}

	sum := 0

	for _, l := range energizedMap {
		for _, e := range l {
			if e {
				sum++
			}
		}
	}

	return sum
}

func getFirstLightDirections(startRune rune, startDirection Direction) []Direction {
	result := []Direction{}
	switch startDirection {
	case Up:
		switch startRune {
		case '.', '|':
			result = append(result, Up)
		case '-':
			result = append(result, Left, Right)
		case '/':
			result = append(result, Right)
		case '\\':
			result = append(result, Left)
		default:
			panic("this shouldn't happen")
		}
	case Down:
		switch startRune {
		case '.', '|':
			result = append(result, Down)
		case '-':
			result = append(result, Left, Right)
		case '/':
			result = append(result, Left)
		case '\\':
			result = append(result, Right)
		default:
			panic("this shouldn't happen")
		}
	case Left:
		switch startRune {
		case '.', '-':
			result = append(result, Left)
		case '|':
			result = append(result, Up, Down)
		case '/':
			result = append(result, Down)
		case '\\':
			result = append(result, Up)
		default:
			panic("this shouldn't happen")
		}
	case Right:
		switch startRune {
		case '.', '-':
			result = append(result, Right)
		case '|':
			result = append(result, Up, Down)
		case '/':
			result = append(result, Up)
		case '\\':
			result = append(result, Down)
		default:
			panic("this shouldn't happen")
		}
	default:
		panic("this shouldn't happen")
	}

	return result
}

func getNextPaths(path LightPath, inputLines []string, doneMap map[DoneType]bool) []LightPath {
	result := []LightPath{}
	source := [2]int{path.source[0], path.source[1]}
	switch path.direction {
	case Up:
		source[0]--
	case Down:
		source[0]++
	case Left:
		source[1]--
	case Right:
		source[1]++
	}

	if source[0] < 0 || source[0] > len(inputLines)-1 || source[1] < 0 || source[1] > len(inputLines[0])-1 {
		return []LightPath{}
	}

	switch inputLines[source[0]][source[1]] {
	case '-':
		if !doneMap[DoneType{x: source[0], y: source[1], d: Left}] {
			result = append(result, LightPath{source: source, direction: Left})
		}
		if !doneMap[DoneType{x: source[0], y: source[1], d: Right}] {
			result = append(result, LightPath{source: source, direction: Right})
		}
	case '|':
		if !doneMap[DoneType{x: source[0], y: source[1], d: Up}] {
			result = append(result, LightPath{source: source, direction: Up})
		}
		if !doneMap[DoneType{x: source[0], y: source[1], d: Down}] {
			result = append(result, LightPath{source: source, direction: Down})
		}
	case '/':
		switch path.direction {
		case Left:
			if !doneMap[DoneType{x: source[0], y: source[1], d: Down}] {
				result = append(result, LightPath{source: source, direction: Down})
			}
		case Right:
			if !doneMap[DoneType{x: source[0], y: source[1], d: Up}] {
				result = append(result, LightPath{source: source, direction: Up})
			}
		case Up:
			if !doneMap[DoneType{x: source[0], y: source[1], d: Right}] {
				result = append(result, LightPath{source: source, direction: Right})
			}
		case Down:
			if !doneMap[DoneType{x: source[0], y: source[1], d: Left}] {
				result = append(result, LightPath{source: source, direction: Left})
			}
		default:
			panic("this shouldn't happen")
		}
	case '\\':
		switch path.direction {
		case Left:
			if !doneMap[DoneType{x: source[0], y: source[1], d: Up}] {
				result = append(result, LightPath{source: source, direction: Up})
			}
		case Right:
			if !doneMap[DoneType{x: source[0], y: source[1], d: Down}] {
				result = append(result, LightPath{source: source, direction: Down})
			}
		case Up:
			if !doneMap[DoneType{x: source[0], y: source[1], d: Left}] {
				result = append(result, LightPath{source: source, direction: Left})
			}
		case Down:
			if !doneMap[DoneType{x: source[0], y: source[1], d: Right}] {
				result = append(result, LightPath{source: source, direction: Right})
			}
		default:
			panic("this shouldn't happen")
		}
	default:
		panic("this shouldn't happen")
	}

	return result
}

func move(path LightPath) LightPath {
	source := [2]int{path.source[0], path.source[1]}
	switch path.direction {
	case Up:
		source[0]--
	case Down:
		source[0]++
	case Left:
		source[1]--
	case Right:
		source[1]++
	}
	return LightPath{direction: path.direction, source: source}
}

func canMove(path LightPath, inputLines []string) bool {
	switch path.direction {
	case Up:
		return path.source[0]-1 >= 0 && strings.ContainsAny(string(inputLines[path.source[0]-1][path.source[1]]), "|.")
	case Down:
		return path.source[0]+1 < len(inputLines) && strings.ContainsAny(string(inputLines[path.source[0]+1][path.source[1]]), "|.")
	case Left:
		return path.source[1]-1 >= 0 && strings.ContainsAny(string(inputLines[path.source[0]][path.source[1]-1]), "-.")
	case Right:
		return path.source[1]+1 < len(inputLines[path.source[0]]) && strings.ContainsAny(string(inputLines[path.source[0]][path.source[1]+1]), "-.")
	default:
		panic("shouldn't happen")
	}
}
