package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	inputLines, _ := utils.GetProblemLines()

	sum := 0

	for _, line := range inputLines {
		sum += getArrangementCount(line)
	}

	fmt.Println(sum)
}

func getArrangementCount(line string) int {
	sum := 0
	split := strings.Split(line, " ")

	positionCount := len(split[0])

	for _, permutation := range getPermutations(positionCount, split[1]) {
		if isMatch(permutation, split[0]) {
			sum++
		}
	}

	return sum
}

func isMatch(permutation string, template string) bool {
	for i := 0; i < len(permutation); i++ {
		if template[i] != '?' && template[i] != permutation[i] {
			return false
		}
	}

	return true
}

func getPermutations(totalLengthOfHistory int, commaSeparatedDamaged string) []string {

	damagedGlyphs := []string{}

	damageChunkSizes := strings.Split(commaSeparatedDamaged, ",")
	damagedTotalLength := 0

	for _, chunkSize := range damageChunkSizes {
		n, _ := strconv.Atoi(chunkSize)
		damagedTotalLength += n
		damagedGlyphs = append(damagedGlyphs, strings.Repeat("#", n))
	}

	workingLength := totalLengthOfHistory - damagedTotalLength

	distributions := generateDistribution(workingLength, len(damagedGlyphs), make([]int, len(damagedGlyphs)+1))

	legitDistributions := [][]int{}

	for _, dis := range distributions {
		invalid := false
		for i := 1; i < len(dis)-1; i++ {
			if dis[i] == 0 {
				invalid = true
				break
			}
		}

		if invalid {
			continue
		}

		legitDistributions = append(legitDistributions, dis)
	}

	result := []string{}

	for _, lDist := range legitDistributions {
		permutation := buildPermutation(damagedGlyphs, lDist)
		if !utils.Contains(result, permutation) {
			result = append(result, permutation)
		}
	}

	return result
}

func generateDistribution(elements int, containers int, distribution []int) [][]int {
	var permutations [][]int

	if elements == 0 {
		result := make([]int, len(distribution))
		copy(result, distribution)
		permutations = append(permutations, result)
		return permutations
	}

	for i := 0; i <= containers; i++ {
		distribution[containers-i]++
		permutations = append(permutations, generateDistribution(elements-1, containers, distribution)...)
		distribution[containers-i]--
	}

	return permutations
}

func buildPermutation(damagedGlyphs []string, workingGlyphs []int) string {

	result := ""
	for index, element := range workingGlyphs {
		result += strings.Repeat(".", element)
		if index < len(damagedGlyphs) {
			result += damagedGlyphs[index]
		}
	}

	return result
}

// #.#.
// #..#
// .#.#
// 3 permutations 1+2

// #.#..
// #..#.
// #...#
// .#.#.
// .#..#
// ..#.#
// 6 permutations 1+2+3

// #.#...
// #..#..
// #...#.
// #....#
// .#.#..
// .#..#.
// .#...#
// ..#.#.
// ..#..#
// ...#.#
// 10 permutations 1+2+3+4

// #.#.#..
// #.#..#.
// #.#...#
// #..#.#.
// #..#..#
// #...#.#
// .#.#.#.
// .#.#..#
// .#..#.#
// ..#.#.#
// 10 permutations 3+2+1[6] + 2+1[3] + 1

// #.#.#...
// #.#..#..
// #.#...#.
// #.#....#
// #..#.#..
// #..#..#.
// #..#...#
// #...#.#.
// #...#..#
// #....#.#
// .#.#.#..
// .#.#..#.
// .#.#...#
// .#..#.#.
// .#..#..#
// .#...#.#
// ..#.#.#.
// ..#.#..#
// ..#..#.#
// ...#.#.#
// 20 permutations 4+3+2+1[10] + 3+2+1[6] + 2+1[3] + 1

// #.##.###....
// #.##..###...
// #.##...###..
// #.##....###.
// #.##.....###
// #..##.###...
// #..##..###..
// #..##...###.
// #..##....###
// #...##.###..
// #...##..###.
// #...##...###
// #....##.###.
// #....##..###
// #.....##.###
// .#.##.###...
// .#.##..###..
// .#.##...###.
// .#.##....###
// .#..##.###..
// .#..##..###.
// .#..##...###
// .#...##.###.
// .#...##..###
// .#....##.###
// ..#.##.###..
// ..#.##..###.
// ..#.##...###
// ..#..##.###.
// ..#..##..###
// ..#...##.###
// ...#.##.###.
// ...#.##..###
// ...#..##.###
// ....#.##.###

// 35 permutations 5+4+3+2+1[15] + 4+3+2+1[10] + 3+2+1[6] + 2+1[3] + 1
