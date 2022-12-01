package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day00"
)

var (
	day = flag.Int("day", 0, "select which day to run")
)

type Day interface {
	Run(input io.Reader) error
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
	days := map[int]Day{
		0: day00.Solution,
	}

	data, err := loadData(*day)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer data.Close()

	days[*day].Run(data)
}
