package day13

import (
	"errors"
	"strings"
)

type grid [][]byte

func (g grid) String() (s string) {
	for _, r := range g {
		s += string(r) + "\n"
	}
	return
}

func (g grid) reflection(diff int) (rs int) {
	for _, p := range g.pairs(diff) {
		if rs = g.hasReflection(p, diff); rs > 0 {
			return
		}
	}
	return
}

func (g grid) pairs(smudge int) (ps [][2]int) {
	for i := 0; i < len(g)-1; i++ {
		if difference(g[i], g[i+1]) <= smudge {
			ps = append(ps, [2]int{i, i + 1})
		}
	}
	return
}

func (g grid) hasReflection(p [2]int, smudge int) int {
	for u, d := p[1], p[0]; d >= 0 && u < len(g); u, d = u+1, d-1 {
		smudge -= difference(g[d], g[u])
		if smudge < 0 {
			return 0
		}
	}

	if smudge > 0 {
		return 0
	}

	return p[0] + 1
}

func difference(s1, s2 []byte) int {
	var count int
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			count++
		}
	}
	return count
}

func (g grid) transpose() grid {
	p := make(grid, len(g[0]))
	for x := 0; x < len(g[0]); x++ {
		row := make([]byte, len(g))
		for y := 0; y < len(g); y++ {
			row = append(row, g[y][x])
		}
		p[x] = row
	}
	return p
}

func newGrid(ls []string) grid {
	g := make(grid, len(ls))
	for i, l := range ls {
		row := make([]byte, len(l))
		copy(row, l)
		g[i] = row
	}
	return g
}

type V1 struct{}

func (V1) Solve(input []byte, part int) (int, error) {
	var (
		fields     = strings.Split(strings.TrimSpace(string(input)), "\n\n")
		vsum, hsum int
		vri, hri   int
	)

	for _, f := range fields {
		v := newGrid(strings.Split(f, "\n"))
		h := v.transpose()

		switch part {
		case 1:
			vri, hri = v.reflection(0), h.reflection(0)
		case 2:
			vri, hri = v.reflection(1), h.reflection(1)
		default:
			return 0, errors.New("invalid part")
		}

		vsum += vri
		hsum += hri

		// fmt.Printf("vertical (ri: %d):\n%s\nhorizontal (ri: %d):\n%s\n\n", vri, v, hri, h)
	}

	return hsum + vsum*100, nil
}
