package day1

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type V1 struct{}

func (V1) Solve(input []byte, part int) (int, error) {
	if part > 2 {
		return 0, errors.New("invalid part")
	}

	r := bytes.NewReader(input)
	s := bufio.NewScanner(r)

	var s1 int
	for s.Scan() {
		var ns []int
		l := s.Text()
		for i, r := range l {
			n := (r - '0')
			if n > 0 && n <= 9 {
				ns = append(ns, int(n))
			}

			if part == 2 {
				sn := spelled(l[i:])
				if sn > 0 {
					ns = append(ns, sn)
				}
			}
		}

		concat := fmt.Sprintf("%d%d", ns[0], ns[len(ns)-1])
		ccn, err := strconv.Atoi(concat)
		if err != nil {
			return 0, fmt.Errorf("concat NaN: %s", concat)
		}

		s1 += ccn
	}

	return s1, nil
}

func toDigit(r rune) (int, bool) {
	number := r - '0'
	if number > 0 && number <= 9 {
		return int(number), true
	}
	return 0, false
}

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func spelled(l string) int {
	for k, v := range digits {
		if len(k) <= len(l) {
			if k == l[:len(k)] {
				return v
			}
		}
	}
	return 0
}
