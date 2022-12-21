package day18_test

import (
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day18"
)

var sampleInput = []byte(`2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`)

func TestPart1(t *testing.T) {
	sol, err := day18.Part1(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Surface Area of Droplet: 64" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestPart2(t *testing.T) {
	sol, err := day18.Part2(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Hull Area of Droplet: 58" {
		t.Errorf("Wrong solution %v", sol)
	}
}
