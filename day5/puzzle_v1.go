package day5

import (
	"bufio"
	"bytes"
	"fmt"
)

type V1 struct{}

type Mapping struct {
	Start  int
	End    int
	Offset int
}

func (V1) Solve(input []byte, part int) (int, error) {
	r := bytes.NewReader(input)
	s := bufio.NewScanner(r)

	var (
		first = true
		ms    []Mapping
		seeds []Mapping
	)

	for s.Scan() {
		line := s.Text()

		// numbers in line
		ns := numbers(line)

		// first seeds line
		if first {
			for _, n := range ns {
				if part == 1 {
					seeds = append(seeds, Mapping{n, n, 0})
				}
			}
			first = false
			continue
		}

		// skip all non number rows
		if len(ns) == 0 {
			// update source numbers with desination
			if len(ms) > 0 {
				var res []Mapping
				for _, s := range seeds {
					res = append(res, Destination(s, ms))
				}
				seeds = res

				ms = []Mapping{}
			}
			continue
		}

		ms = append(ms, Convert(ns))
	}

	// process last map
	if len(ms) > 0 {
		var res []Mapping
		for _, s := range seeds {
			res = append(res, Destination(s, ms))
		}
		seeds = res
	}

	fmt.Println(seeds)

	return Closest(seeds), nil
}

// closest returns the lowest of a range of mappings
func Closest(locations []Mapping) int {
	closest := -1
	for _, l := range locations {
		if l.Start+l.Offset < closest || closest == -1 {
			closest = l.Start + l.Offset
		}
	}
	return closest
}

// Destination updates offset of s if it matches a
// range in ms, otherwise it returns s
func Destination(s Mapping, ms []Mapping) Mapping {
	for _, m := range ms {
		o := Overlapping(s, m)
		empty := Mapping{}
		if o != empty {
			return o
		}
	}
	return s
}

// Overlapping calculates section where a range in m
// overlaps with s or returns empty map if no overlap
func Overlapping(s, m Mapping) Mapping {
	var (
		res        = Mapping{m.Start, m.End, s.Offset + m.Offset}
		start, end = s.Start + s.Offset, s.End + s.Offset
	)

	if start <= m.End && end >= m.Start {
		if start >= m.Start {
			res.Start = s.Start
		}
		if end <= m.End {
			res.End = s.Start
		}
		return res
	}

	return Mapping{}
}

// Convert translate range notation to mapping
func Convert(numbers []int) Mapping {
	return Mapping{
		numbers[1],
		numbers[1] + numbers[2] - 1,
		numbers[0] - numbers[1],
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
