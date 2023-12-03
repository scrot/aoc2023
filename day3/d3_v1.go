package day3

import "fmt"

func SolveV1(input []byte, part int) (int, error) {
	var vlen int
	for _, b := range input {
		if rune(b) == '\n' {
			break
		}
		vlen++
	}

	if part == 1 {
		return partNumbers(input, vlen), nil
	}

	return gearRatio(input, vlen), nil
}

func gearRatio(input []byte, vlen int) int {
	var ans int
	for i, b := range input {
		if rune(b) == '*' {
			ns := numbers(input, i, vlen)
			fmt.Println(ns)
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
	fmt.Println(ss)
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
		-vlen - 2, -vlen - 1, -vlen, -1,
		1, vlen, vlen + 1, vlen + 2,
	}

	surround := make(map[int]byte)
	for _, i := range is {
		if index+i >= 0 && index+i < len(input) {
			surround[index+i] = input[index+i]
		}
	}

	return surround
}

func isDigit(b byte) bool {
	r := rune(b)
	if r-'0' >= 0 && r-'0' <= 9 {
		return true
	}
	return false
}
