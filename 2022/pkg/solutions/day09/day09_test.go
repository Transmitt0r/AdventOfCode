package day09_test

import (
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day09"
)

var sampleInput = []byte(`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`)

var sampleInput2 = []byte(`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`)

func TestPart1(t *testing.T) {
	sol, err := day09.Part1(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Unique visited tiles: 13" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestPart2(t *testing.T) {
	sol, err := day09.Part2(sampleInput2)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Unique visited tiles: 36" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestParse(t *testing.T) {
	parsed := day09.Parse(sampleInput)
	if len(parsed) != 8 {
		t.Error("Parsing failed")
	}
}
