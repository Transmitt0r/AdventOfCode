package day10

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

const (
	EmptyPixel = "."
	DrawnPixel = "#"
)

type Processor struct {
	registerX             int
	registerValuesByCycle []int
	clockCycle            int
	crt                   [][]string
}

func NewProcessor() *Processor {
	crt := make([][]string, 6)
	for i := 0; i < 6; i++ {
		crt[i] = make([]string, 40)
		for j := 0; j < 40; j++ {
			crt[i][j] = EmptyPixel
		}
	}
	p := &Processor{
		registerX:             1,
		registerValuesByCycle: make([]int, 500),
		crt:                   crt,
	}
	return p
}

func (p *Processor) Noop() {
	p.IncreaseClock()
}

func (p *Processor) AddX(num int) {
	p.IncreaseClock()
	p.IncreaseClock()
	p.registerX += num
}

func (p *Processor) IncreaseClock() {
	p.registerValuesByCycle[p.clockCycle] = p.registerX
	p.UpdateCRT()
	p.clockCycle++
}

func (p *Processor) UpdateCRT() {
	horizontalStart := p.registerX - 1
	horizontalEnd := p.registerX + 1

	verticalPosition := p.clockCycle / 40
	horizontalPosition := p.clockCycle % 40

	if horizontalPosition == horizontalStart {
		p.crt[verticalPosition][horizontalPosition] = DrawnPixel
	}
	if horizontalPosition == horizontalEnd {
		p.crt[verticalPosition][horizontalPosition] = DrawnPixel
	}
	if horizontalPosition == p.registerX {
		p.crt[verticalPosition][horizontalPosition] = DrawnPixel
	}
}

func (p Processor) DrawCRT() string {
	lines := []string{}
	for _, line := range p.crt {
		lines = append(lines, strings.Join(line, ""))
	}
	return strings.Join(lines, "\n")
}

func Part1(input []byte) (runner.Solution, error) {
	p := NewProcessor()
	for _, l := range strings.Split(string(input), "\n") {
		inst := strings.Split(l, " ")
		if inst[0] == "noop" {
			p.Noop()
		} else {
			val, _ := strconv.Atoi(inst[1])
			p.AddX(val)
		}
	}
	p.IncreaseClock()

	signalSum := 0
	for i := 19; i < 220; i += 40 {
		signalSum += (i + 1) * p.registerValuesByCycle[i]
	}
	return runner.Solution{Message: fmt.Sprintf("Summed up signal strengh: %v", signalSum)}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	p := NewProcessor()
	for _, l := range strings.Split(string(input), "\n") {
		inst := strings.Split(l, " ")
		if inst[0] == "noop" {
			p.Noop()
		} else {
			val, _ := strconv.Atoi(inst[1])
			p.AddX(val)
		}
	}
	p.IncreaseClock()
	return runner.Solution{Message: fmt.Sprintf("Solution: \n%v", p.DrawCRT())}, nil
}
