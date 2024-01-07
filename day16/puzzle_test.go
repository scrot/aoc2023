package day16_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day16"
)

//go:embed input.txt
var input []byte

const example = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

const edge = `\.
..
..`

const edge2 = `....\....
.........
../.-.\..
.........
..\...\..
.........
.........`

func TestDay16(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		{"p1Example", 1, day16.V1{}, []byte(example), 46},
		// {"p1Edge", 1, day16.V1{}, []byte(edge), 3},
		// {"p1Edge2", 1, day16.V1{}, []byte(edge2), 89},
		{"p1Input", 1, day16.V1{}, input, 7798},
		{"p2Example", 2, day16.V1{}, []byte(example), 51},
		{"p2Input", 2, day16.V1{}, input, 0},
	}

	for _, c := range cs {
		got, _ := c.version.Solve(c.input, c.part)
		if c.want != got {
			t.Errorf("want %d got %d", c.want, got)
		}

	}
}
