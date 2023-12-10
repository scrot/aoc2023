package day7_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day7"
)

//go:embed input.txt
var input []byte

const example = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

const edge = `12345 100
54321 1
12543 10`

func TestDay7(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		{"p1Edge", 1, day7.V1{}, []byte(edge), 123},
		{"p1Example", 1, day7.V1{}, []byte(example), 6440},
		{"p1Input", 1, day7.V1{}, input, 253603890},
		{"p2Example", 2, day7.V1{}, []byte(example), 5905},
		{"p2Input", 2, day7.V1{}, input, 0},
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
