package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

type Data struct {
	seeds                 []int
	seedToSoil            [][3]int
	soilToFertilizer      [][3]int
	fertilizerToWater     [][3]int
	waterToLight          [][3]int
	lightToTemperature    [][3]int
	temperatureToHumidity [][3]int
	humidityToLocation    [][3]int
}

func main() {
	inputLines, _ := utils.GetProblemLines()

	data := BuildData(inputLines)

	fmt.Println(GetLowestLocationNumberCorrespondingToSeed(data))
	fmt.Println(GetLowestLocationNumberCorrespondingToSeedRanges(data))
}

func GetLowestLocationNumberCorrespondingToSeed(data Data) int {
	smallest := 2147483647
	for _, seed := range data.seeds {
		result := -1

		result = GetMappedValue(seed, data.seedToSoil)
		result = GetMappedValue(result, data.soilToFertilizer)
		result = GetMappedValue(result, data.fertilizerToWater)
		result = GetMappedValue(result, data.waterToLight)
		result = GetMappedValue(result, data.lightToTemperature)
		result = GetMappedValue(result, data.temperatureToHumidity)
		result = GetMappedValue(result, data.humidityToLocation)

		if result < smallest {
			smallest = result
		}
	}
	return smallest
}

func GetLowestLocationNumberCorrespondingToSeedRanges(data Data) int {
	smallest := 2147483647
	for i := 0; i < len(data.seeds); i += 2 {
		fmt.Println(data.seeds[i], "->", data.seeds[i]+data.seeds[i+1]-1)
		for j := data.seeds[i]; j < data.seeds[i]+data.seeds[i+1]-1; j++ {

			result := -1

			result = GetMappedValue(j, data.seedToSoil)
			result = GetMappedValue(result, data.soilToFertilizer)
			result = GetMappedValue(result, data.fertilizerToWater)
			result = GetMappedValue(result, data.waterToLight)
			result = GetMappedValue(result, data.lightToTemperature)
			result = GetMappedValue(result, data.temperatureToHumidity)
			result = GetMappedValue(result, data.humidityToLocation)

			if result < smallest {
				smallest = result
			}
		}

	}
	return smallest
}

func GetMappedValue(mapFrom int, valMap [][3]int) int {
	result := -1
	for _, conversion := range valMap {
		if mapFrom >= conversion[1] && mapFrom < conversion[1]+conversion[2] {
			offset := mapFrom - conversion[1]
			result = conversion[0] + offset
		}

		if result != -1 {
			break
		}
	}

	if result == -1 {
		result = mapFrom
	}

	return result
}

func BuildData(inputLines []string) Data {
	data := Data{}
	data.seeds = []int{}
	data.seedToSoil = [][3]int{}
	data.soilToFertilizer = [][3]int{}
	data.fertilizerToWater = [][3]int{}
	data.waterToLight = [][3]int{}
	data.lightToTemperature = [][3]int{}
	data.temperatureToHumidity = [][3]int{}
	data.humidityToLocation = [][3]int{}

	mapper := ""

	for _, line := range inputLines {

		if line == "" {
			continue
		}

		builtSeeds, seeds := BuildSeedData(line)

		if builtSeeds {
			data.seeds = seeds
			continue
		}

		mapper = EnsureMapper(mapper, line)

		if utils.Contains(validMappers, line) {
			continue
		}

		switch mapper {
		case "seed-to-soil map:":
			data.seedToSoil = append(data.seedToSoil, GetMappingData(line, mapper))
		case "soil-to-fertilizer map:":
			data.soilToFertilizer = append(data.soilToFertilizer, GetMappingData(line, mapper))
		case "fertilizer-to-water map:":
			data.fertilizerToWater = append(data.fertilizerToWater, GetMappingData(line, mapper))
		case "water-to-light map:":
			data.waterToLight = append(data.waterToLight, GetMappingData(line, mapper))
		case "light-to-temperature map:":
			data.lightToTemperature = append(data.lightToTemperature, GetMappingData(line, mapper))
		case "temperature-to-humidity map:":
			data.temperatureToHumidity = append(data.temperatureToHumidity, GetMappingData(line, mapper))
		case "humidity-to-location map:":
			data.humidityToLocation = append(data.humidityToLocation, GetMappingData(line, mapper))
		default:
			panic("we shouldn't get here")
		}
	}

	return data
}

func GetMappingData(line string, mapper string) [3]int {
	numbersString := strings.Split(line, " ")

	numbers := [3]int{}

	for index, numberString := range numbersString {
		i, _ := strconv.Atoi(numberString)
		numbers[index] = i
	}

	return numbers
}

var validMappers = []string{
	"seed-to-soil map:",
	"soil-to-fertilizer map:",
	"fertilizer-to-water map:",
	"water-to-light map:",
	"light-to-temperature map:",
	"temperature-to-humidity map:",
	"humidity-to-location map:",
}

func EnsureMapper(mapper string, line string) string {
	if utils.Contains(validMappers, line) {
		return line
	}

	return mapper
}

func BuildSeedData(line string) (builtSeeds bool, data []int) {
	if strings.Index(line, "seeds:") == 0 {
		res := strings.Split(line, ":")
		seedsStrings := strings.Split(strings.TrimSpace(res[1]), " ")
		for _, seed := range seedsStrings {
			no, _ := strconv.Atoi(seed)
			data = append(data, no)
		}
		return true, data
	}
	return
}
