package day13

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

type Result int

const (
	IsEqual     Result = iota
	IsOrdered   Result = iota
	IsUnordered Result = iota
)

func equal(left, right any) Result {
	// returns  0 when equal, <= when left packet comes before right packet
	l, okL := left.(float64)
	r, okR := right.(float64)

	if okL && okR {
		if l-r == 0 {
			return IsEqual
		}
		if l-r < 0 {
			return IsOrdered
		}
		return IsUnordered
	}

	var lList []any
	var rList []any

	switch left.(type) {
	case []any, []float64:
		lList = left.([]any)
	case float64:
		lList = []any{left}
	}

	switch right.(type) {
	case []any, []float64:
		rList = right.([]any)
	case float64:
		rList = []any{right}
	}

	for i := range lList {
		if len(rList) <= i {
			return IsUnordered
		}
		if r := equal(lList[i], rList[i]); r != IsEqual {
			return r
		}
	}
	if len(lList) == len(rList) {
		return IsEqual
	}
	return IsOrdered
}

func Part1(input []byte) (runner.Solution, error) {
	packets := Parse(input)
	sum := 0
	for i := 0; i < len(packets)-1; i += 2 {
		res := equal(packets[i], packets[i+1])
		if res == IsOrdered || res == IsEqual {
			fmt.Printf("In order: %v\n", i/2+1)
			sum += i/2 + 1
		} else {
			fmt.Printf("Not in order: %v\n", i/2+1)
		}
	}
	return runner.Solution{Message: fmt.Sprintf("Sum of indices of packets in right order: %v", sum)}, nil
}

func generateDividerPackets() []any {
	return []any{
		[]any{[]any{float64(2)}},
		[]any{[]any{float64(6)}},
	}
}

func Part2(input []byte) (runner.Solution, error) {
	packets := Parse(input)
	dividers := generateDividerPackets()
	packets = append(packets, dividers...)

	sort.Slice(packets, func(i, j int) bool {
		res := equal(packets[i], packets[j])
		return res == IsEqual || res == IsOrdered
	})

	decodeKey := 1
	for i, p := range packets {
		for _, d := range dividers {
			if equal(p, d) == IsEqual {
				decodeKey *= i + 1
			}
		}
	}
	return runner.Solution{Message: fmt.Sprintf("The decode key is: %v", decodeKey)}, nil
}

func Parse(input []byte) []any {
	var packets []any
	for _, l := range strings.Split(string(input), "\n") {
		if l == "" {
			continue
		}
		packets = append(packets, parsePacket(l))
	}
	return packets
}

func parsePacket(raw string) any {
	p := []any{}
	json.Unmarshal([]byte(raw), &p)
	return p
}
