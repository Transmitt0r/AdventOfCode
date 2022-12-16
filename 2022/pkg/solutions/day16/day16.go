package day16

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

type ValveName string
type Flow int

type CaveSystem struct {
	ReachabilityMatrix map[ValveName]map[ValveName]int
	FlowMap            map[ValveName]Flow
	TimeLimit          int
}

func (c CaveSystem) solveRecursively(currentTime int, currentPressure int, currentFlow int, currentValve ValveName, remaining []ValveName) int {
	scoreIfNoOtherValvesOpenUp := currentPressure + (c.TimeLimit-currentTime)*currentFlow
	max := scoreIfNoOtherValvesOpenUp

	for _, v := range remaining {
		timeToReachAndOpen := c.ReachabilityMatrix[currentValve][v] + 1
		if currentTime+timeToReachAndOpen >= c.TimeLimit {
			continue
		}
		newPressure := currentPressure + timeToReachAndOpen*currentFlow
		newFlow := currentFlow + int(c.FlowMap[v])
		possibleScore := c.solveRecursively(currentTime+timeToReachAndOpen, newPressure, newFlow, v, removeFromList(remaining, v))
		if possibleScore > max {
			max = possibleScore
		}
	}

	return max
}

func (c CaveSystem) solvePart2(valves []ValveName) int {
	maxPressure := 0
	mu := sync.Mutex{}
	for i := 1; i < len(valves)/2; i++ {
		for v := range CombinationsValves(valves, i) {
			go func(v []ValveName) {
				pressuresSanta := c.solveRecursively(0, 0, 0, "AA", v)
				pressureElephant := c.solveRecursively(0, 0, 0, "AA", createOpposite(valves, v))
				mu.Lock()
				if pressuresSanta+pressureElephant > maxPressure {
					maxPressure = pressuresSanta + pressureElephant
				}
				mu.Unlock()
			}(v)

		}
	}
	return maxPressure
}

func removeFromList(in []ValveName, v ValveName) []ValveName {
	new := []ValveName{}
	for _, i := range in {
		if i != v {
			new = append(new, i)
		}
	}
	return new
}

func Part1(input []byte) (runner.Solution, error) {
	adjacencyList, flowMap := Parse(input)
	toCheck := []ValveName{}
	for k := range adjacencyList {
		if flowMap[k] != 0 {
			toCheck = append(toCheck, k)
		}
	}
	reachabilityMatrix := ReachabilityList(adjacencyList)
	cave := CaveSystem{reachabilityMatrix, flowMap, 30}
	maxPressure := cave.solveRecursively(0, 0, 0, ValveName("AA"), toCheck)
	return runner.Solution{Message: fmt.Sprintf("Maximum pressure release: %v", maxPressure)}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	adjacencyList, flowMap := Parse(input)
	toCheck := []ValveName{}
	for k := range adjacencyList {
		if flowMap[k] != 0 {
			toCheck = append(toCheck, k)
		}
	}
	reachabilityMatrix := ReachabilityList(adjacencyList)
	cave := CaveSystem{reachabilityMatrix, flowMap, 26}
	maxPressure := cave.solvePart2(toCheck)
	return runner.Solution{Message: fmt.Sprintf("Maximum pressure release: %v", maxPressure)}, nil
}

func Parse(input []byte) (map[ValveName][]ValveName, map[ValveName]Flow) {
	allValves := map[ValveName][]ValveName{}
	flowMap := map[ValveName]Flow{}
	re := regexp.MustCompile(`^Valve (\w{2}) has flow rate=(\d+); tunnel(s|) lead(s|) to valve(s|) (.+)$`)
	for _, l := range strings.Split(string(input), "\n") {
		res := re.FindAllStringSubmatch(l, -1)[0]
		currentValve := res[1]
		flowRate, _ := strconv.Atoi(res[2])

		adjacentValves := []ValveName{}
		for _, con := range strings.Split(res[6], ", ") {
			adjacentValves = append(adjacentValves, ValveName(con))
		}

		allValves[ValveName(currentValve)] = adjacentValves
		flowMap[ValveName(currentValve)] = Flow(flowRate)
	}
	return allValves, flowMap
}

func ReachabilityList(adjacencyList map[ValveName][]ValveName) map[ValveName]map[ValveName]int {
	res := map[ValveName]map[ValveName]int{}
	for start := range adjacencyList {
		res[start] = Dijkstra(adjacencyList, start)
	}

	return res
}

func Dijkstra(adjacencyList map[ValveName][]ValveName, start ValveName) map[ValveName]int {
	minDistance := func(dists map[ValveName]int, spt map[ValveName]bool) ValveName {
		min := math.MaxInt
		var minValve ValveName

		for k := range adjacencyList {
			if !spt[k] && dists[k] <= min {
				min = dists[k]
				minValve = k
			}
		}

		return minValve
	}
	costs := map[ValveName]int{}
	sptSet := map[ValveName]bool{}
	for k := range adjacencyList {
		costs[k] = math.MaxInt
		sptSet[k] = false
	}

	costs[start] = 0

	for {
		node := minDistance(costs, sptSet)
		if node == "" {
			break
		}
		sptSet[node] = true

		for _, adjacent := range adjacencyList[node] {
			if costs[node]+1 < costs[adjacent] {
				costs[adjacent] = costs[node] + 1
			}
		}
	}

	return costs
}

func CombinationsValves(iterable []ValveName, r int) chan []ValveName {

	ch := make(chan []ValveName)

	go func() {

		length := len(iterable)

		for comb := range GenCombinations(length, r) {
			result := make([]ValveName, r)
			for i, val := range comb {
				result[i] = iterable[val]
			}
			ch <- result
		}

		close(ch)
	}()
	return ch
}

func GenCombinations(n, r int) <-chan []int {

	if r > n {
		panic("Invalid arguments")
	}

	ch := make(chan []int)

	go func() {
		result := make([]int, r)
		for i := range result {
			result[i] = i
		}

		temp := make([]int, r)
		copy(temp, result) // avoid overwriting of result
		ch <- temp

		for {
			for i := r - 1; i >= 0; i-- {
				if result[i] < i+n-r {
					result[i]++
					for j := 1; j < r-i; j++ {
						result[i+j] = result[i] + j
					}
					temp := make([]int, r)
					copy(temp, result) // avoid overwriting of result
					ch <- temp
					break
				}
			}
			if result[0] >= n-r {
				break
			}
		}
		close(ch)

	}()
	return ch
}

func createOpposite(all []ValveName, partial []ValveName) []ValveName {
	new := []ValveName{}

outer:
	for _, v := range all {
		for _, w := range partial {
			if v == w {
				continue outer
			}
		}
		new = append(new, v)
	}

	return new
}
