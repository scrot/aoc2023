package day15_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day15"
)

//go:embed input.txt
var input []byte

const example = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func TestDay15(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		{"p1Example", 1, day15.V1{}, []byte(example), 1320},
		{"p1Input", 1, day15.V1{}, input, 511498},
		{"p2Example", 2, day15.V1{}, []byte(example), 145},
		{"p2Input", 2, day15.V1{}, input, 284674},
	}

	for _, c := range cs {
		got, _ := c.version.Solve(c.input, c.part)
		if got != c.want {
			t.Errorf("want %d got %d", c.want, got)
		}
	}
}
