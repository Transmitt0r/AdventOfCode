package day18

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

type Cube struct {
	X, Y, Z int
}

func (c Cube) neighbors() []Cube {
	return []Cube{
		{c.X - 1, c.Y, c.Z},
		{c.X + 1, c.Y, c.Z},
		{c.X, c.Y - 1, c.Z},
		{c.X, c.Y + 1, c.Z},
		{c.X, c.Y, c.Z - 1},
		{c.X, c.Y, c.Z + 1},
	}
}

type Droplet struct {
	Cubes map[Cube]bool
}

func (d Droplet) surfaceArea() int {
	surfaceArea := 0
	for c := range d.Cubes {
		for _, n := range c.neighbors() {
			if d.Cubes[n] {
				continue
			}
			surfaceArea++
		}
	}
	return surfaceArea
}

func (d Droplet) hullArea() int {
	minLim, maxLim := d.limits()
	minLim.X--
	minLim.Y--
	minLim.Z--
	maxLim.X++
	maxLim.Y++
	maxLim.Z++

	max := d.maxX()
	max.X++

	visited := map[Cube]bool{}
	queue := []Cube{max}

	surfaceArea := 0

	for len(queue) > 0 {
		var next Cube
		next, queue = queue[len(queue)-1], queue[:len(queue)-1]
		if d.Cubes[next] || visited[next] {
			continue
		}
		visited[next] = true

		if next.X < minLim.X || next.Y < minLim.Y || next.Z < minLim.Z || next.X > maxLim.X || next.Y > maxLim.Y || next.Z > maxLim.Z {
			continue
		}
		surfaceArea += d.howManySurfacesDoesCubeTouch(next)
		queue = append(queue, next.neighbors()...)
	}
	return surfaceArea
}

func (d Droplet) howManySurfacesDoesCubeTouch(c Cube) int {
	count := 0
	for _, n := range c.neighbors() {
		if d.Cubes[n] {
			count++
		}
	}
	return count
}

func (d Droplet) limits() (Cube, Cube) {
	min := Cube{math.MaxInt, math.MaxInt, math.MaxInt}
	max := Cube{math.MinInt, math.MinInt, math.MinInt}

	for c := range d.Cubes {
		if c.X < min.X {
			min.X = c.X
		}
		if c.Y < min.Y {
			min.Y = c.Y
		}
		if c.Z < min.Z {
			min.Z = c.Z
		}

		if c.X > max.X {
			max.X = c.X
		}
		if c.Y > max.Y {
			max.Y = c.Y
		}
		if c.Z > max.Z {
			max.Z = c.Z
		}
	}
	return min, max
}

func (d Droplet) maxX() Cube {
	max := Cube{X: math.MinInt}
	for c := range d.Cubes {
		if c.X > max.X {
			max = c
		}
	}
	return max
}

func newDroplet(cubes []Cube) Droplet {
	newDrop := Droplet{make(map[Cube]bool)}
	for _, c := range cubes {
		newDrop.Cubes[c] = true
	}
	return newDrop
}

func ParseCubes(input []byte) []Cube {
	cubes := []Cube{}
	for _, l := range strings.Split(string(input), "\n") {
		vals := strings.Split(l, ",")
		X, _ := strconv.Atoi(vals[0])
		Y, _ := strconv.Atoi(vals[1])
		Z, _ := strconv.Atoi(vals[2])
		cubes = append(cubes, Cube{X, Y, Z})
	}
	return cubes
}

func Part1(input []byte) (runner.Solution, error) {
	drop := newDroplet(ParseCubes(input))
	return runner.Solution{Message: fmt.Sprintf("Surface Area of Droplet: %v", drop.surfaceArea())}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	drop := newDroplet(ParseCubes(input))
	return runner.Solution{Message: fmt.Sprintf("Hull Area of Droplet: %v", drop.hullArea())}, nil
}
