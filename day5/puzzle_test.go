package day5_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023/day5"
)

//go:embed input.txt
var input []byte

//go:embed example.txt
var example []byte

func TestDay5(t *testing.T) {
	cs := []struct {
		name  string
		part  int
		input []byte
		want  int
	}{
		{"example", 1, example, 35},
		{"input", 1, input, 289863851},
		{"p2Example", 2, example, 46},
		{"input", 2, input, 60568880},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got, err := day5.V1{}.Solve(c.input, c.part)
			if err != nil {
				t.Fatal(err)
			}

			if got != c.want {
				t.Errorf("want %d got %d", c.want, got)
			}
		})
	}
}
