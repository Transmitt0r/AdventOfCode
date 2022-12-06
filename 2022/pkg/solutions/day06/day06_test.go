package day06_test

import (
	"fmt"
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day06"
)

var sampleInputs = []struct {
	In           []byte
	FirstMarker  int
	MessageStart int
}{
	{[]byte("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), 7, 19},
	{[]byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), 5, 23},
	{[]byte("nppdvjthqldpwncqszvftbrmjlhg"), 6, 23},
	{[]byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 10, 29},
	{[]byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 11, 26},
}

func TestPart1(t *testing.T) {
	for _, tc := range sampleInputs {
		t.Run(fmt.Sprintf("Test %v", string(tc.In)), func(t *testing.T) {
			sol, err := day06.Part1(tc.In)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if sol.Message != fmt.Sprintf("First marker after: %v", tc.FirstMarker) {
				t.Errorf("Wrong solution %v, expected %v", sol, tc.FirstMarker)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	for _, tc := range sampleInputs {
		t.Run(fmt.Sprintf("Test %v", string(tc.In)), func(t *testing.T) {
			sol, err := day06.Part2(tc.In)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
			if sol.Message != fmt.Sprintf("First message after: %v", tc.MessageStart) {
				t.Errorf("Wrong solution %v, expected %v", sol, tc.MessageStart)
			}
		})
	}
}
