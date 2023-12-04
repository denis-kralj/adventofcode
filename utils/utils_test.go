package utils_test

import (
	"aoc/utils"
	"fmt"
	"regexp"
	"testing"
)

func TestGetProblemLinesWithDefaultPath(t *testing.T) {
	want := regexp.MustCompile("^this is a test input$")
	output, err := utils.GetProblemLines()
	if !want.MatchString(output[0]) || err != nil {
		t.Fatalf(`utils.GetProblemLines() = %q, %v, want match for %#q, nil`, output[0], err, want)
	}
}

func TestGetProblemLinesWithCustomPath(t *testing.T) {
	want := regexp.MustCompile("^this is also a test input$")
	output, err := utils.GetProblemLines("input_alternate_path.txt")
	if !want.MatchString(output[0]) || err != nil {
		t.Fatalf(`utils.GetProblemLines("input_alternate_path.txt") = %q, %v, want match for %#q, nil`, output[0], err, want)
	}
}

func TestGetProblemLinesWithFaultyPath(t *testing.T) {
	output, err := utils.GetProblemLines("non_exists.txt")
	if err == nil {
		t.Fatalf(`utils.GetProblemLines("non_exists.txt") = %q, %v, want error`, output[0], err)
	}
}

func TestGetProblemLinesWithMoreThanOneParam(t *testing.T) {
	output, err := utils.GetProblemLines("non_exists.txt", "another_param")
	if err == nil || fmt.Sprint(err) != "too many arguments" {
		t.Fatalf(`utils.GetProblemLines("non_exists.txt") = %q, %v, want error`, output[0], err)
	}
}
