package day11_test

import (
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day11"
)

var sampleMonkeys = []day11.Monkey{
	{
		Items:       []int{79, 98},
		Operation:   func(i int) int { return i * 19 },
		TestDivisor: 23,
		IfTrue:      2,
		IfFalse:     3,
	},
	{
		Items:       []int{54, 65, 75, 74},
		Operation:   func(i int) int { return i + 6 },
		TestDivisor: 19,
		IfTrue:      2,
		IfFalse:     0,
	},
	{
		Items:       []int{79, 60, 97},
		Operation:   func(i int) int { return i * i },
		TestDivisor: 13,
		IfTrue:      1,
		IfFalse:     3,
	},
	{
		Items:       []int{74},
		Operation:   func(i int) int { return i + 3 },
		TestDivisor: 17,
		IfTrue:      0,
		IfFalse:     1,
	},
}

func TestPart1(t *testing.T) {
	sol := day11.Part1Monkey(sampleMonkeys)
	if sol != 10605 {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestPart2(t *testing.T) {
	sol := day11.Part2Monkey(sampleMonkeys)
	if sol != 2713310158 {
		t.Errorf("Wrong solution %v", sol)
	}
}
