package day14_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day14"
)

//go:embed input.txt
var input []byte

const (
	example = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

	// 8 + 2
	small = `O
.
#
.
.
#
.
O`
	small2 = `.
.
.`
	small3 = `O
O
O`
)

func TestDay13(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		// {"p1Small", 1, day14.V1{}, []byte(small), 10},
		// {"p1Small2", 1, day14.V1{}, []byte(small2), 0},
		// {"p1Small3", 1, day14.V1{}, []byte(small3), 6},
		// {"p1Example", 1, day14.V1{}, []byte(example), 136},
		{"p1Input", 1, day14.V1{}, input, 109098},
	}

	for _, c := range cs {
		got, _ := c.version.Solve(c.input, c.part)
		if got != c.want {
			t.Errorf("want %v got %v", c.want, got)
		}
	}
}
