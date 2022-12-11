package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

type Monkey struct {
	Items       []int
	Operation   func(int) int
	TestDivisor int
	IfTrue      int
	IfFalse     int
}

func (m Monkey) Test(item int) bool {
	return (item % m.TestDivisor) == 0
}

func (m *Monkey) PopItem() (int, error) {
	if len(m.Items) <= 0 {
		return 0, fmt.Errorf("no items available")
	}
	var firstItem int
	firstItem, m.Items = m.Items[0], m.Items[1:]
	return firstItem, nil
}

func (m *Monkey) PushItem(item int) {
	m.Items = append(m.Items, item)
}

func (m Monkey) InspectItem(item int) int {
	newInt := m.Operation(item)
	if newInt < item {
		panic("Overflow!")
	}
	return newInt
}

func (m Monkey) CheckWorryLevel(item int) int {
	if m.Test(item) {
		return m.IfTrue
	} else {
		return m.IfFalse
	}
}

func SimulateMonkeys(monkeys []Monkey, iterations int, reductionFunc func(int) int) int {
	inspectedItems := make([]int, len(monkeys))
	for i := 0; i < iterations; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := &monkeys[j]
			for {
				item, err := monkey.PopItem()
				if err != nil {
					break
				}
				item = monkey.InspectItem(item)
				inspectedItems[j]++
				item = reductionFunc(item)
				throwTo := monkey.CheckWorryLevel(item)
				monkeys[throwTo].PushItem(item)
			}
		}
	}
	sort.Ints(inspectedItems)
	return inspectedItems[len(monkeys)-1] * inspectedItems[len(monkeys)-2]
}

func findDivisor(monkeys []Monkey) int {
	divisor := 1
	for _, m := range monkeys {
		divisor *= m.TestDivisor
	}
	return divisor
}

func Part1(input []byte) (runner.Solution, error) {
	monkeys, err := Parse(input)
	return runner.Solution{Message: fmt.Sprintf("Monkey Business: %v", SimulateMonkeys(monkeys, 20, func(i int) int { return i / 3 }))}, err
}

func Part2(input []byte) (runner.Solution, error) {
	monkeys, err := Parse(input)
	multiple := findDivisor(monkeys)
	return runner.Solution{Message: fmt.Sprintf("Monkey Business: %v", SimulateMonkeys(monkeys, 10000, func(i int) int { return i % multiple }))}, err
}

func Parse(input []byte) ([]Monkey, error) {
	monkeys := []Monkey{}
	prefixStart := "Monkey "
	prefixStartingItems := "  Starting items: "
	prefixOperation := "  Operation: new = old "
	prefixTest := "  Test: divisible by "
	prefixIfTrue := "    If true: throw to monkey "
	prefixIfFalse := "    If false: throw to monkey "

	var curMonkey Monkey
	for _, l := range strings.Split(string(input), "\n")[1:] {
		switch {
		case strings.HasPrefix(l, prefixStart):
			monkeys = append(monkeys, curMonkey)
			curMonkey = Monkey{}
		case strings.HasPrefix(l, prefixStartingItems):
			items := strings.Split(l[len(prefixStartingItems):], ", ")
			for _, i := range items {
				asInt, err := strconv.Atoi(i)
				if err != nil {
					return monkeys, nil
				}
				curMonkey.Items = append(curMonkey.Items, asInt)
			}
		case strings.HasPrefix(l, prefixOperation):
			wihoutPrefix := l[len(prefixOperation):]
			noNum := false
			var num int
			if strings.HasSuffix(wihoutPrefix, "old") {
				noNum = true
			} else {
				var err error
				num, err = strconv.Atoi(wihoutPrefix[2:])
				if err != nil {
					return monkeys, err
				}
			}
			switch wihoutPrefix[0] {
			case '+':
				if noNum {
					curMonkey.Operation = func(i int) int {
						return i + i
					}
				} else {
					curMonkey.Operation = func(i int) int {
						return i + num
					}
				}
			case '*':
				if noNum {
					curMonkey.Operation = func(i int) int {
						return i * i
					}
				} else {
					curMonkey.Operation = func(i int) int {
						return i * num
					}
				}
			}
		case strings.HasPrefix(l, prefixTest):
			divisor, err := strconv.Atoi(l[len(prefixTest):])
			if err != nil {
				return monkeys, err
			}
			curMonkey.TestDivisor = divisor
		case strings.HasPrefix(l, prefixIfTrue):
			trueThrow, err := strconv.Atoi(l[len(prefixIfTrue):])
			if err != nil {
				return monkeys, err
			}
			curMonkey.IfTrue = trueThrow
		case strings.HasPrefix(l, prefixIfFalse):
			falseThrow, err := strconv.Atoi(l[len(prefixIfFalse):])
			if err != nil {
				return monkeys, err
			}
			curMonkey.IfFalse = falseThrow
		}
	}
	monkeys = append(monkeys, curMonkey)
	return monkeys, nil
}
