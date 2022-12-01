package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day00"
	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day01"
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
		0: day00.Solution,
		1: day01.Solution,
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
