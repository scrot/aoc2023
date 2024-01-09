package day16

import (
	"bufio"
	"bytes"
)

type V1 struct{}

func (V1) Solve(input []byte, part int) (int, error) {
	var (
		r = bytes.NewReader(input)
		s = bufio.NewScanner(r)
		g = newGrid(s)
	)

	if part == 1 {
		start := beam{loc{-1, 0}, DirRight}
		return energize(start, g), nil
	}

	if part == 2 {
		var maxEnergized int
		for _, b := range g.genBeams() {
			maxEnergized = max(maxEnergized, energize(b, g))
		}
		return maxEnergized, nil
	}

	return 0, nil
}

func energize(start beam, g grid) int {
	var (
		energized = make(map[loc]bool)
		visited   = make(map[beam]bool)
		b, bs     = beam{}, []beam{start}
	)

	for len(bs) > 0 {
		b, bs = bs[0], bs[1:]

		if _, ok := visited[b]; ok {
			continue
		}
		visited[b] = true

		if b != start {
			// g[b.pos.y][b.pos.x].energized = true
			energized[b.pos] = true
		}

		bs = append(bs, b.step(&g)...)
	}

	return len(energized)
}

type grid [][]tile

func newGrid(s *bufio.Scanner) (g grid) {
	for s.Scan() {
		var r []tile
		for _, b := range s.Bytes() {
			r = append(r, tile{b, false})
		}
		g = append(g, r)
	}
	return
}

func (g grid) genBeams() (bs []beam) {
	for y := range g {
		bs = append(bs, beam{loc{-1, y}, DirRight})
		bs = append(bs, beam{loc{len(g[0]), y}, DirLeft})
	}
	for x := range g[0] {
		bs = append(bs, beam{loc{x, -1}, DirDown})
		bs = append(bs, beam{loc{x, len(g)}, DirUp})
	}
	return
}

func (g grid) energized() (count int) {
	for _, r := range g {
		for _, t := range r {
			if t.energized {
				count++
			}
		}
	}
	return
}

type tile struct {
	sym       byte
	energized bool
}

type loc struct {
	x, y int
}

var (
	DirUp    = loc{0, -1}
	DirDown  = loc{0, 1}
	DirLeft  = loc{-1, 0}
	DirRight = loc{1, 0}
)

type beam struct {
	pos, dir loc
}

func (b beam) step(g *grid) []beam {
	// out of bound
	xm, ym := len((*g)[0]), len((*g))
	n := loc{b.pos.x + b.dir.x, b.pos.y + b.dir.y}
	if n.x < 0 || n.x >= xm || n.y < 0 || n.y >= ym {
		return []beam{}
	}

	switch (*g)[n.y][n.x].sym {
	case '/':
		switch b.dir {
		case DirUp:
			return []beam{{n, DirRight}}
		case DirDown:
			return []beam{{n, DirLeft}}
		case DirLeft:
			return []beam{{n, DirDown}}
		case DirRight:
			return []beam{{n, DirUp}}
		}
	case '\\':
		switch b.dir {
		case DirUp:
			return []beam{{n, DirLeft}}
		case DirDown:
			return []beam{{n, DirRight}}
		case DirLeft:
			return []beam{{n, DirUp}}
		case DirRight:
			return []beam{{n, DirDown}}
		}
	case '|':
		switch b.dir {
		case DirLeft, DirRight:
			return []beam{
				{n, DirUp},
				{n, DirDown},
			}
		case DirUp, DirDown:
			return []beam{{n, b.dir}}
		}
	case '-':
		switch b.dir {
		case DirLeft, DirRight:
			return []beam{{n, b.dir}}
		case DirUp, DirDown:
			return []beam{
				{n, DirLeft},
				{n, DirRight},
			}
		}
	default:
		return []beam{{n, b.dir}}
	}

	return []beam{}
}
