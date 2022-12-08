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

func (f Forest) IsVisible(x, y int) bool {
	if x == 0 || x == len(f.heights[y])-1 || y == 0 || y == len(f.heights[x])-1 {
		// on the edge
		return true
	}
	treeHeight := f.heights[y][x]
	// check north
	isVisibleNorth := true
	for yIter := y - 1; yIter >= 0; yIter-- {
		otherTree := f.heights[yIter][x]
		if otherTree >= treeHeight {
			isVisibleNorth = false
			break
		}
	}
	// check east
	isVisibleEast := true
	for xIter := x + 1; xIter < len(f.heights[y]); xIter++ {
		otherTree := f.heights[y][xIter]
		if otherTree >= treeHeight {
			isVisibleEast = false
			break
		}
	}
	// check south
	isVisibleSouth := true
	for yIter := y + 1; yIter < len(f.heights); yIter++ {
		otherTree := f.heights[yIter][x]
		if otherTree >= treeHeight {
			isVisibleSouth = false
			break
		}
	}
	// check west
	isVisibleWest := true
	for xIter := x - 1; xIter >= 0; xIter-- {
		otherTree := f.heights[y][xIter]
		if otherTree >= treeHeight {
			isVisibleWest = false
			break
		}
	}

	return isVisibleNorth || isVisibleEast || isVisibleSouth || isVisibleWest
}

func (f Forest) ScenicScore(x, y int) int {
	if x == 0 || x == len(f.heights[y])-1 || y == 0 || y == len(f.heights[x])-1 {
		// on the edge
		return 0
	}
	treeHeight := f.heights[y][x]
	// check north
	var northDistance int
	var foundNorthTree bool
	for yIter := y - 1; yIter >= 0; yIter-- {
		otherTree := f.heights[yIter][x]
		if otherTree >= treeHeight {
			northDistance = y - yIter
			foundNorthTree = true
			break
		}
	}
	if !foundNorthTree {
		northDistance = y
	}
	// check east
	var eastDistance int
	var foundEastTree bool
	for xIter := x + 1; xIter < len(f.heights[y]); xIter++ {
		otherTree := f.heights[y][xIter]
		if otherTree >= treeHeight {
			eastDistance = xIter - x
			foundEastTree = true
			break
		}
	}
	if !foundEastTree {
		eastDistance = len(f.heights[y]) - x - 1
	}
	// check south
	var southDistance int
	var foundSouthTree bool
	for yIter := y + 1; yIter < len(f.heights); yIter++ {
		otherTree := f.heights[yIter][x]
		if otherTree >= treeHeight {
			southDistance = yIter - y
			foundSouthTree = true
			break
		}
	}
	if !foundSouthTree {
		southDistance = len(f.heights) - y - 1
	}
	// check west
	var westDistance int
	var foundWestTree bool
	for xIter := x - 1; xIter >= 0; xIter-- {
		otherTree := f.heights[y][xIter]
		if otherTree >= treeHeight {
			westDistance = x - xIter
			foundWestTree = true
			break
		}
	}
	if !foundWestTree {
		westDistance = x
	}

	return northDistance * eastDistance * southDistance * westDistance
}

func Part1(input []byte) (runner.Solution, error) {
	p := Parse(input)
	xSize, ySize := p.Size()

	visibleTrees := 0
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			if p.IsVisible(x, y) {
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
			score := p.ScenicScore(x, y)
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
