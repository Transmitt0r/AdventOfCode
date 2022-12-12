package day12_test

import (
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day12"
)

var sampleInput = []byte(`Sabqponm
abcryxxl
accszExk
acctuvwj
abdefghi`)

func TestPart1(t *testing.T) {
	sol, err := day12.Part1(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Least steps to Target: 31" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestPart2(t *testing.T) {
	sol, err := day12.Part2(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Least steps from any starting a: 29" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestParse(t *testing.T) {
	emap := day12.Parse(sampleInput)
	expectedStart := day12.Point{X: 0, Y: 0}
	expectedTarget := day12.Point{X: 5, Y: 2}
	if emap.Start != expectedStart || emap.Target != expectedTarget {
		t.Errorf("Parsing error!")
	}
}
