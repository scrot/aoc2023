package day12_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day12"
)

//go:embed input.txt
var input []byte

const example = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func TestDay12(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		{"p1Single", 1, day12.V1{}, []byte("???.### 1,1,3"), 1},
		{"p1Single", 1, day12.V1{}, []byte("?###???????? 3,2,1"), 10},
		{"p1Example", 1, day12.V1{}, []byte(example), 21},
		{"p1Input", 1, day12.V1{}, input, 7674},
		{"p2Example", 2, day12.V1{}, []byte(example), 525152},
		{"p2Input", 2, day12.V1{}, input, 4443895258186},
	}

	for _, c := range cs {
		got, _ := c.version.Solve(c.input, c.part)
		if got != c.want {
			t.Errorf("want %v got %v", c.want, got)
		}
	}
}
