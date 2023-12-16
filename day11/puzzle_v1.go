package day11

import (
	"bufio"
	"bytes"
	"fmt"
	"slices"
)

type V1 struct{}

type galaxy struct {
	id  int
	loc [2]int
}

func (g galaxy) String() string {
	return fmt.Sprintf("%d (%d, %d)", g.id, g.loc[0], g.loc[1])
}

func NewCosmos(s *bufio.Scanner, part int) (gs []galaxy) {
	// find galaxy rows and columns
	var (
		gi, ci, ri   int
		gcols, grows []int
		gss          []galaxy
	)
	for ri = 0; s.Scan(); ri++ {
		l := s.Text()
		for ci = range l {
			r := l[ci]
			if r == '#' {
				gi++
				gss = append(gss, galaxy{gi, [2]int{ci, ri}})
				grows = append(grows, ri)
				gcols = append(gcols, ci)
			}
		}
	}

	// empty spaces
	var emptyRows []int
	for i := 0; i < ri; i++ {
		if !slices.Contains(grows, i) {
			emptyRows = append(emptyRows, i)
		}
	}

	var emptyCols []int
	for i := 0; i < ri; i++ {
		if !slices.Contains(gcols, i) {
			emptyCols = append(emptyCols, i)
		}
	}

	// calculate galaxy locations
	for _, g := range gss {
		var xexp int
		for _, gi := range emptyCols {
			if gi > g.loc[0] {
				break
			}
			if part == 1 {
				xexp += 1
			}
			if part == 2 {
				xexp += 1000000 - 1
			}
		}

		var yexp int
		for _, gi := range emptyRows {
			if gi > g.loc[1] {
				break
			}
			if part == 1 {
				yexp += 1
			}
			if part == 2 {
				yexp += 1000000 - 1
			}
		}

		gs = append(gs, galaxy{g.id, [2]int{g.loc[0] + xexp, g.loc[1] + yexp}})
	}

	return
}

func (V1) Solve(input []byte, part int) (int, error) {
	var (
		r  = bytes.NewReader(input)
		s  = bufio.NewScanner(r)
		gs = NewCosmos(s, part)
	)

	ds := distances(gs)

	var sum int
	for _, d := range ds {
		sum += d.distance
	}

	return sum, nil
}

type distance struct {
	from, to, distance int
}

func (d distance) String() string {
	return fmt.Sprintf("%d > %d: %d", d.from, d.to, d.distance)
}

func distances(gs []galaxy) (ds []distance) {
	for i := 0; i < len(gs); i++ {
		for j := i + 1; j < len(gs); j++ {
			dist := abs(gs[i].loc[0]-gs[j].loc[0]) + abs(gs[i].loc[1]-gs[j].loc[1])
			d := distance{gs[i].id, gs[j].id, dist}
			ds = append(ds, d)
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
