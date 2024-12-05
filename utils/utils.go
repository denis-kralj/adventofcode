package utils

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strconv"
)

type Comparable interface {
	int | float64 | string
}

func Contains[T Comparable](array []T, element T) bool {
	for _, e := range array {
		if e == element {
			return true
		}
	}
	return false
}

func GetProblemLines(arg ...string) ([]string, error) {
	path := ""
	data := []string{}
	if len(arg) == 0 {
		path = "input.txt"
	} else if len(arg) == 1 {
		path = arg[0]
	} else {
		return data, errors.New("too many arguments")
	}

	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	if err != nil {
		return data, err
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	output := []string{}

	for sc.Scan() {
		output = append(output, sc.Text())
	}

	return output, nil
}

func GetNumberFromString(input string) int {
	number, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	return number
}
