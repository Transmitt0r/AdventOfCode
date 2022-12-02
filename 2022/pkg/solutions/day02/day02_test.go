package day02_test

import (
	"fmt"
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day02"
)

var sampleInput = []byte(`A Y
B X
C Z`)

func TestPart1(t *testing.T) {
	sol, err := day02.Part1(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Total Score: 15" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestPart2(t *testing.T) {
	sol, err := day02.Part2(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Total Score: 12" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestParse(t *testing.T) {
	rounds := day02.Parse(sampleInput)

	if len(rounds) != 3 {
		t.Errorf("Expected 3 rounds, got %v", len(rounds))
	}
	if rounds[0].Opponent != day02.Rock || rounds[0].Self != day02.Paper {
		t.Errorf("Parsing error")
	}
}

func TestParseOutcomes(t *testing.T) {
	rounds := day02.ParseOutcomes(sampleInput)

	if len(rounds) != 3 {
		t.Errorf("Expected 3 rounds, got %v", len(rounds))
	}
	if rounds[0].Opponent != day02.Rock || rounds[0].DesiredOutcome != day02.Draw {
		t.Errorf("Parsing error")
	}
}

func TestRound(t *testing.T) {
	rounds := []struct {
		round day02.Round
		score int
	}{
		{day02.Round{day02.Rock, day02.Scissors}, 7},
		{day02.Round{day02.Paper, day02.Scissors}, 2},
		{day02.Round{day02.Paper, day02.Paper}, 5},
	}
	for _, test := range rounds {
		t.Run(fmt.Sprintf("Testing %v", test.round), func(t *testing.T) {
			if test.score != test.round.Score() {
				t.Errorf("Wrong score, expected %v got %v for %v", test.score, test.round.Score(), test.round)
			}
		})
	}
}

func TestRoundOutcome(t *testing.T) {
	rounds := []struct {
		round day02.RoundOutcome
		score int
	}{
		{day02.RoundOutcome{day02.Rock, day02.Won}, 8},
		{day02.RoundOutcome{day02.Paper, day02.Lost}, 1},
		{day02.RoundOutcome{day02.Paper, day02.Draw}, 5},
	}
	for _, test := range rounds {
		t.Run(fmt.Sprintf("Testing %v", test.round), func(t *testing.T) {
			if test.score != test.round.Score() {
				t.Errorf("Wrong score, expected %v got %v for %v", test.score, test.round.Score(), test.round)
			}
		})
	}
}
