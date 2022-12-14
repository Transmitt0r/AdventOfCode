package day14_test

import (
	"fmt"
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day14"
)

var sampleInput = []byte(`498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`)

func TestPart1(t *testing.T) {
	sol, err := day14.Part1(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Units of sand until full: 24" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestPart2(t *testing.T) {
	sol, err := day14.Part2(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Units of sand until hole is plugged: 93" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestParse(t *testing.T) {
	insts := day14.Parse(sampleInput)
	min, max := day14.FindLimits(insts)
	minExp := day14.Point{494, 4}
	maxExp := day14.Point{503, 9}
	if min != minExp || max != maxExp {
		t.Error("Unexpected Limits")
	}
	plot := day14.NewPlot(min, max, 0)
	for _, inst := range insts {
		plot.Draw(inst)
	}

	out := plot.Render()
	outExp := `..........
..........
..........
..........
....#...##
....#...#.
..###...#.
........#.
........#.
#########.
`
	fmt.Print(out)
	if out != outExp {
		t.Error("Rendering failed")
	}
}
