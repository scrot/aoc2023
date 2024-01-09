package day16

import (
	"bufio"
	"bytes"
)

type V2 struct{}

func (V2) Solve(input []byte, part int) (int, error) {
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
		var (
			mx int
			bs = g.genBeams()
			ch = make(chan int, len(bs))
		)

		for _, b := range bs {
			go pEnergize(b, g, ch)
		}

		for i := 0; i < len(bs); i++ {
			mx = max(mx, <-ch)
		}

		return mx, nil
	}

	return 0, nil
}

func pEnergize(b beam, g grid, ch chan<- int) {
	ch <- energize(b, g)
}
