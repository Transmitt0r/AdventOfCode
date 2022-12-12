package day12

import (
	"fmt"
	"math"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

type Point struct {
	X int
	Y int
}

type ElevationMap struct {
	Elevations [][]int
	Start      Point
	Target     Point
}

func (m ElevationMap) possibleMoves(p Point) []Point {
	possible := []Point{}
	if p.X > 0 {
		possible = append(possible, Point{p.X - 1, p.Y})
	}
	if p.X < len(m.Elevations[p.Y])-1 {
		possible = append(possible, Point{p.X + 1, p.Y})
	}
	if p.Y > 0 {
		possible = append(possible, Point{p.X, p.Y - 1})
	}
	if p.Y < len(m.Elevations)-1 {
		possible = append(possible, Point{p.X, p.Y + 1})
	}

	reallyPossible := []Point{}
	for _, possiblePoint := range possible {
		originalHeight := m.Elevations[p.Y][p.X]
		possibleHeight := m.Elevations[possiblePoint.Y][possiblePoint.X]
		if originalHeight >= possibleHeight || possibleHeight-1 == originalHeight {
			reallyPossible = append(reallyPossible, possiblePoint)
		}
	}
	return reallyPossible
}

func (m ElevationMap) FindShortestPath(from, to Point) (int, error) {
	vis := map[Point]bool{}

	type node struct {
		P        Point
		Distance int
	}
	nextNodes := []node{{from, 0}}
	for len(nextNodes) > 0 {
		next := nextNodes[0]
		nextNodes = nextNodes[1:]
		if vis[next.P] {
			continue
		}

		vis[next.P] = true
		if next.P == to {
			return next.Distance, nil
		}

		neighbors := m.possibleMoves(next.P)
		for _, n := range neighbors {
			if !vis[n] {
				nextNodes = append(nextNodes, node{P: n, Distance: next.Distance + 1})
			}
		}
	}
	return 0, fmt.Errorf("no path to target")
}

func (m ElevationMap) FindAllPossibleStarts() []Point {
	possibleStarts := []Point{}
	for y := 0; y < len(m.Elevations); y++ {
		for x := 0; x < len(m.Elevations[y]); x++ {
			if m.Elevations[y][x] == 0 {
				possibleStarts = append(possibleStarts, Point{x, y})
			}
		}
	}
	return possibleStarts
}

func Part1(input []byte) (runner.Solution, error) {
	emap := Parse(input)
	steps, err := emap.FindShortestPath(emap.Start, emap.Target)
	return runner.Solution{Message: fmt.Sprintf("Least steps to Target: %v", steps)}, err
}

func Part2(input []byte) (runner.Solution, error) {
	emap := Parse(input)
	allStartingPoints := emap.FindAllPossibleStarts()

	min := math.MaxInt
	for _, s := range allStartingPoints {
		localMin, err := emap.FindShortestPath(s, emap.Target)
		if err != nil {
			continue
		}
		if localMin < min {
			min = localMin
		}
	}
	return runner.Solution{Message: fmt.Sprintf("Least steps from any starting a: %v", min)}, nil
}

func Parse(input []byte) ElevationMap {
	emap := ElevationMap{}
	emap.Elevations = [][]int{}
	for y, l := range strings.Split(string(input), "\n") {
		mapLine := []int{}
		for x, char := range l {
			var value int
			switch char {
			case 'S':
				value = 0
				emap.Start.X = x
				emap.Start.Y = y
			case 'E':
				value = int('z' - 'a')
				emap.Target.X = x
				emap.Target.Y = y
			default:
				value = int(char - 'a')
			}
			mapLine = append(mapLine, value)
		}
		emap.Elevations = append(emap.Elevations, mapLine)
	}
	return emap
}
