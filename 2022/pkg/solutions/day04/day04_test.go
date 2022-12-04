package day04_test

import (
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day04"
)

var sampleInput = []byte(`2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`)

func TestPart1(t *testing.T) {
	sol, err := day04.Part1(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Contained count: 2" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestPart2(t *testing.T) {
	sol, err := day04.Part2(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Overlaps count: 4" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestParse(t *testing.T) {
	assignments, err := day04.Parse(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(assignments) != 6 {
		t.Errorf("Wrong amount of assignments: %v", len(assignments))
	}

	if assignments[3][1].Start != 3 || assignments[2][0].End != 7 {
		t.Errorf("Parsing Error!")
	}
}
