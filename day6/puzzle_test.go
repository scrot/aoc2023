package day6_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023/day6"
)

//go:embed input.txt
var input []byte

const example = `Time:      7  15   30
Distance:  9  40  200`

func TestDay6(t *testing.T) {
	cs := []struct {
		name  string
		part  int
		input []byte
		want  int
	}{
		{"example", 1, []byte(example), 288},
		{"input", 1, input, 1159152},
		{"example", 2, []byte(example), 71503},
		{"input", 2, input, 0},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got, _ := day6.V1{}.Solve(c.input, c.part)
			if got != c.want {
				t.Errorf("expected %v got %v", c.want, got)
			}
		})
	}
}
