package day11_test

import (
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day11"
)

var sampleInput = []byte(`Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1`)

func TestPart1(t *testing.T) {
	sol, err := day11.Part1(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Monkey Business: 10605" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestPart2(t *testing.T) {
	sol, err := day11.Part2(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Monkey Business: 2713310158" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestParse(t *testing.T) {
	monkeys, err := day11.Parse(sampleInput)
	if err != nil {
		t.Error("Parsing error!")
	}
	if len(monkeys) != 4 {
		t.Errorf("Parsing error!")
	}
}
