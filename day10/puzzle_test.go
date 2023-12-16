package day10_test

import (
	_ "embed"
	"testing"

	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day10"
)

//go:embed input.txt
var input []byte

const example = `.....
.S-7.
.|.|.
.L-J.
.....`

const example2 = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

const example31 = `...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`

const example32 = `..........
.S------7.
.|F----7|.
.||....||.
.||....||.
.|L-7F-J|.
.|..||..|.
.L--JL--J.
..........`

const example4 = `.F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ...`

const example5 = `SF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJIF7FJ-
L---JF-JLJ....FJLJJ7
|F|F-JF---7...L7L|7|
|FFJF7L7F-JF7..L---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func TestDay10(t *testing.T) {
	cs := []struct {
		name    string
		part    int
		version aoc2023.Solver
		input   []byte
		want    int
	}{
		{"p1Example1", 1, day10.V1{}, []byte(example), 4},
		{"p1Example2", 1, day10.V1{}, []byte(example2), 8},
		{"p1Input", 1, day10.V1{}, input, 7086},
		{"p1Example31", 2, day10.V1{}, []byte(example31), 4},
		{"p1Example31", 2, day10.V1{}, []byte(example32), 4},
		{"p1Example4", 2, day10.V1{}, []byte(example4), 8},
		{"p1Example5", 2, day10.V1{}, []byte(example5), 10},
		{"p2Input", 2, day10.V1{}, input, 317},
	}

	for _, c := range cs {
		got, _ := c.version.Solve(c.input, c.part)
		if got != c.want {
			t.Errorf("want %v got %v", c.want, got)
		}
	}
}
