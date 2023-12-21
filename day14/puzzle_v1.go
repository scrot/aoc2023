package day14

import (
	"bufio"
	"bytes"
	"fmt"
)

type V1 struct{}

type grid [][]byte

func (g grid) String() (s string) {
	for _, r := range g {
		s += string(r) + "\n"
	}
	return
}

func newGrid(s *bufio.Scanner) (g grid) {
	for s.Scan() {
		row := make([]byte, len(s.Bytes()))
		copy(row, s.Bytes())
		g = append(g, row)
	}
	fmt.Printf("%d %d\n", len(g), len(g[0]))
	return
}

func (V1) Solve(input []byte, part int) (int, error) {
	var (
		s = bufio.NewScanner(bytes.NewReader(input))
		g = newGrid(s)
	)

	fmt.Println(g)

	var sum int
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
			// fmt.Printf("weight:%d total:%d sum:%d\n", weight, total, sum)
		}
		sum += total
		// fmt.Printf("col: %d sum: %d\n\n", x, sum)
	}

	return sum, nil
}
