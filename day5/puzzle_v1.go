package day5

import (
	"bufio"
	"bytes"
	"cmp"
	"slices"
)

type V1 struct{}

type rule struct {
	start  int
	end    int
	offset int
}

func newRule(src int, dst int, length int) rule {
	return rule{src, src + length - 1, dst - src}
}

type mapping = []rule

func (V1) Solve(input []byte, part int) (int, error) {
	r := bytes.NewReader(input)
	s := bufio.NewScanner(r)

	// first line to seed ranges
	s.Scan()
	source := seeds(s.Text(), part)

	// for each map translate sources to destinations
	var m mapping
	for i := 0; i < 7; {
		s.Scan()
		ns := numbers(s.Text())

		if len(ns) > 0 {
			m = append(m, newRule(ns[1], ns[0], ns[2]))
			continue
		}

		if len(m) > 0 {
			slices.SortFunc(m, func(a, b rule) int {
				return cmp.Compare(a.start, b.start)
			})

			var destination [][2]int
			for _, s := range source {
				destination = append(destination, destinations(s, m)...)
			}
			source = destination

			m = mapping{}
			i++
		}
	}

	// find lowest location
	closest := slices.MinFunc(source, func(a, b [2]int) int {
		return cmp.Compare(a[0], b[0])
	})[0]

	return closest, nil
}

// seeds parses seeds numbers
// for part 1 each number is a seed range of 1
// for part 2 each pair is transformed to start-end
func seeds(line string, part int) [][2]int {
	var (
		seeds [][2]int
		ns    = numbers(line)
	)

	switch part {
	case 1:
		for _, n := range ns {
			seeds = append(seeds, [2]int{n, n})
		}
	case 2:
		for i := 0; i < len(ns)-1; i += 2 {
			seeds = append(seeds,
				[2]int{ns[i], ns[i] + ns[i+1] - 1},
			)
		}
	}

	return seeds
}

// destinations recusively checks a seed against all
// mapping rules returning the destinations ranges
func destinations(src [2]int, rs []rule) [][2]int {
	var (
		ds    [][2]int
		right [2]int
	)

	// tail: no source range
	if src == [2]int{} {
		return [][2]int{}
	}

	// tail: no rules left
	if len(rs) == 0 {
		return [][2]int{src}
	}

	// split for current rule
	left, mid, right := splitv2(src, rs[0])
	if left != [2]int{} {
		ds = append(ds, left)
	}

	if mid != [2]int{} {
		ds = append(ds, mid)
	}

	// recusively append destinations of remaining rules
	ds = append(ds, destinations(right, rs[1:])...)

	return ds
}

func splitv2(src [2]int, r rule) ([2]int, [2]int, [2]int) {
	if r.start <= src[0] && src[1] <= r.end {
		mid := [2]int{src[0] + r.offset, src[1] + r.offset}
		return [2]int{}, mid, [2]int{}
	}

	left, rest := cut(src, r.start-1)
	mid, right := cut(rest, r.end)

	if mid != [2]int{} {
		mid[0] += r.offset
		mid[1] += r.offset
	}

	return left, mid, right
}

func cut(src [2]int, i int) ([2]int, [2]int) {
	if i < src[0] {
		return [2]int{}, src
	}
	if i >= src[1] {
		return src, [2]int{}
	}
	return [2]int{src[0], i}, [2]int{i + 1, src[1]}
}

func numbers(line string) []int {
	var (
		number   int
		numbers  []int
		isNumber bool
	)

	for _, r := range line {
		if '0' <= r && r <= '9' {
			d := int(r - '0')
			number *= 10
			number += d
			isNumber = true
		} else {
			if isNumber {
				numbers = append(numbers, number)
				number = 0
			}
			isNumber = false
		}
	}

	if isNumber {
		numbers = append(numbers, number)
	}

	return numbers
}
