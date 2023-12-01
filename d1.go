package aoc2023

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func day1(input string, part int) (int, error) {
	r := strings.NewReader(input)
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
