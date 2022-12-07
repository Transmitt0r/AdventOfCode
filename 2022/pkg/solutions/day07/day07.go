package day07

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

type Node struct {
	Name     string
	IsDir    bool
	size     int
	Children []*Node
}

type StatefulParser struct {
	Root        *Node
	currentPath []*Node
}

func (s *StatefulParser) ParseLine(line string) {
	args := strings.Split(line, " ")
	if args[0] == "$" {
		// is a command
		switch args[1] {
		case "cd":
			switch args[2] {
			case "..":
				s.currentPath = s.currentPath[:len(s.currentPath)-1]
			case "/":
				s.currentPath = []*Node{s.Root}
			default:
				for _, c := range s.currentPath[len(s.currentPath)-1].Children {
					if c.Name == args[2] {
						s.currentPath = append(s.currentPath, c)
						break
					}
				}
			}
		case "ls":
			//nothing really happens here
		}
	} else {
		currentNode := s.currentPath[len(s.currentPath)-1]
		if !currentNode.IsDir {
			panic("This should not happen")
		}

		newNode := &Node{
			Name: args[1],
		}

		if args[0] == "dir" {
			newNode.IsDir = true
		} else {
			size, _ := strconv.Atoi(args[0])
			newNode.size = size
		}
		currentNode.Children = append(currentNode.Children, newNode)
	}
}

func (n *Node) Size() int {
	// Base Case, Node is file
	if !n.IsDir {
		return n.size
	}

	// Node is dir, so sum up all the sizes of the children
	size := 0
	for _, child := range n.Children {
		size += child.Size()
	}
	return size
}

func Part1(input []byte) (runner.Solution, error) {
	sizeUnder100k := 0
	p := Parse(input)
	nodesToCheck := []*Node{p.Root}
	for len(nodesToCheck) > 0 {
		currentNode := nodesToCheck[0]
		nodesToCheck = nodesToCheck[1:]

		if currentNode.IsDir {
			size := currentNode.Size()
			if size <= 100000 {
				sizeUnder100k += size
			}
			nodesToCheck = append(nodesToCheck, currentNode.Children...)
		}
	}
	return runner.Solution{Message: fmt.Sprintf("Total size of directories of at most 100k: %v", sizeUnder100k)}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	totalDiskSize := 70000000
	necessarySpace := 30000000
	p := Parse(input)
	unusedSpace := totalDiskSize - p.Root.Size()
	additionalSpaceRequired := necessarySpace - unusedSpace
	nodesToCheck := []*Node{p.Root}
	candidates := []int{}
	for len(nodesToCheck) > 0 {
		currentNode := nodesToCheck[0]
		nodesToCheck = nodesToCheck[1:]

		if currentNode.IsDir {
			size := currentNode.Size()
			if size >= additionalSpaceRequired {
				candidates = append(candidates, size)
			}
			nodesToCheck = append(nodesToCheck, currentNode.Children...)
		}
	}

	sort.Ints(candidates)
	return runner.Solution{Message: fmt.Sprintf("Smallest directory that can be deleted: %v", candidates[0])}, nil
}

func Parse(input []byte) *StatefulParser {
	parser := StatefulParser{
		Root: &Node{
			Name:  "/",
			IsDir: true,
		},
	}

	for _, l := range strings.Split(string(input), "\n") {
		parser.ParseLine(l)
	}
	return &parser
}
