package day2

import (
	"bufio"
	"bytes"
	"errors"
)

type V2 struct{}

func (_ V2) Solve(input []byte, part int) (int, error) {
	if part > 2 {
		return 0, errors.New("invalid part")
	}

	r := bytes.NewReader(input)
	s := bufio.NewScanner(r)

	var a1 int
	var a2 int

	for s.Scan() {
		l := s.Text()

		// skip 'Game '
		i := 5

		// game index
		var gi int
		for isDigit(rune(l[i])) {
			gi *= 10
			gi += int(l[i] - '0')
			i++
		}

		// skip ': '
		i += 2

		// max cube counts
		var ci, r, g, b int
		for i < len(l) {
			ru := rune(l[i])
			switch {
			case isDigit(ru):
				ci *= 10
				ci += int(ru - '0')
				i++
			case ru == 'r':
				if ci > r {
					r = ci
				}
				ci = 0
				i += 3
			case ru == 'g':
				if ci > g {
					g = ci
				}
				i += 5
				ci = 0
			case ru == 'b':
				if ci > b {
					b = ci
				}
				i += 4
				ci = 0
			default:
				i++
			}
		}

		if r <= 12 && g <= 13 && b <= 14 {
			a1 += gi
		}

		a2 += r * b * g
	}

	if part == 1 {
		return a1, nil
	}

	return a2, nil
}

func isDigit(r rune) bool {
	return r-'0' >= 0 && r-'0' <= 9
}
