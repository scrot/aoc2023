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

func TestDay16(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		{"p1Example", 1, day16.V1{}, []byte(example), 46},
		{"p1Input", 1, day16.V1{}, input, 7798},
		{"p2Example", 2, day16.V1{}, []byte(example), 51},
		{"p2Input", 2, day16.V1{}, input, 8026},
		{"p2Example", 2, day16.V2{}, []byte(example), 51},
		{"p2Input", 2, day16.V2{}, input, 8026},
	}

	for _, c := range cs {
		got, _ := c.version.Solve(c.input, c.part)
		if c.want != got {
			t.Errorf("want %d got %d", c.want, got)
		}

	}
}

func benchmarkDay16(b *testing.B, s aoc2023.Solver, part int) {
	for i := 0; i < b.N; i++ {
		s.Solve(input, part)
	}
}

// func BenchmarkDay16Part1V1(b *testing.B) { benchmarkDay16(b, day16.V1{}, 1) }
func BenchmarkDay16Part2V1(b *testing.B) { benchmarkDay16(b, day16.V1{}, 2) }

// func BenchmarkDay16Part1V2(b *testing.B) { benchmarkDay16(b, day16.V2{}, 1) }
func BenchmarkDay16Part2V2(b *testing.B) { benchmarkDay16(b, day16.V2{}, 2) }
