package day14

import (
	"bufio"
	"bytes"
)

type V1 struct{}

type grid [][]byte

func (g grid) northSum() (sum int) {
	for x := 0; x < len(g[0]); x++ {
		total, weight := 0, len(g)
		for y := 0; y < len(g); y++ {
			switch g[y][x] {
			case 'O':
				total += weight
				weight--
			case '#':
				sum += total
				total = 0
				weight = len(g) - (y + 1)
			}
		}
		sum += total
	}
	return
}

func (g grid) String() (s string) {
	for _, r := range g {
		s += string(r) + "\n"
	}
	return
}

func newGrid(s *bufio.Scanner) (g grid) {
	for y := 0; s.Scan(); y++ {
		row := make([]byte, len(s.Bytes()))
		copy(row, s.Bytes())
		g = append(g, row)
	}
	return
}

func (V1) Solve(input []byte, part int) (int, error) {
	var (
		r = bytes.NewReader(input)
		s = bufio.NewScanner(r)
		g = newGrid(s)
	)

	// part 1
	return g.northSum(), nil
}
