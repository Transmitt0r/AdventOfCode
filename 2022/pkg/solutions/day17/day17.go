package day17

import (
	"fmt"
	"strings"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/runner"
)

var Solution = runner.Runner{Runnables: []runner.Runnable{Part1, Part2}}

type JetDirection rune
type CaveMaterial rune

type Shape [][]CaveMaterial

func (s Shape) Height() int {
	return len(s)
}

func (s Shape) Width() int {
	return len(s[0])
}

const (
	JetRight JetDirection = '>'
	JetLeft  JetDirection = '<'

	Air          CaveMaterial = '.'
	Rock         CaveMaterial = '#'
	FallingShape CaveMaterial = '@'
)

var (
	Shape1 Shape = Shape{
		{FallingShape, FallingShape, FallingShape, FallingShape},
	}
	Shape2 Shape = Shape{
		{Air, FallingShape, Air},
		{FallingShape, FallingShape, FallingShape},
		{Air, FallingShape, Air},
	}
	Shape3 Shape = Shape{
		{Air, Air, FallingShape},
		{Air, Air, FallingShape},
		{FallingShape, FallingShape, FallingShape},
	}
	Shape4 Shape = Shape{
		{FallingShape},
		{FallingShape},
		{FallingShape},
		{FallingShape},
	}
	Shape5 Shape = Shape{
		{FallingShape, FallingShape},
		{FallingShape, FallingShape},
	}
)

func Generator[P interface{}](pattern []P) <-chan P {
	patternChan := make(chan P)
	go func() {
		i := 0
		for {
			patternChan <- pattern[i]
			i++
			if i >= len(pattern) {
				i = 0
			}
		}
	}()
	return patternChan
}

type Cave struct {
	curShape Shape
	width    int
	caveMap  [][]CaveMaterial
	offset   int
}

func NewCave(width int) *Cave {
	initialMap := [][]CaveMaterial{{}}
	for i := 0; i < width; i++ {
		initialMap[0] = append(initialMap[0], Air)
	}
	return &Cave{width: width, caveMap: initialMap}
}

func (c *Cave) AddShape(s Shape) {
	c.curShape = s

	newCaveMap := [][]CaveMaterial{}

	renderedTowerheight := c.MaxHeight() - c.MinHeight()

	caveMapHeight := renderedTowerheight + 3 + s.Height()

	for i := 0; i < caveMapHeight; i++ {
		line := make([]CaveMaterial, c.width)
		if i <= renderedTowerheight {
			line = c.caveMap[i]
		} else {
			for j := 0; j < c.width; j++ {
				line[j] = Air
			}
		}

		newCaveMap = append(newCaveMap, line)
	}
	c.caveMap = newCaveMap

	c.putShapeIntoCaveMap(s, 2, caveMapHeight-s.Height())
}

func (c Cave) Render() string {
	lines := []string{}
	for i := len(c.caveMap) - 1; i >= 0; i-- {
		line := []string{}
		for j := 0; j < len(c.caveMap[i]); j++ {
			line = append(line, string(c.caveMap[i][j]))
		}
		lines = append(lines, strings.Join(line, ""))
	}
	lines = append(lines, "-------------")
	return strings.Join(lines, "\n")
}

func (c *Cave) putShapeIntoCaveMap(s Shape, bottomLeftX, bottomLeftY int) {
	for y := 0; y < len(s); y++ {
		for x := 0; x < len(s[y]); x++ {
			c.caveMap[bottomLeftY+y][bottomLeftX+x] = s[len(s)-y-1][x]
		}
	}
}

func (c *Cave) SolidifyFallingShape() {
	for y := 0; y < len(c.caveMap); y++ {
		for x := 0; x < len(c.caveMap[y]); x++ {
			if c.caveMap[y][x] == FallingShape {
				c.caveMap[y][x] = Rock
			}
		}
	}
}

