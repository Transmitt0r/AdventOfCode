package day15_test

import (
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day15"
)

var sampleInput = []byte(`Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3`)

func TestPart1(t *testing.T) {
	sol, err := day15.Part1(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "How many positions cannot contain a beacon: 26" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestPart2(t *testing.T) {
	sol, err := day15.Part2(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Tuning frequency: 56000011" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestDistanceToBeacon(t *testing.T) {
	s := day15.Sensor{day15.Point{5, -6}, &day15.Beacon{-15, 0}}
	if s.DistanceToBeacon() != 26 {
		t.Error("Wrong distance calculated!")
	}
}

func TestParse(t *testing.T) {
	sensors := day15.Parse(sampleInput)
	if len(sensors) != 14 {
		t.Error("Parsing error")
	}
}
