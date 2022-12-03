package day03

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

func CalculatePriority(in rune) int {
	asciiVal := int(in)
	if unicode.IsLower(in) {
		return asciiVal - 96
	} else {
		return asciiVal - 64 + 26
	}
}

func Part1(input []byte) (runner.Solution, error) {
	prioritySum := 0
	for _, line := range strings.Split(string(input), "\n") {
		sets := []map[rune]bool{{}, {}}

		for i, char := range line {
			if i < len(line)/2 {
				sets[0][char] = true
			} else {
				sets[1][char] = true
			}
		}

		for k := range sets[0] {
			if ok := sets[1][k]; ok {
				prioritySum += CalculatePriority(k)

			}
		}
	}
	return runner.Solution{Message: fmt.Sprintf("Priority Sum: %v", prioritySum)}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	prioritySum := 0
	elfGroup := []string{}
	for _, line := range strings.Split(string(input), "\n") {
		if len(elfGroup) < 3 {
			elfGroup = append(elfGroup, line)
			if len(elfGroup) < 3 {
				continue
			}
		}

		sets := []map[rune]bool{{}, {}, {}}
		for i, elf := range elfGroup {
			for _, r := range elf {
				sets[i][r] = true
			}
		}

		for k := range sets[0] {
			if ok1, ok2 := sets[1][k], sets[2][k]; ok1 && ok2 {
				prioritySum += CalculatePriority(k)
			}
		}

		elfGroup = []string{}
	}
	return runner.Solution{Message: fmt.Sprintf("Priority Sum: %v", prioritySum)}, nil
}
