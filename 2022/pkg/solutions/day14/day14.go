package day14

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

type Material rune

const (
	Rock        Material = '#'
	Air         Material = '.'
	Sand        Material = '+'
	SettledSand Material = 'o'
	OutOfBounds Material = 'x'
)

var (
	ErrorOutOfBounds error = errors.New("out of bounds")
)

type Point struct {
	X int
	Y int
}

type Instruction []Point

type Plot struct {
	Offset Point
	plot   [][]Material
}

func (p Plot) Render() string {
	out := ""
	for _, y := range p.plot {
		for _, x := range y {
			out += string(x)
		}
		out += "\n"
	}
	return out
}

func (p Plot) getMaterialAt(coord Point) Material {
	if coord.Y < p.Offset.Y || coord.Y >= len(p.plot)+p.Offset.Y || coord.X < p.Offset.X || coord.X >= len(p.plot[coord.Y-p.Offset.Y])+p.Offset.X {
		return OutOfBounds
	}
	return p.plot[coord.Y-p.Offset.Y][coord.X-p.Offset.X]
}

func (p *Plot) move(sand *Point) (bool, error) {
	sandDown := Point{sand.X, sand.Y + 1}
	sandLeft := Point{sand.X - 1, sand.Y + 1}
	sandRight := Point{sand.X + 1, sand.Y + 1}

	switch p.getMaterialAt(sandDown) {
	case OutOfBounds:
		return false, ErrorOutOfBounds
	case Air:
		sand.Y++
		return true, nil
	}

	switch p.getMaterialAt(sandLeft) {
	case OutOfBounds:
		return false, ErrorOutOfBounds
	case Air:
		sand.Y++
		sand.X--
		return true, nil
	}

	switch p.getMaterialAt(sandRight) {
	case OutOfBounds:
		return false, ErrorOutOfBounds
	case Air:
		sand.Y++
		sand.X++
		return true, nil
	}

	return false, nil
}

func (p *Plot) AddSand(sand Point) {
	p.plot[sand.Y][sand.X-p.Offset.X] = SettledSand
}

func (p *Plot) Simulate() int {
	sandEntryPoint := Point{500, 0}
	sandCounter := 0

topLoop:
	for {
		sandCorn := sandEntryPoint
		canMove := true
		var err error
		for canMove {
			canMove, err = p.move(&sandCorn)
			if err != nil {
				break topLoop
			}
		}
		p.AddSand(sandCorn)
		sandCounter++
		if sandCorn == sandEntryPoint {
			break
		}
	}

	return sandCounter
}

func Part1(input []byte) (runner.Solution, error) {
	insts := Parse(input)
	min, max := FindLimits(insts)
	plot := NewPlot(min, max, 0)
	for _, inst := range insts {
		plot.Draw(inst)
	}
	plot.Render()
	res := plot.Simulate()
	return runner.Solution{Message: fmt.Sprintf("Units of sand until full: %v", res)}, nil
}
func Part2(input []byte) (runner.Solution, error) {
	insts := Parse(input)
	min, max := FindLimits(insts)
	min.X = 0
	max.X = max.X * 2
	plot := NewPlot(min, max, 2)
	for _, inst := range insts {
		plot.Draw(inst)
	}
	plot.Render()
	res := plot.Simulate()
	return runner.Solution{Message: fmt.Sprintf("Units of sand until hole is plugged: %v", res)}, nil
}

func Parse(input []byte) []Instruction {
	results := []Instruction{}
	for _, l := range strings.Split(string(input), "\n") {
		points := strings.Split(l, " -> ")

		inst := Instruction{}
		for i := 0; i < len(points); i++ {
			p := strings.Split(points[i], ",")
			pX, _ := strconv.Atoi(p[0])
			pY, _ := strconv.Atoi(p[1])
			inst = append(inst, Point{pX, pY})

		}
		results = append(results, inst)
	}
	return results
}

func FindLimits(insts []Instruction) (Point, Point) {
	max := Point{0, 0}
	min := Point{math.MaxInt, math.MaxInt}
	for _, i := range insts {
		for _, p := range i {
			if p.X > max.X {
				max.X = p.X
			}
			if p.Y > max.Y {
				max.Y = p.Y
			}
			if p.X < min.X {
				min.X = p.X
			}
			if p.Y < min.Y {
				min.Y = p.Y
			}
		}
	}
	return min, max
}

func NewPlot(min, max Point, floorOffset int) Plot {
	newPlot := Plot{}
	newPlot.Offset = min
	newPlot.Offset.Y = 0
	newPlot.plot = make([][]Material, max.Y+1+floorOffset)
	for y := 0; y < len(newPlot.plot); y++ {
		newPlot.plot[y] = make([]Material, max.X-newPlot.Offset.X+1)
		for x := 0; x < len(newPlot.plot[y]); x++ {
			mat := Air
			if floorOffset > 0 && y == len(newPlot.plot)-1 {
				mat = Rock
			}
			newPlot.plot[y][x] = mat
		}
	}
	return newPlot
}

func (p *Plot) Draw(inst Instruction) {
	for i := 0; i < len(inst)-1; i++ {
		start := inst[i]
		end := inst[i+1]

		if start.X != end.X {
			// draw x
			var startX int
			var endX int
			if start.X < end.X {
				startX = start.X
				endX = end.X
			} else {
				startX = end.X
				endX = start.X
			}
			for i := startX; i <= endX; i++ {
				p.plot[start.Y-p.Offset.Y][i-p.Offset.X] = Rock
			}
		}
		if start.Y != end.Y {
			// draw y
			var startY int
			var endY int
			if start.Y < end.Y {
				startY = start.Y
				endY = end.Y
			} else {
				startY = end.Y
				endY = start.Y
			}
			for i := startY; i <= endY; i++ {
				p.plot[i-p.Offset.Y][start.X-p.Offset.X] = Rock
			}
		}
	}
}
