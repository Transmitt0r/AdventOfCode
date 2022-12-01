package day00

import (
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

func Part1(input []byte) (runner.Solution, error) {
	return runner.Solution{Message: string(input) + "part1"}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	return runner.Solution{Message: string(input) + "part2"}, nil
}
