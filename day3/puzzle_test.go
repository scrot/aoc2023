package day3_test

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/scrot/aoc2023"
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

var bench int

func benchmarkDay3(version, part int, b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		s, err := newSolver(version)
		if err != nil {
			b.Fatal(err)
		}
		r, _ = s.Solve(input, part)
	}
	bench = r
}
func BenchmarkDay3Part1V1(b *testing.B) { benchmarkDay3(1, 1, b) }
func BenchmarkDay3Part2V1(b *testing.B) { benchmarkDay3(1, 2, b) }

func newSolver(version int) (aoc2023.Solver, error) {
	var s aoc2023.Solver
	switch version {
	case 1:
		s = day3.V1{}
	default:
		return s, fmt.Errorf("invalid version %d", version)
	}
	return s, nil
}