func (c *Cave) MoveCurrentShape(j JetDirection) {
	caveMapCopy := [][]CaveMaterial{}
	for y := 0; y < len(c.caveMap); y++ {
		line := []CaveMaterial{}
		for x := 0; x < len(c.caveMap[y]); x++ {
			line = append(line, c.caveMap[y][x])
		}
		caveMapCopy = append(caveMapCopy, line)
	}
	replace(caveMapCopy, FallingShape, Air)

	for y := 0; y < len(c.caveMap); y++ {
		for x := 0; x < len(c.caveMap[y]); x++ {
			switch j {
			case JetLeft:
				if c.caveMap[y][x] == FallingShape && (x-1 < 0 || c.caveMap[y][x-1] == Rock) {
					return
				}
				if c.caveMap[y][x] == FallingShape {
					caveMapCopy[y][x-1] = FallingShape
				}
			case JetRight:
				if c.caveMap[y][x] == FallingShape && (x+1 >= len(c.caveMap[y]) || c.caveMap[y][x+1] == Rock) {
					return
				}
				if c.caveMap[y][x] == FallingShape {
					caveMapCopy[y][x+1] = FallingShape
				}
			}
		}
	}
	c.caveMap = caveMapCopy
}

func replace(caveMap [][]CaveMaterial, original CaveMaterial, replacement CaveMaterial) [][]CaveMaterial {
	for y := 0; y < len(caveMap); y++ {
		for x := 0; x < len(caveMap[y]); x++ {
			if caveMap[y][x] == original {
				caveMap[y][x] = replacement
			}
		}
	}
	return caveMap
}

func (c *Cave) DropShape() error {
	caveMapCopy := [][]CaveMaterial{}
	for y := 0; y < len(c.caveMap); y++ {
		line := []CaveMaterial{}
		for x := 0; x < len(c.caveMap[y]); x++ {
			line = append(line, c.caveMap[y][x])
		}
		caveMapCopy = append(caveMapCopy, line)
	}
	replace(caveMapCopy, FallingShape, Air)

	for y := 0; y < len(caveMapCopy); y++ {
		for x := 0; x < len(caveMapCopy[y]); x++ {
			if c.caveMap[y][x] == FallingShape && (y-1 < 0 || c.getMaterialAt(x, y-1) == Rock) {
				return fmt.Errorf("cannot drop")
			}
			if c.caveMap[y][x] == FallingShape {
				caveMapCopy[y-1][x] = FallingShape
			}
		}
	}
	c.caveMap = caveMapCopy
	return nil
}

func (c *Cave) getMaterialAt(x, y int) CaveMaterial {
	return c.caveMap[y][x]
}

func (c Cave) MaxHeight() int {
	for y := len(c.caveMap) - 1; y >= 0; y-- {
		for x := 0; x < len(c.caveMap[y]); x++ {
			if c.caveMap[y][x] == Rock {
				return y + 1
			}
		}
	}
	return 0
}

func (c Cave) MinHeight() int {
	for h, y := range c.caveMap {
		for _, x := range y {
			if x == Rock {
				return h
			}
		}
	}
	return 0
}

func Part1(input []byte) (runner.Solution, error) {
	jetPattern := Parse(input)
	jetChan := Generator(jetPattern)
	shapeChan := Generator([]Shape{Shape1, Shape2, Shape3, Shape4, Shape5})
	i := 0
	c := NewCave(7)
	for i < 2022 {
		c.AddShape(<-shapeChan)
		dropsNext := false
		for {
			//fmt.Println(c.Render())
			if !dropsNext {
				c.MoveCurrentShape(<-jetChan)
				dropsNext = true
			} else {
				err := c.DropShape()
				dropsNext = false
				if err != nil {
					c.SolidifyFallingShape()
					break
				}
			}
		}
		//fmt.Println(c.Render())
		i++
	}
	return runner.Solution{Message: fmt.Sprintf("Height of the tower: %v", c.MaxHeight())}, nil
}

func Part2(input []byte) (runner.Solution, error) {
	return runner.Solution{Message: fmt.Sprintf("Maximum pressure release: %v", 0)}, nil
}

func Parse(input []byte) []JetDirection {
	pattern := []JetDirection{}
	for _, r := range string(input) {
		pattern = append(pattern, JetDirection(r))
	}
	return pattern
}
