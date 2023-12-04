package day1

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strconv"
)

type V2 struct{}

func (_ V2) Solve(input []byte, part int) (int, error) {
	if part > 2 {
		return 0, errors.New("invalid part")
	}

	r := bytes.NewReader(input)
	s := bufio.NewScanner(r)

	var answer int
	for s.Scan() {
		l := s.Text()

		var (
			ok          bool
			first, last int
		)

		for i := 0; i < len(l); i++ {
			first, ok = toDigit(rune(l[i]))
			if ok {
				break
			}
			if part == 2 {
				if n := spelled(l[i:]); n > 0 {
					first = n
					break
				}
			}
		}

		for i := len(l) - 1; i >= 0; i-- {
			last, ok = toDigit(rune(l[i]))
			if ok {
				break
			}
			if part == 2 {
				if n := spelled(l[i:]); n > 0 {
					last = n
					break
				}
			}
		}

		concat := fmt.Sprintf("%d%d", first, last)
		number, err := strconv.Atoi(concat)
		if err != nil {
			return 0, fmt.Errorf("concat NaN: %s", concat)
		}

		answer += number
	}

	return answer, nil
}
