package day10

import (
	"bufio"
	"bytes"
	"fmt"
)

type V1 struct{}

// Maze represents the pipesystem
type maze [][]rune

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
				start = loc{[2]int{x, y}, &loc{}, pretty[r]}
			}
			row = append(row, pretty[r])
		}
		maze = append(maze, row)
	}
	return maze, start
}

func (m maze) String() (s string) {
	for _, row := range m {
		s += string(row) + "\n"
	}
	return
}

func (m maze) filter(end loc) maze {
	for y, row := range m {
		for x := range row {
			m[y][x] = ' '
		}
	}

	for end.coord != (loc{}).coord {
		m[end.coord[1]][end.coord[0]] = end.symbol
		end = *end.prev
	}

	return m
}

// loc represents a location in a maze
type loc struct {
	coord  [2]int
	prev   *loc
	symbol rune
}

func (l loc) String() string {
	return fmt.Sprintf("%c (%d, %d)", l.symbol, l.coord[0], l.coord[1])
}

// directions returns all locations to go to
// in the maze for a given location
func (l loc) directions(m maze) []loc {
	var ds []loc

	for dir := 0; dir < 4; dir++ {
		var nl loc
		nl.coord = l.direction(dir)

		if nl.coord[1] >= 0 && nl.coord[1] < len(m) &&
			nl.coord[0] >= 0 && nl.coord[0] < len(m[0]) &&
			l.prev.coord != nl.coord {

			nl.symbol = m[nl.coord[1]][nl.coord[0]]
			nl.prev = &l

			// fmt.Printf("cur: %s, dir: %d, new: %s\n", l, dir, nl)

			if fits(l.symbol, nl.symbol, dir) {
				// fmt.Printf("%s != %s\n", l.prev, nl)
				ds = append(ds, nl)
			}
		}
	}
	// fmt.Printf("valid dirs: %s\n\n", ds)
	return ds
}

// toXY provides a new location given a
// direction dir (0=left, 1=up, 2=right, 3=down)
func (l loc) direction(dir int) [2]int {
	ls := map[int][2]int{
		0: {-1, 0}, 1: {0, -1},
		2: {1, 0}, 3: {0, 1},
	}

	return [2]int{l.coord[0] + ls[dir][0], l.coord[1] + ls[dir][1]}
}

func (V1) Solve(input []byte, part int) (int, error) {
	var (
		r         = bytes.NewReader(input)
		s         = bufio.NewScanner(r)
		mz, start = newMaze(s)
	)
	// fmt.Println(mz)

	var (
		count = 0
		end   = loc{start.coord, start.prev, '┌'}
	)

	for len(end.directions(mz)) > 0 {
		// fmt.Printf("cw: %s, prev: %s\n", end, end.prev)
		end = end.directions(mz)[0]
		count++
	}

	if part == 2 {
		var (
			mz    = mz.filter(end)
			count int
			in    bool
		)

		for y, row := range mz {
			for x, col := range row {
				in = updateIn(col, in)
				if col == ' ' && in {
					mz[y][x] = '░'
					count++
				}
			}
			// fmt.Println(mz)
		}

		// fmt.Println(mz)

		return count, nil
	}

	return count/2 + 1, nil
}

func updateIn(r rune, in bool) bool {
	switch r {
	// case '│', '┌', '└', '┐', '┘':
	// 	return !in
	case '│', '┌', '┐':
		return !in
	default:
		return in
	}
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

// connectedchecks if q fits onto p from
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
