package day3

import (
	_ "embed"
	"testing"
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
		name    string
		part    int
		version int
		input   []byte
		want    int
	}{
		{"p1empty", 1, 1, []byte(empty), 0},
		{"p1around", 1, 1, []byte(around), 8},
		{"p1corner", 1, 1, []byte(corner), 1},
		{"p1example", 1, 1, []byte(example), 4361},
		{"p1input", 1, 1, input, 556057},
		{"p2example", 2, 1, []byte(example), 467835},
		{"p2input", 2, 1, input, 0},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			res, err := Solve(c.input, c.part, c.version)
			if err != nil {
				t.Fatalf("expected no error but got '%v'", err)
			}
			if res != c.want {
				t.Errorf("expected %d but got %d", c.want, res)
			}
		})
	}
}
