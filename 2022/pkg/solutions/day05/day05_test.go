package day05_test

import (
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day05"
)

var sampleInput = []byte(`    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2`)

func TestPart1(t *testing.T) {
	sol, err := day05.Part1(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Top Crates: CMZ" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestPart2(t *testing.T) {
	sol, err := day05.Part2(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Top Crates: MCD" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestParseStart(t *testing.T) {
	sampleStart := []string{
		"    [D]    ",
		"[N] [C]    ",
		"[Z] [M] [P]",
		" 1   2   3 ",
	}

	res := day05.ParseStart(sampleStart)
	if res.Len() != 3 {
		t.Errorf("Parsing Error")
	}
}

func TestParseEnd(t *testing.T) {
	sampleEnd := []string{
		"move 1 from 2 to 1",
		"move 3 from 1 to 3",
		"move 2 from 2 to 1",
		"move 1 from 1 to 2",
	}

	res := day05.ParseEnd(sampleEnd)
	if len(res) != 4 {
		t.Errorf("Parsing Error")
	}
}

func TestParse(t *testing.T) {
	stack, instructions := day05.Parse(sampleInput)
	if stack.Len() != 3 || len(instructions) != 4 {
		t.Errorf("Parsing Error!")
	}
}
