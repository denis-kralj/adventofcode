package main

import "fmt"

var input = [4][2]int{{57, 291}, {72, 1172}, {69, 1176}, {92, 2026}}

func main() {
	product := 1
	for _, race := range input {
		product *= GetNumberOfWaysToWin(race)
	}

	fmt.Println(product)
	fmt.Println(GetNumberOfWaysToWin([2]int{57726992, 291117211762026}))
}

func GetNumberOfWaysToWin(inputs [2]int) int {
	fmt.Println(inputs[0], "=>", inputs[1])
	sum := 0
	for i := 1; i < inputs[0]; i++ {
		raceTime := inputs[0] - i
		distanceCovered := raceTime * i
		if distanceCovered >= inputs[1] {
			sum++
		}
	}
	return sum
}
