package day02

import (
	"fmt"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

type Shape int
type Outcome int

const (
	// shapes one can pick
	Rock     Shape = 1
	Paper    Shape = 2
	Scissors Shape = 3

	// outcome scores
	Lost Outcome = 0
	Draw Outcome = 3
	Won  Outcome = 6
)

var Rules = map[Shape]map[Shape]Outcome{
	Rock: {
		Rock:     Draw,
		Paper:    Lost,
		Scissors: Won,
	},
	Paper: {
		Rock:     Won,
		Paper:    Draw,
		Scissors: Lost,
	},
	Scissors: {
		Rock:     Lost,
		Paper:    Won,
		Scissors: Draw,
	},
}

var ParseMap = map[string]Shape{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var ParseMapOutcomes = map[string]Outcome{
	"X": Lost,
	"Y": Draw,
	"Z": Won,
}

type Round struct {
	Self     Shape
	Opponent Shape
}

type RoundOutcome struct {
	Opponent       Shape
	DesiredOutcome Outcome
}

func (r Round) Score() int {
	return int(r.Self) + int(Rules[r.Self][r.Opponent])
}

func (r RoundOutcome) Score() int {
	var selfPick Shape
	for k := range Rules {
		if Rules[k][r.Opponent] == r.DesiredOutcome {
			selfPick = k
			break
		}
	}
	return int(selfPick) + int(r.DesiredOutcome)
}

func Part1(input []byte) (runner.Solution, error) {
	rounds := Parse(input)
	totalScore := 0

	for _, round := range rounds {
		totalScore += round.Score()
	}

	return runner.Solution{Message: fmt.Sprintf("Total Score: %v", totalScore)}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	rounds := ParseOutcomes(input)
	totalScore := 0

	for _, round := range rounds {
		totalScore += round.Score()
	}

	return runner.Solution{Message: fmt.Sprintf("Total Score: %v", totalScore)}, nil
}

func Parse(input []byte) []Round {
	rounds := []Round{}
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		nextRound := Round{}
		shapesRaw := strings.Split(line, " ")
		nextRound.Opponent = ParseMap[shapesRaw[0]]
		nextRound.Self = ParseMap[shapesRaw[1]]
		rounds = append(rounds, nextRound)
	}
	return rounds
}

func ParseOutcomes(input []byte) []RoundOutcome {
	rounds := []RoundOutcome{}
	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		nextRound := RoundOutcome{}
		shapesRaw := strings.Split(line, " ")
		nextRound.Opponent = ParseMap[shapesRaw[0]]
		nextRound.DesiredOutcome = ParseMapOutcomes[shapesRaw[1]]
		rounds = append(rounds, nextRound)
	}
	return rounds
}
