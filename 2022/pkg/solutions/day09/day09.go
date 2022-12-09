package day09

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

type InstDirection string

const (
	DirectionUp    InstDirection = "U"
	DirectionDown  InstDirection = "D"
	DirectionLeft  InstDirection = "L"
	DirectionRight InstDirection = "R"
)

var directionMap = map[string]InstDirection{
	"U": DirectionUp,
	"D": DirectionDown,
	"L": DirectionLeft,
	"R": DirectionRight,
}

type Instruction struct {
	Direction InstDirection
	Steps     int
}

type Coordinate struct {
	X int
	Y int
}

type Grid struct {
	Knots         []Coordinate
	visitedByTail map[Coordinate]bool
}

func NewGrid(knots int) *Grid {
	return &Grid{Knots: make([]Coordinate, knots), visitedByTail: map[Coordinate]bool{{0, 0}: true}}
}

func (g Grid) VisitedByTail() int {
	count := 0
	for _, v := range g.visitedByTail {
		if v {
			count++
		}
	}
	return count
}

func (g *Grid) ApplyInstruction(inst Instruction) {
	for i := 0; i < inst.Steps; i++ {
		switch inst.Direction {
		case DirectionUp:
			g.Knots[0].Y++
		case DirectionDown:
			g.Knots[0].Y--
		case DirectionLeft:
			g.Knots[0].X--
		case DirectionRight:
			g.Knots[0].X++
		}
		for k := 1; k < len(g.Knots); k++ {
			g.moveKnot(k)
		}
	}
}

func abs(num int) (int, int) {
	// return absolute value and sign
	if num < 0 {
		return -num, -1
	}
	return num, 1
}

func (g *Grid) moveKnot(knot int) {
	if g.Knots[knot-1].X == g.Knots[knot].X || g.Knots[knot-1].Y == g.Knots[knot].Y {
		// move directly up or down
		xDist, xSign := abs(g.Knots[knot-1].X - g.Knots[knot].X)
		yDist, ySign := abs(g.Knots[knot-1].Y - g.Knots[knot].Y)

		if xDist >= 2 && g.Knots[knot-1].Y == g.Knots[knot].Y {
			g.Knots[knot].X += xSign
		}

		if yDist >= 2 && g.Knots[knot-1].X == g.Knots[knot].X {
			g.Knots[knot].Y += ySign
		}
	} else {
		// move diagonally
		xDist, xSign := abs(g.Knots[knot-1].X - g.Knots[knot].X)
		yDist, ySign := abs(g.Knots[knot-1].Y - g.Knots[knot].Y)
		if xDist >= 2 || yDist >= 2 {
			g.Knots[knot].X += xSign
			g.Knots[knot].Y += ySign
		}
	}
	if knot == len(g.Knots)-1 {
		g.visitedByTail[g.Knots[knot]] = true
	}
}

func Part1(input []byte) (runner.Solution, error) {
	instructions := Parse(input)
	grid := NewGrid(2)

	for _, inst := range instructions {
		grid.ApplyInstruction(inst)
	}

	return runner.Solution{Message: fmt.Sprintf("Unique visited tiles: %v", grid.VisitedByTail())}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	instructions := Parse(input)
	grid := NewGrid(10)

	for _, inst := range instructions {
		grid.ApplyInstruction(inst)
	}

	return runner.Solution{Message: fmt.Sprintf("Unique visited tiles: %v", grid.VisitedByTail())}, nil
}

func Parse(input []byte) []Instruction {
	lines := strings.Split(string(input), "\n")
	insts := []Instruction{}

	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		direction := directionMap[splitLine[0]]
		count, _ := strconv.Atoi(splitLine[1])

		insts = append(insts, Instruction{direction, count})
	}
	return insts
}
