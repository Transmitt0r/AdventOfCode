package day01

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Day01Part1}}

type Elf struct {
	Calories int
}

func Day01Part1(input []byte) (runner.Solution, error) {
	elves, err := Parse(input)
	if err != nil {
		return runner.Solution{}, err
	}

	maxCalories := 0
	for _, elf := range elves {
		if elf.Calories > maxCalories {
			maxCalories = elf.Calories
		}
	}

	return runner.Solution{Message: fmt.Sprintf("Max Calories are %v", maxCalories)}, nil
}

func Parse(input []byte) ([]Elf, error) {
	elves := []Elf{}
	lines := strings.Split(string(input), "\n")

	currentElf := Elf{}
	for _, line := range lines {
		if line == "" {
			elves = append(elves, currentElf)
			currentElf = Elf{}
			continue
		}
		calories, err := strconv.Atoi(line)
		if err != nil {
			return elves, err
		}
		currentElf.Calories += calories
	}
	elves = append(elves, currentElf)
	return elves, nil
}
