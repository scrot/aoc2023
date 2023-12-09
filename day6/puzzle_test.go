package day6_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day6"
)

//go:embed input.txt
var input []byte

const example = `Time:      7  15   30
Distance:  9  40  200`

func TestDay6(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		{"p1 example", 1, day6.V1{}, []byte(example), 288},
		{"p1 input", 1, day6.V1{}, input, 1159152},
		{"p1 example", 2, day6.V1{}, []byte(example), 71503},
		{"p2 input", 2, day6.V1{}, input, 41513103},
		{"p1 example", 1, day6.V2{}, []byte(example), 288},
		{"p1 input", 1, day6.V2{}, input, 1159152},
		{"p1 example", 2, day6.V2{}, []byte(example), 71503},
		{"p2 input", 2, day6.V2{}, input, 41513103},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got, _ := c.version.Solve(c.input, c.part)
			if got != c.want {
				t.Errorf("expected %v got %v", c.want, got)
			}
		})
	}
}

func benchmarkDay6(b *testing.B, s aoc2023.Solver, part int) {
	for i := 0; i < b.N; i++ {
		s.Solve(input, part)
	}
}

func BenchmarkDay6Part1Version1(b *testing.B) { benchmarkDay6(b, day6.V1{}, 1) }
func BenchmarkDay6Part2Version1(b *testing.B) { benchmarkDay6(b, day6.V1{}, 2) }
func BenchmarkDay6Part1Version2(b *testing.B) { benchmarkDay6(b, day6.V2{}, 1) }
func BenchmarkDay6Part2Version2(b *testing.B) { benchmarkDay6(b, day6.V2{}, 2) }
