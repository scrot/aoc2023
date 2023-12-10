package day8_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day8"
)

//go:embed input.txt
var input []byte

const example = `RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`

const example2 = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

const example3 = `LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`

func TestDay8(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		{"p1Example", 1, day8.V1{}, []byte(example), 2},
		{"p1Example2", 1, day8.V1{}, []byte(example2), 6},
		{"p1Input", 1, day8.V1{}, input, 19241},
		{"p2Example3", 2, day8.V1{}, []byte(example3), 6},
		{"p2Input", 2, day8.V1{}, input, 9606140307013},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got, _ := c.version.Solve(c.input, c.part)
			if got != c.want {
				t.Errorf("want %v got %v", c.want, got)
			}
		})
	}
}

func benchmarkDay8(b *testing.B, s aoc2023.Solver, part int) {
	for i := 0; i < b.N; i++ {
		s.Solve(input, part)
	}
}

func BenchmarkDay8Part1Version1(b *testing.B) { benchmarkDay8(b, day8.V1{}, 1) }

// func BenchmarkDay8Part2Version1(b *testing.B) { benchmarkDay8(b, day8.V1{}, 2) }
// func BenchmarkDay8Part1Version2(b *testing.B) { benchmarkDay8(b, day8.V2{}, 1) }
// func BenchmarkDay8Part2Version2(b *testing.B) { benchmarkDay8(b, day8.V2{}, 2) }
