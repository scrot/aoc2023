package day4_test

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day4"
)

//go:embed input.txt
var input []byte

const example = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`

func TestDay4(t *testing.T) {
	cs := []struct {
		name  string
		part  int
		input []byte
		want  int
	}{
		{"none", 1, []byte("Card   1: 0 1 2 | 3 4 5"), 0},
		{"all", 1, []byte("Card   1: 1 2 3 | 1 2 3"), 4},
		{"mixed", 1, []byte("Card   1: 2 1 3 | 1 2 3"), 4},
		{"large", 1, []byte("Card   1: 200 111 9 | 111 200 3"), 2},
		{"sum", 1, []byte("Card   1: 1 2 | 1 2\nCard   2: 1 2 | 1 2"), 4},
		{"example", 1, []byte(example), 13},
		{"input", 1, input, 23750},
		{"example", 2, []byte(example), 30},
		{"input", 2, input, 13261850},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got, err := day4.V1{}.Solve(c.input, c.part)
			if err != nil {
				t.Fatal(err)
			}

			if got != c.want {
				t.Errorf("want %d got %d", c.want, got)
			}
		})
	}
}

var bench int

func benchmarkDay4(version, part int, b *testing.B) {
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
func BenchmarkDay4Part1V1(b *testing.B) { benchmarkDay4(1, 1, b) }
func BenchmarkDay4Part2V1(b *testing.B) { benchmarkDay4(1, 2, b) }

func newSolver(version int) (aoc2023.Solver, error) {
	var s aoc2023.Solver
	switch version {
	case 1:
		s = day4.V1{}
	default:
		return s, fmt.Errorf("invalid version %d", version)
	}
	return s, nil
}
