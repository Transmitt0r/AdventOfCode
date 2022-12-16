package day15

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

type Point struct {
	X int
	Y int
}

type Sensor struct {
	Point
	Closest *Beacon
}

func (s Sensor) DistanceToBeacon() int {
	return calculateManhattanDistance(s.Point, Point(*s.Closest))
}

func calculateManhattanDistance(p1, p2 Point) int {
	distance := 0
	if p1.X > p2.X {
		distance += p1.X - p2.X
	} else {
		distance += p2.X - p1.X
	}

	if p1.Y > p2.Y {
		distance += p1.Y - p2.Y
	} else {
		distance += p2.Y - p1.Y
	}

	return distance
}

func (s Sensor) xWidthAtY(y int) (int, int) {
	absDeltaY := y - s.Y
	if absDeltaY < 0 {
		absDeltaY = -absDeltaY
	}

	return s.X + absDeltaY - s.DistanceToBeacon(), s.X - absDeltaY + s.DistanceToBeacon()
}

func (s Sensor) pointIsOutOfRange(p Point) bool {
	d := calculateManhattanDistance(s.Point, p)
	return d > s.DistanceToBeacon()
}

func pointIsOutOfRangeForAllSensors(sensors []Sensor, p Point) bool {
	for _, s := range sensors {
		if !s.pointIsOutOfRange(p) {
			return false
		}
	}
	return true
}

type Beacon Point

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

func Part1(input []byte) (runner.Solution, error) {
	sensors := Parse(input)

	line := 2000000
	if len(sensors) == 14 {
		// hacky to detect test
		line = 10
	}

	max := math.MinInt
	min := math.MaxInt
	for _, s := range sensors {
		sMin, sMax := s.xWidthAtY(line)

		if sMax > max {
			max = sMax
		}
		if sMin < min {
			min = sMin
		}
	}
	return runner.Solution{Message: fmt.Sprintf("How many positions cannot contain a beacon: %v", max-min)}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	sensors := Parse(input)
	limit := 4000000
	if len(sensors) == 14 {
		// hacky to detect test
		limit = 20
	}
	foundPoint := Point{}
	for _, s1 := range sensors {
		for _, s2 := range sensors {
			s1dist := s1.X - s1.Y - s1.DistanceToBeacon()

			s2dist := s2.X + s2.Y + s2.DistanceToBeacon()

			x := (s2dist + s1dist) / 2 // midpoint?
			y := (s2dist-s1dist)/2 + 1 // wtf?

			potentialPoint := Point{x, y}
			if 0 < x && x < limit && 0 < y && y < limit && pointIsOutOfRangeForAllSensors(sensors, potentialPoint) {
				foundPoint = potentialPoint
				break
			}
		}
	}
	return runner.Solution{Message: fmt.Sprintf("Tuning frequency: %v", calculateTuningFrequency(foundPoint))}, nil
}

func calculateTuningFrequency(p Point) int {
	return p.X*4000000 + p.Y
}

func Parse(input []byte) []Sensor {
	sensors := []Sensor{}
	re := regexp.MustCompile(`^Sensor at (x=(-?\d+), y=(-?\d+)): closest beacon is at (x=(-?\d+), y=(-?\d+))$`)
	for _, l := range strings.Split(string(input), "\n") {
		res := re.FindAllStringSubmatch(l, -1)[0]
		sX, _ := strconv.Atoi(res[2])
		sY, _ := strconv.Atoi(res[3])
		bX, _ := strconv.Atoi(res[5])
		bY, _ := strconv.Atoi(res[6])
		beacon := Beacon{bX, bY}
		sensors = append(sensors, Sensor{Point{sX, sY}, &beacon})
	}
	return sensors
}
