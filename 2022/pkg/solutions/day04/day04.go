package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

type Assignment struct {
	Start int
	End   int
}

func (a Assignment) Contains(other Assignment) bool {
	return a.Start <= other.Start && a.End >= other.End
}

func (a Assignment) Overlaps(other Assignment) bool {
	return !(other.Start > a.End || a.Start > other.End)
}

func Part1(input []byte) (runner.Solution, error) {
	count := 0
	assignmentPairs, err := Parse(input)
	if err != nil {
		return runner.Solution{}, err
	}

	for _, pair := range assignmentPairs {
		if pair[0].Contains(pair[1]) || pair[1].Contains(pair[0]) {
			count++
		}
	}

	return runner.Solution{Message: fmt.Sprintf("Contained count: %v", count)}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	count := 0
	assignmentPairs, err := Parse(input)
	if err != nil {
		return runner.Solution{}, err
	}

	for _, pair := range assignmentPairs {
		if pair[0].Overlaps(pair[1]) {
			count++
		}
	}

	return runner.Solution{Message: fmt.Sprintf("Overlaps count: %v", count)}, nil
}

func Parse(input []byte) ([][]Assignment, error) {
	out := [][]Assignment{}
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		pair := make([]Assignment, 2)
		for i, elf := range strings.Split(line, ",") {

			area := strings.Split(elf, "-")
			start, err := strconv.Atoi(area[0])
			if err != nil {
				return out, err
			}
			end, err := strconv.Atoi(area[1])
			if err != nil {
				return out, err
			}
			pair[i] = Assignment{Start: start, End: end}
		}
		out = append(out, pair)
	}
	return out, nil
}
