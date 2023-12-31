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
)

func TestDay13(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		{"p1Example", 1, day14.V1{}, []byte(example), 136},
		{"p1Input", 1, day14.V1{}, input, 109098},
		{"p2Example", 2, day14.V2{}, []byte(example), 64},
		{"p2Input", 2, day14.V2{}, input, 100064},
	}

	for _, c := range cs {
		got, _ := c.version.Solve(c.input, c.part)
		if got != c.want {
			t.Errorf("want %v got %v", c.want, got)
		}
	}
}

func benchmarkDay14(b *testing.B, s aoc2023.Solver, part int) {
	for i := 0; i < b.N; i++ {
		s.Solve(input, part)
	}
}
func BenchmarkDay1Part1V1(b *testing.B) { benchmarkDay14(b, day14.V1{}, 1) }
func BenchmarkDay1Part1V2(b *testing.B) { benchmarkDay14(b, day14.V2{}, 1) }
