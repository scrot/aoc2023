package day1_test

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/scrot/aoc2023"
	"github.com/scrot/aoc2023/day1"
)

//go:embed input.txt
var input []byte

func TestDay1(t *testing.T) {
	const (
		example1 = `1abc2
    pqr3stu8vwx
    a1b2c3d4e5f
    treb7uchet`

		example2 = `two1nine
    eightwothree
    abcone2threexyz
    xtwone3four
    4nineeightseven2
    zoneight234
    7pqrstsixteen`
	)
	cs := []struct {
		name  string
		part  int
		input []byte
		want  int
	}{
		{"outside", 1, []byte("1aabb2"), 12},
		{"inside", 1, []byte("a1ab2b"), 12},
		{"left", 1, []byte("1ab2b"), 12},
		{"right", 1, []byte("a1ab2"), 12},
		{"onechar", 1, []byte("1"), 11},
		{"example1", 1, []byte(example1), 142},
		{"input1", 1, input, 54605},
		{"txtleft", 2, []byte("oneaabb2"), 12},
		{"txtright", 2, []byte("1aabbtwo"), 12},
		{"txtboth", 2, []byte("oneaa1bbtwo"), 12},
		{"txtmiddle", 2, []byte("aaone3twobb"), 12},
		{"tricky", 2, []byte("threeight1twone"), 31},
		{"txtsingle", 2, []byte("aonea"), 11},
		{"txtshort", 2, []byte("one"), 11},
		{"example2", 2, []byte(example2), 281},
		{"input2", 2, input, 55429},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got, err := day1.V1{}.Solve(c.input, c.part)
			if err != nil {
				t.Fatalf("expected no error but got %s", err)
			}

			t.Logf("answer: %d", got)
			if !cmp.Equal(c.want, got) {
				t.Fatalf("want %d got %d", c.want, got)
			}
		})

		t.Run("v2_"+c.name, func(t *testing.T) {
			got, err := day1.V2{}.Solve(c.input, c.part)
			if err != nil {
				t.Fatalf("expected no error but got %s", err)
			}

			t.Logf("answer: %d", got)
			if !cmp.Equal(c.want, got) {
				t.Fatalf("want %d got %d", c.want, got)
			}
		})
	}
}

func benchmarkDay1(b *testing.B, s aoc2023.Solver, part int) {
	for i := 0; i < b.N; i++ {
		s.Solve(input, part)
	}
}
func BenchmarkDay1Part1V1(b *testing.B) { benchmarkDay1(b, day1.V1{}, 1) }
func BenchmarkDay1Part2V1(b *testing.B) { benchmarkDay1(b, day1.V1{}, 2) }
func BenchmarkDay1Part1V2(b *testing.B) { benchmarkDay1(b, day1.V2{}, 1) }
func BenchmarkDay1Part2V2(b *testing.B) { benchmarkDay1(b, day1.V2{}, 2) }
