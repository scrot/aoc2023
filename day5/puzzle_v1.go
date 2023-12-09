package day5

import (
	"bufio"
	"bytes"
	"sort"
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

type mapping struct {
	rs []rule
}

func (m mapping) Len() int {
	return len(m.rs)
}

func (m mapping) Less(i, j int) bool {
	return m.rs[i].start < m.rs[j].start
}

func (m mapping) Swap(i, j int) {
	m.rs[i], m.rs[j] = m.rs[j], m.rs[i]
}

func (V1) Solve(input []byte, part int) (int, error) {
	r := bytes.NewReader(input)
	s := bufio.NewScanner(r)

	var (
		seeds   [][2]int
		almanac []mapping
	)

	if s.Scan() {
		ns := numbers(s.Text())

		if part == 1 {
			for _, n := range ns {
				seeds = append(seeds, [2]int{n, n})
			}
		}

		if part == 2 {
			for i := 0; i < len(ns)-1; i += 2 {
				seeds = append(seeds, [2]int{ns[i], ns[i] + ns[i+1] - 1})
			}
		}
	}

	var m mapping
	for s.Scan() {
		ns := numbers(s.Text())

		if len(ns) > 0 {
			m.rs = append(m.rs, newRule(ns[1], ns[0], ns[2]))
		} else {
			if m.Len() > 0 {
				sort.Sort(m)
				almanac = append(almanac, m)
				m = mapping{}
			}
		}
	}

	if m.Len() > 0 {
		sort.Sort(m)
		almanac = append(almanac, m)
		m = mapping{}
	}

	locs := seeds
	for _, m := range almanac {
		ds := locs
		locs = [][2]int{}
		for _, d := range ds {
			locs = append(locs, destination(d, m.rs)...)
		}
	}

	closest := -1
	for _, l := range locs {
		if l[0] < closest || closest == -1 {
			closest = l[0]
		}
	}

	return closest, nil
}

// destination recusively checks a seed against all
// mapping rules returning the destination ranges
func destination(src [2]int, rs []rule) [][2]int {
	var (
		ds        [][2]int
		remainder [2]int
	)

	// tail condition
	if src == [2]int{} {
		return [][2]int{}
	}

	if len(rs) == 0 {
		return [][2]int{src}
	}

	// split for current rule
	xs, remainder := split(src, rs[0])
	ds = append(ds, xs...)

	// remainder recursive
	ds = append(ds, destination(remainder, rs[1:])...)

	// fmt.Printf("ds: %v\n", ds)

	return ds
}

func split(src [2]int, r rule) ([][2]int, [2]int) {
	switch {
	case src[1] < r.start:
		return [][2]int{src}, [2]int{}
	case src[0] > r.end:
		return [][2]int{}, src
	case r.start <= src[0] && src[1] <= r.end:
		return [][2]int{{src[0] + r.offset, src[1] + r.offset}}, [2]int{}
	case src[0] < r.start && src[1] <= r.end:
		lhs := [2]int{src[0], r.start - 1}
		rhs := [2]int{r.start + r.offset, src[0] + r.offset}
		return [][2]int{lhs, rhs}, [2]int{}
	case src[0] >= r.start && src[1] > r.start:
		lhs := [2]int{src[0] + r.offset, r.end + r.offset}
		rhs := [2]int{r.end + 1, src[1]}
		return [][2]int{lhs}, rhs
	case src[0] < r.start && src[1] > r.end:
		lhs := [2]int{src[0], r.start - 1}
		mid := [2]int{r.start + r.offset, r.end + r.offset}
		rhs := [2]int{r.end + 1, src[1]}
		return [][2]int{lhs, mid}, rhs
	default:
		return [][2]int{}, [2]int{}
	}
}

func numbers(line string) []int {
	var (
		isNumber bool
		number   int
		numbers  []int
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
