package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file := Getfile("input_1.txt")
	defer file.Close()

	sc := bufio.NewScanner(file)

	var biggest int = 0
	var secondBiggest int = 0
	var thirdBiggest int = 0
	var subsum int = 0

	for sc.Scan() {
		text := sc.Text()
		if text == "" {
			if subsum > biggest {
				thirdBiggest = secondBiggest
				secondBiggest = biggest				
				biggest = subsum
			}
			subsum = 0
		} else {
			subsum += Getnumber(text)
		}
	}

	fmt.Println(biggest + secondBiggest + thirdBiggest)
}

func Getnumber(input string) int {
	number, err := strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	return number
}

func Getfile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatal(err)
	}

	return file
}
