package day03_test

import (
	"fmt"
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day03"
)

var sampleInput = []byte(`vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`)

func TestPart1(t *testing.T) {
	sol, err := day03.Part1(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Priority Sum: 157" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestPart2(t *testing.T) {
	sol, err := day03.Part2(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Priority Sum: 70" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestCalculatePriority(t *testing.T) {
	testCases := []struct {
		Expected int
		Input    rune
	}{
		{
			Expected: 16,
			Input:    'p',
		},
		{
			Expected: 38,
			Input:    'L',
		},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("Test %v", test), func(t *testing.T) {
			val := day03.CalculatePriority(test.Input)
			if val != test.Expected {
				t.Errorf("Expected %v, got %v", test.Expected, val)
			}
		})
	}
}
