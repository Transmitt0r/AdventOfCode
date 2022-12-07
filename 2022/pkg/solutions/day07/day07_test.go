package day07_test

import (
	"testing"

	"github.com/Transmitt0r/AdventOfCode/2022/pkg/solutions/day07"
)

var sampleInput = []byte(`$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`)

func TestPart1(t *testing.T) {
	sol, err := day07.Part1(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Total size of directories of at most 100k: 95437" {
		t.Errorf("Wrong solution %v", sol)
	}
}

func TestPart2(t *testing.T) {
	sol, err := day07.Part2(sampleInput)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if sol.Message != "Smallest directory that can be deleted: 24933642" {
		t.Errorf("Wrong solution %v", sol)
	}
}
