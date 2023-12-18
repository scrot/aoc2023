package day13_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day13"
)

//go:embed input.txt
var input []byte

const (
	singleH = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.`

	singleV = `#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

	edgeCase = `..#.#.#..#...
..#.#.#..#...
....#...####.
..##.#.#.#..#
.#.#....#....
#####..#.#.##
.####.##.....
..######.#...
..##.#.#..###
..##...#..###
..######.#...
.####.##.....
#####..#.#.##
.#.#....#....
..##.#.#.#..#
....#...####.
..#.#.#..#...`

	edgeCase2 = `..#.#.#..#...
..#.#.#..#...
....#...####.
..##.#.#.#..#
.#.#....#....
#####..#.#.##
.####.##.....
..######.#...
..##.#.#..###
..##...#..###
..######.#...
.####.##.....
#####..#.#.##
.#.#....#....
..##.#.#.#..#
....#...####.
..#.#.#..#..`

	example = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`
)

func TestDay13(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		// {"p1SingleH", 1, day13.V1{}, []byte(singleH), 5},
		// {"p1SingleV", 1, day13.V1{}, []byte(singleV), 400},
		// {"p1EdgeCase", 1, day13.V1{}, []byte(edgeCase), 100},
		// // {"p1EdgeCase2", 1, day13.V1{}, []byte(edgeCase2), 0},
		{"p1Example", 1, day13.V1{}, []byte(example), 405},
		{"p1Input", 1, day13.V1{}, input, 27202},
	}

	for _, c := range cs {
		got, _ := c.version.Solve(c.input, c.part)
		if got != c.want {
			t.Errorf("want %v got %v", c.want, got)
		}
	}
}
