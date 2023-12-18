package day13

import (
	"fmt"
	"slices"
	"strings"
)

type grid [][]rune

func (g grid) String() (s string) {
	for _, r := range g {
		s += string(r) + "\n"
	}
	return
}

func (g grid) Reflection() int {
	// find pair
	var ris []int
	for i := 0; i < len(g)-1; i++ {
		if slices.Compare(g[i], g[i+1]) == 0 {
			ris = append(ris, i)
		}
	}

	// no matching pair found
	if len(ris) == 0 {
		return 0
	}

	// complete reflection
	for _, ri := range ris {
		reflection := true
		up, down := ri+1, ri
		for down >= 0 && up < len(g) {
			if slices.Compare(g[down], g[up]) != 0 {
				reflection = false
				break
			}
			down--
			up++
		}
		if reflection {
			return ri + 1
		}
	}

	return 0
}

func (g grid) Pivot() (p grid) {
	for x := 0; x < len(g[0]); x++ {
		var row []rune
		for y := 0; y < len(g); y++ {
			row = append(row, g[y][x])
		}
		p = append(p, row)
	}

	return
}

func newGrid(ls []string) (g grid) {
	for _, l := range ls {
		var row []rune
		for _, r := range l {
			row = append(row, r)
		}
		if len(row) > 0 {
			g = append(g, row)
		}
	}
	return
}

type V1 struct{}

func (V1) Solve(input []byte, part int) (int, error) {
	var (
		fields     = strings.Split(string(input), "\n\n")
		vsum, hsum int
	)

	for _, f := range fields {
		v := newGrid(strings.Split(f, "\n"))
		h := v.Pivot()

		vri, hri := v.Reflection(), h.Reflection()
		vsum += vri
		hsum += hri

		fmt.Printf("vertical (ri: %d):\n%s\nhorizontal (ri: %d):\n%s\n\n", vri, v, hri, h)
	}

	return hsum + vsum*100, nil
}
