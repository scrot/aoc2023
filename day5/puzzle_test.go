package day5_test

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/scrot/aoc2023/day5"
)

//go:embed input.txt
var input []byte

//go:embed example.txt
var example []byte

const (
	leftof = `seeds: 1 2
seeds-to-soil-map:
0 4 1
`
	rightof = `seeds: 5 2
seeds-to-soil-map:
0 4 1
`
	within = `seeds: 5 1
seeds-to-soil-map:
0 4 3
`
	over = `seeds: 3 5
seeds-to-soil-map:
0 4 3
`
	exactly = `seeds: 4 3
seeds-to-soil-map:
0 4 3
`
	leftoverlap = `seeds: 3 3
seeds-to-soil-map:
0 4 3
`
	rightoverlap = `seeds: 5 3
seeds-to-soil-map:
0 4 3
`
)

func TestDay5(t *testing.T) {
	cs := []struct {
		name  string
		part  int
		input []byte
		want  int
	}{
		// {"overlap", 1, []byte(within), 1},
		{"example", 1, example, 35},
		// {"input", 1, input, 289863851},
		// {"p2LeftOf", 2, []byte(leftof), 1},
		// {"p2RightOf", 2, []byte(rightof), 1},
		// {"p2Within", 2, []byte(within), 1},
		// {"p2Over", 2, []byte(over), 0},
		// {"p2Exactly", 2, []byte(exactly), 0},
		// {"p2OLeftOverlap", 2, []byte(leftoverlap), 0},
		// {"p2RightOverlap", 2, []byte(rightoverlap), 1},
		// {"p2Example", 2, example, 46},
		// {"input", 2, input, 0},
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

func TestDay5Convert(t *testing.T) {
	got := day5.Convert([]int{1, 2, 5})
	want := day5.Mapping{2, 6, -1}
	if !cmp.Equal(got, want) {
		t.Errorf("want %v got %v", want, got)
	}
}

func TestDay5Overlapping(t *testing.T) {
	single := day5.Mapping{5, 5, 0}
	offset := day5.Mapping{4, 4, 1}

	cs := []struct {
		name string
		seed day5.Mapping
		dest day5.Mapping
		want day5.Mapping
	}{
		{"singleLeft", single, day5.Mapping{6, 7, 2}, day5.Mapping{}},
		{"singleRight", single, day5.Mapping{2, 4, 2}, day5.Mapping{}},
		{"singleIn1", single, day5.Mapping{5, 6, 3}, day5.Mapping{single.Start, single.End, 3}},
		{"singleIn2", single, day5.Mapping{4, 5, 3}, day5.Mapping{single.Start, single.End, 3}},
		{"offsetLeft", offset, day5.Mapping{6, 7, 2}, day5.Mapping{}},
		{"offsetRight", offset, day5.Mapping{2, 4, 2}, day5.Mapping{}},
		{"offsetIn1", offset, day5.Mapping{5, 6, 3}, day5.Mapping{offset.Start, offset.End, 4}},
		{"offsetIn2", offset, day5.Mapping{4, 5, 3}, day5.Mapping{offset.Start, offset.End, 4}},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got := day5.Overlapping(c.seed, c.dest)
			if !cmp.Equal(got, c.want) {
				t.Errorf("want %v got %v", c.want, got)
			}
		})
	}
}

func TestDay5Destination(t *testing.T) {
	single := day5.Mapping{5, 5, 0}
	offset := day5.Mapping{5, 5, -4}

	rs1 := []day5.Mapping{{0, 1, 5}, {3, 5, -3}}

	cs := []struct {
		name string
		seed day5.Mapping
		dest []day5.Mapping
		want day5.Mapping
	}{
		{"single", single, rs1, day5.Mapping{5, 5, -3}},
		{"offset", offset, rs1, day5.Mapping{5, 5, 1}},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got := day5.Destination(c.seed, c.dest)
			if !cmp.Equal(got, c.want) {
				t.Errorf("want %v got %v", c.want, got)
			}
		})
	}
}
