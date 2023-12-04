package day3

import "errors"

type V1 struct{}

func (V1) Solve(input []byte, part int) (int, error) {
	var width int
	for _, b := range input {
		if rune(b) == '\n' {
			break
		}
		width++
	}

	switch part {
	case 1:
		return partNumbers(input, width), nil
	case 2:
		return gearRatio(input, width), nil
	default:
		return 0, errors.New("invalid part")
	}
}

func gearRatio(input []byte, vlen int) int {
	var ans int
	for i, b := range input {
		if rune(b) == '*' {
			ns := numbers(input, i, vlen)
			if len(ns) == 2 {
				ans += ns[0] * ns[1]
			}
		}
	}
	return ans
}

func partNumbers(input []byte, vlen int) int {
	var (
		n, ans int
		adj    bool
	)

	for i, b := range input {
		r := rune(b)
		switch {
		case r-'0' >= 0 && r-'0' <= 9:
			n *= 10
			n += int(r - '0')
			if adjecent(input, i, vlen) {
				adj = true
			}
		default:
			if adj {
				ans += n
			}
			n = 0
			adj = false
		}
		// fmt.Printf("r:%c n:%d adj:%t ans:%d\n", r, n, adj, ans)
	}

	// handle number with adj at end of array
	if adj {
		ans += n
	}

	return ans
}

func numbers(input []byte, index, vlen int) []int {
	var is []int
	ss := surround(input, index, vlen)
	for i := range ss {
		// skip if surrounding not number
		if !isDigit(input[i]) {
			continue
		}

		// find first digit of number
		var number int
		for i > 0 && isDigit(input[i]) {
			i--
		}
		if i > 0 {
			i++
		}

		// parse number and remove redudant
		for i < len(input) && isDigit(input[i]) {
			number *= 10
			d := rune(input[i]) - '0'
			number += int(d)
			delete(ss, i)
			i++
		}

		if number > 0 {
			is = append(is, number)
		}
	}
	return is
}

func adjecent(input []byte, index, vlen int) bool {
	ss := surround(input, index, vlen)
	for _, s := range ss {
		if !isDigit(s) && s != '.' && s != '\n' {
			return true
		}
	}
	return false
}

func surround(input []byte, index, vlen int) map[int]byte {
	is := []int{
		-vlen - 2, -vlen - 1, -vlen,
		-1, 1,
		vlen, vlen + 1, vlen + 2,
	}

	surround := make(map[int]byte)
	for _, i := range is {
		x := index + i
		if x >= 0 && x < len(input) {
			surround[x] = input[x]
		}
	}

	return surround
}

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}
