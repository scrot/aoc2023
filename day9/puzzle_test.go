package day9_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day9"
)

//go:embed input.txt
var input []byte

const example = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestDay9(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		{"p1Example", 1, day9.V1{}, []byte(example), 114},
		{"p1Input", 1, day9.V1{}, input, 1743490457},
		{"p2Example", 2, day9.V1{}, []byte(example), 2},
		{"p2Input", 2, day9.V1{}, input, 1053},
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
