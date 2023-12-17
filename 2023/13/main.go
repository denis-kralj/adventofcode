package main

import (
	"aoc/utils"
	"fmt"
)

func main() {
	inputLines, _ := utils.GetProblemLines()

	segments := [][]string{{}}

	index := 0
	segments[index] = []string{}

	for _, line := range inputLines {
		if line == "" {
			segments = append(segments, []string{})
			index++

			continue
		}
		segments[index] = append(segments[index], line)
	}

	values := [][]int{}

	sum := 0
	sum2 := 0

	for i, segment := range segments {
		values = append(values, []int{-1, -1})
		values[i][0], values[i][1] = getValuesForSegment(segment)
		sum += values[i][0]
		sum2 += values[i][1]
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}

func getValuesForSegment(segment []string) (int, int) {
	sum1 := getReflectionLocationValue(segment, -1)

	for l, line := range segment {
		for c, char := range line {
			replace := ' '
			if char == '.' {
				replace = '#'
			} else {
				replace = '.'
			}

			oldLine := segment[l]

			segment[l] = line[:c] + string(replace) + line[c+1:]
			sum2 := getReflectionLocationValue(segment, sum1)
			if sum2 != 0 {
				return sum1, sum2
			}
			segment[l] = oldLine
		}
	}

	panic("this shouldn't happen")
}

func getReflectionLocationValue(segment []string, toCheck int) int {
	for i := 0; i < len(segment)-1; i++ {
		if segment[i] == segment[i+1] && reflects(segment, i, i+1) {
			result := (i + 1) * 100
			if result != toCheck {
				return result
			}

		}
	}

	rotated := rotateBy90Degrees(segment)

	for i := 0; i < len(rotated)-1; i++ {
		if rotated[i] == rotated[i+1] && reflects(rotated, i, i+1) {
			result := (i + 1)
			if result != toCheck {
				return result
			}
		}
	}

	return 0
}

func reflects(lines []string, indexTop int, indexBottom int) bool {

	for indexTop >= 0 && indexBottom < len(lines) {
		if lines[indexTop] != lines[indexBottom] {
			return false
		} else {
			indexTop--
			indexBottom++
		}
	}

	return true
}

func rotateBy90Degrees(segment []string) []string {
	result := []string{}

	for range segment[0] {
		result = append(result, "")
	}

	for i := range segment[0] {
		for _, line := range segment {
			result[i] += line[i : i+1]
		}
	}

	// for i := range result {
	// 	if i >= len(result)-1-i {
	// 		break
	// 	}
	// 	result[i], result[len(result)-1-i] = result[len(result)-1-i], result[i]
	// }

	return result
}
