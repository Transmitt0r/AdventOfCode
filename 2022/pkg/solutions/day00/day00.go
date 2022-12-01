package day00

import (
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnable: Day00}

func Day00(input []byte) (runner.Solution, error) {
	return runner.Solution{Message: string(input)}, nil
}
