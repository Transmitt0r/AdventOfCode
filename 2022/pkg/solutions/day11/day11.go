package day11

import (
	"fmt"
	"sort"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

var inputMonkeys = []Monkey{
	{
		Items:       []int{62, 92, 50, 63, 62, 93, 73, 50},
		Operation:   func(i int) int { return i * 7 },
		TestDivisor: 2,
		IfTrue:      7,
		IfFalse:     1,
	},
	{
		Items:       []int{51, 97, 74, 84, 99},
		Operation:   func(i int) int { return i + 3 },
		TestDivisor: 7,
		IfTrue:      2,
		IfFalse:     4,
	},
	{
		Items:       []int{98, 86, 62, 76, 51, 81, 95},
		Operation:   func(i int) int { return i + 4 },
		TestDivisor: 13,
		IfTrue:      5,
		IfFalse:     4,
	},
	{
		Items:       []int{53, 95, 50, 85, 83, 72},
		Operation:   func(i int) int { return i + 5 },
		TestDivisor: 19,
		IfTrue:      6,
		IfFalse:     0,
	},
	{
		Items:       []int{59, 60, 63, 71},
		Operation:   func(i int) int { return i * 5 },
		TestDivisor: 11,
		IfTrue:      5,
		IfFalse:     3,
	},
	{
		Items:       []int{92, 65},
		Operation:   func(i int) int { return i * i },
		TestDivisor: 5,
		IfTrue:      6,
		IfFalse:     3,
	},
	{
		Items:       []int{78},
		Operation:   func(i int) int { return i + 8 },
		TestDivisor: 3,
		IfTrue:      0,
		IfFalse:     7,
	},
	{
		Items:       []int{84, 93, 54},
		Operation:   func(i int) int { return i + 1 },
		TestDivisor: 17,
		IfTrue:      2,
		IfFalse:     1,
	},
}

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

func Part1Monkey(monkeys []Monkey) int {
	inspectedItems := make([]int, len(monkeys))
	for i := 0; i < 20; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := &monkeys[j]
			for {
				item, err := monkey.PopItem()
				if err != nil {
					break
				}
				item = monkey.InspectItem(item)
				inspectedItems[j]++
				item = item / 3 // reduce worry level
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

func Part2Monkey(monkeys []Monkey) int {
	inspectedItems := make([]int, len(monkeys))
	lcm := findDivisor(monkeys)
	for i := 0; i < 10000; i++ {
		for j := 0; j < len(monkeys); j++ {
			monkey := &monkeys[j]
			for {
				item, err := monkey.PopItem()
				if err != nil {
					break
				}
				item = monkey.InspectItem(item)
				inspectedItems[j]++
				throwTo := monkey.CheckWorryLevel(item)
				item = item % lcm
				monkeys[throwTo].PushItem(item)
			}
		}
	}
	sort.Ints(inspectedItems)
	return inspectedItems[len(monkeys)-1] * inspectedItems[len(monkeys)-2]
}

func Part1(input []byte) (runner.Solution, error) {
	return runner.Solution{Message: fmt.Sprintf("Monkey Business: %v", Part1Monkey(inputMonkeys))}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	return runner.Solution{Message: fmt.Sprintf("Monkey Business: %v", Part2Monkey(inputMonkeys))}, nil
}
