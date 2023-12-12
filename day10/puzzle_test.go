package day10_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day10"
)

//go:embed input.txt
var input []byte

const example = `.....
.S-7.
.|.|.
.L-J.
.....`

const example2 = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

func TestDay10(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		{"p1Example1", 1, day10.V1{}, []byte(example), 4},
		{"p1Example2", 1, day10.V1{}, []byte(example2), 8},
		{"p1Input", 1, day10.V1{}, input, -1},
	}

	for _, c := range cs {
		got, _ := c.version.Solve(c.input, c.part)
		if got != c.want {
			t.Errorf("want %v got %v", c.want, got)
		}
	}
}
