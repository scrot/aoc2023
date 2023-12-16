package day11_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day11"
)

//go:embed input.txt
var input []byte

const example = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestDay11(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		{"p1Example", 1, day11.V1{}, []byte(example), 374},
		{"p1Input", 1, day11.V1{}, input, 9418609},
		{"p2Input", 2, day11.V1{}, input, 593821230983},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got, _ := c.version.Solve(c.input, c.part)
			if got != c.want {
				t.Errorf("want %v got %v", c.want, got)
			}
		})
	}
}
