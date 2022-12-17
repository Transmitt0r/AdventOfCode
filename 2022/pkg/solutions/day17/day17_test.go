package day17_test

import (
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day17"
)

var sampleInput = []byte(`>>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>`)

func TestPart1(t *testing.T) {
	sol, err := day17.Part1(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Height of the tower: 3068" {
		t.Errorf("Wrong solution %v", sol)
	}
}
