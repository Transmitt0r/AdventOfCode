package day16

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

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
	return runner.Solution{Message: fmt.Sprintf("Tuning frequency: %v", 0)}, nil
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
