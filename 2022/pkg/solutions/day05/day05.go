package day05

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

type Instruction struct {
	Count int
	From  int
	To    int
}

type CrateStack struct {
	crates [][]rune
}

func (c *CrateStack) ApplyMultiple(inst Instruction) {
	c.crates[inst.To-1] = append(c.crates[inst.To-1], c.crates[inst.From-1][len(c.crates[inst.From-1])-inst.Count:]...)
	c.crates[inst.From-1] = c.crates[inst.From-1][:len(c.crates[inst.From-1])-inst.Count]
}

func (c *CrateStack) Apply(inst Instruction) {
	for i := 0; i < inst.Count; i++ {
		c.crates[inst.To-1] = append(c.crates[inst.To-1], c.crates[inst.From-1][len(c.crates[inst.From-1])-1])
		c.crates[inst.From-1] = c.crates[inst.From-1][:len(c.crates[inst.From-1])-1]
	}
}

func (c CrateStack) TopCrates() string {
	out := ""
	for _, crate := range c.crates {
		if len(crate) == 0 {
			out += " "
			continue
		}
		out += string(crate[len(crate)-1])
	}
	return out
}

func (c CrateStack) Len() int {
	return len(c.crates)
}

func Part1(input []byte) (runner.Solution, error) {
	stack, instructions := Parse(input)

	for _, inst := range instructions {
		stack.Apply(inst)
	}

	return runner.Solution{Message: fmt.Sprintf("Top Crates: %v", stack.TopCrates())}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	stack, instructions := Parse(input)

	for _, inst := range instructions {
		stack.ApplyMultiple(inst)
	}

	return runner.Solution{Message: fmt.Sprintf("Top Crates: %v", stack.TopCrates())}, nil
}

func ParseStart(lines []string) CrateStack {
	columns := len(lines[0]) / 3
	out := make([][]rune, columns)
	for _, line := range lines {
		for i := 1; i < len(line); i += 4 {
			index := i / 4
			char := line[i]
			if string(char) == " " {
				continue
			}
			if unicode.IsNumber(rune(char)) {
				break
			}
			out[index] = append([]rune{rune(char)}, out[index]...)
		}
	}
	return CrateStack{out}
}

func ParseEnd(lines []string) []Instruction {
	out := []Instruction{}
	re := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)

	for _, line := range lines {
		res := re.FindAllStringSubmatch(line, -1)[0]
		count, _ := strconv.Atoi(res[1])
		from, _ := strconv.Atoi(res[2])
		to, _ := strconv.Atoi(res[3])
		out = append(out, Instruction{Count: count, From: from, To: to})
	}
	return out
}

func Parse(input []byte) (CrateStack, []Instruction) {
	lines := strings.Split(string(input), "\n")
	var seperatorIndex int
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			seperatorIndex = i
			break
		}
	}
	return ParseStart(lines[:seperatorIndex]), ParseEnd(lines[1+seperatorIndex:])
}
