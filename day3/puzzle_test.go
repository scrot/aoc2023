package day3_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023/day3"
)

//go:embed input.txt
var input []byte

const (
	empty = `...
...
...`
	around = `$...$...$...............
.1..1..1.$1..1$.1..1..1.
...............$...$...$`
	corner = `...
.$.
..1`
	cicled = `111
1*1
111`
	spaced = `11111
1...1
1.*.1
1...1
11111`
	example = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
)

func TestDay3(t *testing.T) {
	cs := []struct {
		name  string
		part  int
		input []byte
		want  int
	}{
		{"p1empty", 1, []byte(empty), 0},
		{"p1around", 1, []byte(around), 8},
		{"p1corner", 1, []byte(corner), 1},
		{"p1example", 1, []byte(example), 4361},
		{"p1input", 1, input, 556057},
		{"p2digits", 2, []byte(cicled), 0},
		{"p2spaced", 2, []byte(spaced), 0},
		{"p2example", 2, []byte(example), 467835},
		{"p2input", 2, input, 82824352},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			res, err := day3.V1{}.Solve(c.input, c.part)
			if err != nil {
				t.Fatalf("expected no error but got '%v'", err)
			}
			if res != c.want {
				t.Errorf("expected %d but got %d", c.want, res)
			}
		})
	}
}
