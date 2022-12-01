package day01_test

import (
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day01"
)

var sampleInput = []byte(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)

var expectedElves = []day01.Elf{
	{Calories: 5000},
	{Calories: 4000},
	{Calories: 11000},
	{Calories: 24000},
	{Calories: 10000},
}

func TestParser(t *testing.T) {
	elves, err := day01.Parse(sampleInput)
	if err != nil {
		t.Errorf("Did not expect error: %v", err)
	}
	if len(elves) != len(expectedElves) {
		t.Errorf("Number of elves is not %v, but %v", len(expectedElves), len(elves))
	}
}

func TestDay01Part1(t *testing.T) {
	sol, err := day01.Day01Part1(sampleInput)
	if err != nil {
		t.Errorf("Did not expect error: %v", err)
	}
	if sol.Message != "Max Calories are 24000" {
		t.Errorf("Wrong solution: %v", sol)
	}
}
