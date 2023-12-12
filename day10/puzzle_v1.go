package day10

import (
	"bufio"
	"bytes"
	"fmt"
)

type V1 struct{}

type maze [][]rune

func (m maze) String() string {
	var s string
	for _, row := range m {
		s += string(row) + "\n"
	}
	return s
}

func (V1) Solve(input []byte, part int) (int, error) {
	var (
		r         = bytes.NewReader(input)
		s         = bufio.NewScanner(r)
		mz, start = newMaze(s)
		count     = 1
	)

	fmt.Println(mz)

	ds := start.directions(mz)
	cw, ccw := ds[0], ds[1]

	for cw.String() != ccw.String() {
		fmt.Printf("cw: %s, prev: %s\n", cw, cw.prev)
		cw = cw.directions(mz)[0]
		fmt.Printf("ccw: %s, prev: %s\n", ccw, ccw.prev)
		ccw = ccw.directions(mz)[0]
		count++
	}

	return count, nil
}

type loc struct {
	x, y   int
	prev   *loc
	symbol rune
}

func (l loc) String() string {
	return fmt.Sprintf("%c (%d, %d)", l.symbol, l.x, l.y)
}

// directions returns all locations to go to
// in the maze for a given location
func (l loc) directions(m maze) []loc {
	var ds []loc

	for dir := 0; dir < 4; dir++ {
		var nl loc
		nl.x, nl.y = l.direction(dir)

		if nl.x >= 0 && nl.x < len(m) &&
			nl.y >= 0 && nl.y < len(m[0]) {

			nl.symbol = m[nl.y][nl.x]
			nl.prev = &l

			fmt.Printf("cur: %s, dir: %d, new: %s\n", l, dir, nl)

			if l.prev.String() != nl.String() &&
				fits(l.symbol, nl.symbol, dir) {
				fmt.Printf("%s != %s\n", l.prev, nl)
				ds = append(ds, nl)
			}
		}
	}
	fmt.Printf("valid dirs: %s\n\n", ds)
	return ds
}

// toXY provides a new location given a
// direction dir (0=left, 1=up, 2=right, 3=down)
func (l loc) direction(dir int) (int, int) {
	ls := map[int][2]int{
		0: {-1, 0}, 1: {0, -1},
		2: {1, 0}, 3: {0, 1},
	}

	return l.x + ls[dir][0], l.y + ls[dir][1]
}

// fitting maps pipe to fitting pipes for
// direction left, up, right, down
var fitting = map[rune][4][]rune{
	'│': {{}, {'┌', '┐', '│'}, {}, {'┘', '└', '│'}},
	'─': {{'┌', '└', '─'}, {}, {'┐', '┘', '─'}, {}},
	'┌': {{}, {}, {'┐', '┘', '─'}, {'┘', '└', '│'}},
	'└': {{}, {'┌', '┐', '│'}, {'┐', '┘', '─'}, {}},
	'┐': {{'┌', '└', '─'}, {}, {}, {'┘', '└', '│'}},
	'┘': {{'┌', '└', '─'}, {'┌', '┐', '│'}, {}, {}},
	' ': {{}, {}, {}, {}},

	'🐿': {
		{'┌', '└', '─'},
		{'┌', '┐', '│'},
		{'┐', '┘', '─'},
		{'┘', '└', '│'},
	},
}

// connected checks if q fits onto p from
// direction dir (0=left, 1=up, 2=right, 3=down)
func fits(p, q rune, dir int) bool {
	// fmt.Printf("\tc: %c fits nl: %c\n", p, q)
	for _, f := range fitting[p][dir] {
		if q == f {
			return true
		}
	}
	return false
}

// pretty maps ugly symbols to pretty ones
var pretty = map[rune]rune{
	'|': '│',
	'-': '─',
	'F': '┌',
	'L': '└',
	'7': '┐',
	'J': '┘',
	'.': ' ',
	'S': '🐿',
}

// build new pretty maze
func newMaze(s *bufio.Scanner) (maze, loc) {
	var (
		start loc
		maze  [][]rune
	)

	for y := 0; s.Scan(); y++ {
		var row []rune
		for x, r := range s.Text() {
			if r == 'S' {
				start = loc{x, y, &loc{}, pretty[r]}
			}
			row = append(row, pretty[r])
		}
		maze = append(maze, row)
	}
	return maze, start
}
