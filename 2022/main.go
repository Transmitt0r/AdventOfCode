package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day00"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day01"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day02"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day03"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day04"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day05"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day06"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day07"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day08"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day09"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day10"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day11"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day12"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day13"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day14"
)

var (
	day  = flag.Int("day", 0, "select which day to run")
	part = flag.Int("part", 1, "select which part to run")
)

type Day interface {
	Run(input io.Reader, partSelect int) error
}

func loadData(fileNameNoExt int) (io.ReadCloser, error) {
	file, err := os.Open("inputs/" + fmt.Sprintf("%02d", fileNameNoExt) + ".txt")
	if err != nil {
		return file, err
	}
	return file, nil
}

func main() {
	flag.Parse()
	if !(*part == 1 || *part == 2) {
		log.Fatalf("Invalid Part: %v, Please select part 1 or 2", *part)
	}
	days := map[int]Day{
		0:  day00.Solution,
		1:  day01.Solution,
		2:  day02.Solution,
		3:  day03.Solution,
		4:  day04.Solution,
		5:  day05.Solution,
		6:  day06.Solution,
		7:  day07.Solution,
		8:  day08.Solution,
		9:  day09.Solution,
		10: day10.Solution,
		11: day11.Solution,
		12: day12.Solution,
		13: day13.Solution,
		14: day14.Solution,
	}
	if _, ok := days[*day]; !ok {
		log.Fatalln("Selected invalid day!")
	}
	data, err := loadData(*day)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer data.Close()

	err = days[*day].Run(data, *part-1)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
