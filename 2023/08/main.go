package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

type targetingFunction func(string) bool

func main() {
	inputLines, _ := utils.GetProblemLines()

	directions := inputLines[0]

	nodes := map[string][2]string{}

	for i := 2; i < len(inputLines); i++ {
		node, left, right := GetNode(inputLines[i])
		nodes[node] = [2]string{left, right}
	}

	steps := GetStepCount("AAA", directions, nodes, func(target string) bool { return target == "ZZZ" })

	fmt.Println(steps)

	stepList := []int{}

	for key := range nodes {
		if key[2:3] == "A" {
			stepList = append(stepList, GetStepCount(key, directions, nodes, func(target string) bool { return target[2:3] == "Z" }))
		}
	}

	fmt.Println(LCM(stepList[0], stepList[1], stepList[2:]...))
}

func GetStepCount(seed string, directions string, nodeMap map[string][2]string, isTarget targetingFunction) int {
	i := 0
	steps := 0
	current := seed

	for !isTarget(current) {
		if directions[i:i+1] == "L" {
			current = nodeMap[current][0]
		} else {
			current = nodeMap[current][1]
		}

		i++

		if i == len(directions) {
			i = 0
		}

		steps++
	}

	return steps
}

func GetNode(line string) (string, string, string) {
	keyDirectionSplit := strings.Split(line, " = ")

	leftRight := keyDirectionSplit[1][1 : len(keyDirectionSplit[1])-1]
	leftRightSplit := strings.Split(leftRight, ", ")

	return keyDirectionSplit[0], leftRightSplit[0], leftRightSplit[1]
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
