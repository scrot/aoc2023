package day14

import (
	"bufio"
	"bytes"
	"fmt"
	"hash/fnv"
	"slices"
)

type V2 struct{}

type (
	loc  [2]int
	locs []loc
)

type agrid struct {
	// field without O stones
	field [][]byte

	// stone locations
	stones []loc

	// look-up tables for squares
	squares map[byte][]locs
}

func newAbstractGrid(s *bufio.Scanner) (g agrid) {
	var sqs []loc
	for y := 0; s.Scan(); y++ {
		var r []byte
		for x, b := range s.Bytes() {
			switch b {
			case 'O':
				g.stones = append(g.stones, loc{x, y})
				r = append(r, '.')
			case '#':
				sqs = append(sqs, loc{x, y})
				r = append(r, '#')
			case '.':
				r = append(r, '.')
			}
		}
		g.field = append(g.field, r)
	}

	g.squares = make(map[byte][]locs)
	g.squares['N'] = g.lookupTable('N', sqs)
	g.squares['E'] = g.lookupTable('E', sqs)
	g.squares['S'] = g.lookupTable('S', sqs)
	g.squares['W'] = g.lookupTable('W', sqs)

	return
}

func (g agrid) String() (s string) {
	var tmp [][]byte
	for _, row := range g.field {
		var r []byte
		r = append(r, row...)
		tmp = append(tmp, r)
	}

	for _, st := range g.stones {
		tmp[st[1]][st[0]] = 'O'
	}

	for _, row := range tmp {
		s += fmt.Sprintf("%s\n", row)
	}
	return
}

func (g agrid) hash() uint64 {
	h := fnv.New64a()
	for _, st := range g.stones {
		fmt.Fprintf(h, "%d%d", st[0], st[1])
	}
	return h.Sum64()
}

func (g agrid) load() (s int) {
	for _, st := range g.stones {
		s += len(g.field) - st[1]
	}
	return
}

func (g agrid) lookupTable(dir byte, ls []loc) (m []locs) {
	var (
		idx    int
		length int
	)
	switch dir {
	case 'N', 'S':
		idx = 0
		length = len(g.field[0])
	case 'W', 'E':
		idx = 1
		length = len(g.field)
	}

	for y := 0; y < length; y++ {
		var row locs
		for _, st := range ls {
			if st[idx] == y {
				row = append(row, st)
			}
		}
		if dir == 'S' || dir == 'E' {
			slices.Reverse(row)
		}
		m = append(m, row)
	}
	return
}

func (g *agrid) tilt(dir byte) {
	var (
		squares = g.squares[dir]
		stones  = g.lookupTable(dir, g.stones)
	)

	g.stones = locs{}

	switch dir {
	case 'N':
		for i := 0; i < len(g.field[0]); i++ {
			var (
				open = 0
				sts  = stones[i]
				sqs  = append(squares[i], loc{i, len(g.field[0])})
			)
			for sqI, stI := 0, 0; sqI < len(sqs) && stI < len(sts); {
				if sts[stI][1] < sqs[sqI][1] {
					g.stones = append(g.stones, loc{sts[stI][0], open})
					open++
					stI++
				} else {
					open = sqs[sqI][1] + 1
					sqI++
				}
			}
		}
	case 'S':
		for i := 0; i < len(g.field[0]); i++ {
			var (
				open = len(g.field[0]) - 1
				sts  = stones[i]
				sqs  = append(squares[i], loc{len(g.field[0]) - i - 1, -1})
			)
			for sqI, stI := 0, 0; sqI < len(sqs) && stI < len(sts); {
				if sts[stI][1] > sqs[sqI][1] {
					g.stones = append(g.stones, loc{sts[stI][0], open})
					open--
					stI++
				} else {
					open = sqs[sqI][1] - 1
					sqI++
				}
			}
		}
	case 'W':
		for i := 0; i < len(g.field); i++ {
			var (
				open = 0
				sts  = stones[i]
				sqs  = append(squares[i], loc{len(g.field[0]), i})
			)

			for sqI, stI := 0, 0; sqI < len(sqs) && stI < len(sts); {
				if sts[stI][0] < sqs[sqI][0] {
					g.stones = append(g.stones, loc{open, sts[stI][1]})
					open++
					stI++
				} else {
					open = sqs[sqI][0] + 1
					sqI++
				}
			}
		}
	case 'E':
		for i := 0; i < len(g.field); i++ {
			var (
				open = len(g.field) - 1
				sts  = stones[i] // pointer?
				sqs  = append(squares[i], loc{-1, len(g.field) - i - 1})
			)

			for sqI, stI := 0, 0; sqI < len(sqs) && stI < len(sts); {
				if sts[stI][0] > sqs[sqI][0] {
					g.stones = append(g.stones, loc{open, sts[stI][1]})
					open--
					stI++
				} else {
					open = sqs[sqI][0] - 1
					sqI++
				}
			}
		}
	}
}

func (V2) Solve(input []byte, part int) (int, error) {
	var (
		r = bytes.NewReader(input)
		s = bufio.NewScanner(r)
		g = newAbstractGrid(s)
	)

	if part == 1 {
		g.tilt('N')
		return g.load(), nil
	}

	// TODO: cycle detection loop
	var (
		cycles                = 1_000_000_000
		loopstart, looplength int
		loads                 []int
		seen                  = make(map[uint64]int)
	)

	for i := 0; i < cycles; i++ {
		g.tilt('N')
		g.tilt('W')
		g.tilt('S')
		g.tilt('E')

		loads = append(loads, g.load())

		h := g.hash()
		if _, ok := seen[h]; !ok {
			seen[h] = i
		} else {
			loopstart = seen[h]
			looplength = i - seen[h]
			break
		}
	}

	idx := (cycles-loopstart)%looplength + loopstart - 1
	return loads[idx], nil
}
