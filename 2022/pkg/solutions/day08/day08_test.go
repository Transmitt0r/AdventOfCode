package day08_test

import (
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day08"
)

var sampleInput = []byte(`30373
25512
65332
33549
35390`)

func TestPart1(t *testing.T) {
	sol, err := day08.Part1(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Number of visible trees: 21" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestPart2(t *testing.T) {
	sol, err := day08.Part2(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Highest scenic score: 8" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestParse(t *testing.T) {
	parsed := day08.Parse(sampleInput)
	xLen, yLen := parsed.Size()
	if xLen != 5 && yLen != 5 {
		t.Error("Parsing failed")
	}
}
