package utils

import (
	"bufio"
	"errors"
	"os"
)

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
