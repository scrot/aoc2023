package day17

import (
	"bufio"
	"bytes"
	"fmt"
)

type V1 struct{}

func (V1) Solve(input []byte, part int) (int, error) {
	bt := newBlockTree(input)
	fmt.Println(bt)
	return 0, nil
}

type block struct {
	pos  loc
	loss int
	bs   []block
}

func newBlockTree(input []byte) (bt block) {
	s := bufio.NewScanner(bytes.NewReader(input))

	var m []string
	for s.Scan() {
		m = append(m, s.Text())
	}

	var (
		bs     []block
		last   = -1
		dirs   = []loc{{0, -1}, {0, 1}, {1, 0}}
		xm, ym = len(m[0]), len(m)
	)

	for i, dir := range dirs {
		if i == last {
			continue
		}

		for count := 1; count <= 3; count++ {
			pos := loc{bt.pos.x + dir.x*count, bt.pos.y + dir.y*count}
			if pos.x < 0 || pos.x >= xm || pos.y < 0 || pos.y >= ym {
				continue
			}
			loss := m[pos.y][pos.x]
			bs = append(bs, block{pos, int(loss - '0'), []block{}})
		}
	}

	bt.bs = bs

	return
}

type loc struct {
	x, y int
}
