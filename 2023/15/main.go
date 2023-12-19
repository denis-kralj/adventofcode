package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type LenseBox struct {
	boxNumber int
	lenses    []Lense
}

func (box *LenseBox) ApplyLense(data LenseData) {
	if data.isRemove {
		box.RemoveLense(data.lenseLabel)
	} else {
		if box.ContainsLense(data.lenseLabel) {
			box.UpdateLense(data.lenseLabel, data.lenseFocalLength)
		} else {
			lense := Lense{label: data.lenseLabel, focalLength: data.lenseFocalLength}
			box.lenses = append(box.lenses, lense)
		}
	}
}

func (box *LenseBox) RemoveLense(label string) {
	for i := 0; i < len(box.lenses); i++ {
		if box.lenses[i].label == label {
			box.lenses = append(box.lenses[:i], box.lenses[i+1:]...)
		}
	}
}

func (box *LenseBox) UpdateLense(label string, focalLength int) {
	for i, lense := range box.lenses {
		if lense.label == label {
			box.lenses[i].focalLength = focalLength
			return
		}
	}
}

func (box LenseBox) ContainsLense(label string) bool {
	for _, lense := range box.lenses {
		if lense.label == label {
			return true
		}
	}

	return false
}

func (box LenseBox) GetFocusingPower() int {
	result := 0

	if len(box.lenses) != 0 {
		for i, l := range box.lenses {
			result += (box.boxNumber + 1) * (i + 1) * l.focalLength
		}
	}

	return result
}

type Lense struct {
	label       string
	focalLength int
}

type LenseData struct {
	boxNumber        int
	lenseLabel       string
	lenseFocalLength int
	isRemove         bool
}

func main() {
	inputLines, _ := utils.GetProblemLines()

	sum := 0

	lenseBoxes := []LenseBox{}

	for i := 0; i < 256; i++ {
		lenseBoxes = append(lenseBoxes, LenseBox{boxNumber: i, lenses: []Lense{}})
	}

	for _, element := range strings.Split(inputLines[0], ",") {
		sum += hash(element)

		lenseData := getLenseData(element)
		lenseBoxes[lenseData.boxNumber].ApplyLense(lenseData)
	}

	sum2 := 0

	for _, box := range lenseBoxes {
		sum2 += box.GetFocusingPower()
	}

	fmt.Println(sum)
	fmt.Println(sum2)
}

func getLenseData(element string) LenseData {
	result := LenseData{}

	if strings.Contains(element, "-") {
		result.boxNumber = hash(element[:len(element)-1])
		result.isRemove = true
		result.lenseLabel = element[:len(element)-1]
	} else {
		split := strings.Split(element, "=")
		result.boxNumber = hash(split[0])
		result.isRemove = false
		result.lenseLabel = split[0]
		result.lenseFocalLength, _ = strconv.Atoi(split[1])
	}

	return result
}

func hash(input string) int {
	currentValue := 0
	for _, c := range input {
		currentValue += int(c)
		currentValue *= 17
		currentValue %= 256
	}

	return currentValue
}
