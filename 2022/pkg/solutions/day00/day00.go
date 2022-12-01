package day00

import (
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Day00Part1, Day00Part2}}

func Day00Part1(input []byte) (runner.Solution, error) {
	return runner.Solution{Message: string(input) + "part1"}, nil
}

func Day00Part2(input []byte) (runner.Solution, error) {
	return runner.Solution{Message: string(input) + "part2"}, nil
}
