package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func main() {
	inputLines, _ := utils.GetProblemLines()

	expansionCoeficientFirst := 2
	expansionCoeficientSecond := 1_000_000

	emptySpaceMap := getEmptySpace(inputLines)

	galaxyMap := getGalaxyMap(inputLines)

	sum := getSumOfDistances(galaxyMap, emptySpaceMap, expansionCoeficientFirst)
	sum2 := getSumOfDistances(galaxyMap, emptySpaceMap, expansionCoeficientSecond)

	fmt.Println(sum)
	fmt.Println(sum2)
}

func getEmptySpace(inputLines []string) []map[int]bool {
	emptyVerticalSpaceMap := map[int]bool{}
	emptyHorizontalSpaceMap := map[int]bool{}

	for position := range inputLines[0] {
		emptyVerticalSpaceMap[position] = true
	}

	for position := range inputLines {
		emptyHorizontalSpaceMap[position] = true
	}

	for lineIndex, line := range inputLines {
		if strings.Contains(line, "#") {
			indexes := getAllGalaxyIndexes(line)
			for _, index := range indexes {
				emptyVerticalSpaceMap[index] = false
			}
			emptyHorizontalSpaceMap[lineIndex] = false
		}
	}

	return []map[int]bool{emptyHorizontalSpaceMap, emptyVerticalSpaceMap}
}

func getSumOfDistances(galaxyMap map[int][2]int, emptySpaceMap []map[int]bool, expansionCoeficient int) int {
	sum := 0
	for k := 0; k < len(galaxyMap); k++ {
		if k+1 == len(galaxyMap) {
			continue
		}

		for i := k + 1; i < len(galaxyMap); i++ {
			switch {
			case galaxyMap[k][0] > galaxyMap[i][0]:
				sum+= getSumForSection(galaxyMap[i][0], galaxyMap[k][0], emptySpaceMap[0], expansionCoeficient)
			case galaxyMap[k][0] < galaxyMap[i][0]:
				sum+= getSumForSection(galaxyMap[k][0], galaxyMap[i][0], emptySpaceMap[0], expansionCoeficient)
			}

			switch {
			case galaxyMap[k][1] > galaxyMap[i][1]:
				sum+= getSumForSection(galaxyMap[i][1], galaxyMap[k][1], emptySpaceMap[1], expansionCoeficient)
			case galaxyMap[k][1] < galaxyMap[i][1]:
				sum+= getSumForSection(galaxyMap[k][1], galaxyMap[i][1], emptySpaceMap[1], expansionCoeficient)
			}
		}
	}

	return sum
}

func getSumForSection(start int, end int, emptySpaceMap map[int]bool, expansionCoeficient int) int {
	sum:=0
	for j := start; j < end; j++ {
		if emptySpaceMap[j] {
			sum += expansionCoeficient
		} else {
			sum++
		}
	}

	return sum
}

func getGalaxyMap(expandedSpace []string) map[int][2]int {
	result := map[int][2]int{}
	galaxyNum := 0

	for x, l := range expandedSpace {
		for y, c := range l {
			if c == '#' {
				result[galaxyNum] = [2]int{x, y}
				galaxyNum++
			}
		}
	}

	return result
}

func getAllGalaxyIndexes(line string) []int {
	result := []int{}

	start := 0
	end := len(line)
	i := strings.Index(line[start:end], "#")
	for i != -1 {
		result = append(result, i)
		pad := i
		i = strings.Index(line[start+i+1:end], "#")
		if i != -1 {
			i += pad + 1
		}
	}

	return result
}
