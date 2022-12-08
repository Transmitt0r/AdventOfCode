package day08

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

type Forest struct {
	heights [][]int
}

func (f Forest) Size() (int, int) {
	return len(f.heights[0]), len(f.heights)
}

func (f Forest) IsVisibleWithScore(x, y int) (bool, int) {
	if x == 0 || x == len(f.heights[y])-1 || y == 0 || y == len(f.heights[x])-1 {
		// on the edge
		return true, 0
	}
	treeHeight := f.heights[y][x]
	// check north
	var northDistance int
	isVisibleNorth := true
	for yIter := y - 1; yIter >= 0; yIter-- {
		otherTree := f.heights[yIter][x]
		if otherTree >= treeHeight {
			northDistance = y - yIter
			isVisibleNorth = false
			break
		}
	}
	if isVisibleNorth {
		northDistance = y
	}
	// check east
	var eastDistance int
	isVisibleEast := true
	for xIter := x + 1; xIter < len(f.heights[y]); xIter++ {
		otherTree := f.heights[y][xIter]
		if otherTree >= treeHeight {
			eastDistance = xIter - x
			isVisibleEast = false
			break
		}
	}
	if isVisibleEast {
		eastDistance = len(f.heights[y]) - x - 1
	}
	// check south
	var southDistance int
	isVisibleSouth := true
	for yIter := y + 1; yIter < len(f.heights); yIter++ {
		otherTree := f.heights[yIter][x]
		if otherTree >= treeHeight {
			southDistance = yIter - y
			isVisibleSouth = false
			break
		}
	}
	if isVisibleSouth {
		southDistance = len(f.heights) - y - 1
	}
	// check west
	var westDistance int
	isVisibleWest := true
	for xIter := x - 1; xIter >= 0; xIter-- {
		otherTree := f.heights[y][xIter]
		if otherTree >= treeHeight {
			westDistance = x - xIter
			isVisibleWest = false
			break
		}
	}
	if isVisibleWest {
		westDistance = x
	}

	scenicScore := northDistance * eastDistance * southDistance * westDistance
	isVisible := isVisibleNorth || isVisibleEast || isVisibleSouth || isVisibleWest
	return isVisible, scenicScore
}

func Part1(input []byte) (runner.Solution, error) {
	p := Parse(input)
	xSize, ySize := p.Size()

	visibleTrees := 0
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			isVisible, _ := p.IsVisibleWithScore(x, y)
			if isVisible {
				visibleTrees++
			}
		}
	}

	return runner.Solution{Message: fmt.Sprintf("Number of visible trees: %v", visibleTrees)}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	p := Parse(input)
	xSize, ySize := p.Size()

	maxScenicScore := 0
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			_, score := p.IsVisibleWithScore(x, y)
			if score > maxScenicScore {
				maxScenicScore = score
			}
		}
	}
	return runner.Solution{Message: fmt.Sprintf("Highest scenic score: %v", maxScenicScore)}, nil
}

func Parse(input []byte) Forest {
	lines := strings.Split(string(input), "\n")
	intLines := [][]int{}

	for _, line := range lines {
		intL := make([]int, len(line))
		for i, char := range line {
			asInt, _ := strconv.Atoi(string(char))
			intL[i] = asInt
		}
		intLines = append(intLines, intL)
	}
	return Forest{intLines}
}
