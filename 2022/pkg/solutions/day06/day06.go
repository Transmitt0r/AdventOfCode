package day06

import (
	"errors"
	"fmt"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

func Part1(input []byte) (runner.Solution, error) {
	marker, err := UniqueSlidingWindow(input, 4)
	return runner.Solution{Message: fmt.Sprintf("First marker after: %v", marker)}, err
}

func Part2(input []byte) (runner.Solution, error) {
	marker, err := UniqueSlidingWindow(input, 14)
	return runner.Solution{Message: fmt.Sprintf("First message after: %v", marker)}, err
}

func UniqueSlidingWindow(input []byte, length int) (int, error) {
	for start := 0; start < len(input)-length; start++ {
		uniqueSet := map[byte]bool{}
		foundMultiple := false
		for i := start; i < start+length; i++ {
			if ok := uniqueSet[input[i]]; ok {
				foundMultiple = true
				break
			}
			uniqueSet[input[i]] = true
		}
		if !foundMultiple {
			return start + length, nil
		}
	}
	return 0, errors.New("no marker found")
}
