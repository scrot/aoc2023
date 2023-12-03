package day2

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	//go:embed input.txt
	input []byte

	//go:embed example.txt
	example []byte
)

func TestDay2(t *testing.T) {
	cs := []struct {
		name  string
		part  int
		input []byte
		want  int
	}{
		{"example1", 1, example, 8},
		{"input", 1, input, 2268},
		{"part2example", 2, example, 2286},
		{"part2input", 2, input, 63542},
	}

	for _, c := range cs {
		t.Run(c.name, func(t *testing.T) {
			got, err := SolveV1(c.input, c.part)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(got, c.want) {
				t.Errorf("want %d got %d", c.want, got)
			}
		})

		t.Run("alt_"+c.name, func(t *testing.T) {
			got, err := SolveV2(c.input, c.part)
			if err != nil {
				t.Fatal(err)
			}

			if !cmp.Equal(got, c.want) {
				t.Errorf("want %d got %d", c.want, got)
			}
		})
	}
}

var bench int

func benchmarkDay2(version, part int, b *testing.B) {
	var r int
	for i := 0; i < b.N; i++ {
		r, _ = Solve(input, version, part)
	}
	bench = r
}
func BenchmarkDay2Part1V1(b *testing.B) { benchmarkDay2(1, 1, b) }
func BenchmarkDay2Part2V1(b *testing.B) { benchmarkDay2(1, 2, b) }
func BenchmarkDay2Part1V2(b *testing.B) { benchmarkDay2(2, 1, b) }
func BenchmarkDay2Part2V2(b *testing.B) { benchmarkDay2(2, 2, b) }
