package day2_test

import (
	_ "embed"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day2"
)

//go:embed input.txt
var input []byte

func TestDay2(t *testing.T) {
	const example = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

	cs := []struct {
		name  string
		part  int
		input []byte
		want  int
	}{
		{"example1", 1, []byte(example), 8},
		{"input", 1, input, 2268},
		{"part2example", 2, []byte(example), 2286},
		{"part2input", 2, input, 63542},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got, err := day2.V1{}.Solve(c.input, c.part)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(got, c.want) {
				t.Errorf("want %d got %d", c.want, got)
			}
		})

		t.Run("v2_"+c.name, func(t *testing.T) {
			got, err := day2.V2{}.Solve(c.input, c.part)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(got, c.want) {
				t.Errorf("want %d got %d", c.want, got)
			}
		})
	}
}

var bench int

func benchmarkDay2(version, part int, b *testing.B) {
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
func BenchmarkDay2Part1V1(b *testing.B) { benchmarkDay2(1, 1, b) }
func BenchmarkDay2Part2V1(b *testing.B) { benchmarkDay2(1, 2, b) }
func BenchmarkDay2Part1V2(b *testing.B) { benchmarkDay2(2, 1, b) }
func BenchmarkDay2Part2V2(b *testing.B) { benchmarkDay2(2, 2, b) }

func newSolver(version int) (aoc2023.Solver, error) {
	var s aoc2023.Solver
	switch version {
	case 1:
		s = day2.V1{}
	case 2:
		s = day2.V2{}
	default:
		return s, fmt.Errorf("invalid version %d", version)
	}
	return s, nil
}
