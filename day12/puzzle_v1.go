package day12

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type V1 struct{}

func (V1) Solve(input []byte, part int) (int, error) {
	r := bytes.NewReader(input)
	s := bufio.NewScanner(r)

	var sum int
	for s.Scan() {
		springs, after, _ := strings.Cut(s.Text(), " ")

		if part == 2 {
			springs = strings.Repeat(springs+"?", 5)
			springs = springs[:len(springs)-1]
			after = strings.Repeat(","+after, 5)
			after = after[1:]

		}
		var groups []int
		for _, seg := range strings.Split(after, ",") {
			n, _ := strconv.Atoi(seg)
			groups = append(groups, n)
		}

		ps := permuations(springs)
		cs := configurations(ps, groups)

		sum += len(cs)

		fmt.Printf("%v %v > %d\n", springs, groups, len(cs))
		// fmt.Println(strings.Join(cs, "\n"))
		// fmt.Printf("seq: %s %v\n", springs, groups)

	}
	return sum, nil
}

var cache = make(map[string]string)

func permuations(springs string) (p []string) {
	i := strings.Index(springs, "?")

	// tail
	if i == -1 {
		p = append(p, springs)
		return
	}

	operational := springs[:i] + "." + springs[i+1:]
	p = append(p, permuations(operational)...)

	broken := springs[:i] + "#" + springs[i+1:]
	p = append(p, permuations(broken)...)

	return
}

func configurations(permutations []string, groups []int) (c []string) {
	for _, p := range permutations {
		// fmt.Printf("p: %s gs: %v\n", p, groups)
		if matches(p, groups) {
			// fmt.Printf("match: %t\n\n\n", true)
			c = append(c, p)
		}
	}
	return
}

func matches(p string, gs []int) bool {
	var matched int
	var count int
	for _, r := range p {
		// fmt.Printf("i: %d last: %d count: %d matched: %d\n", i, len(p)-1, count, matched)

		switch r {
		case '#':
			if matched == len(gs) {
				return false
			}
			count++
		case '.':
			if count == 0 {
				continue
			}

			if matched >= len(gs) {
				return false
			}

			if count != gs[matched] {
				return false
			}

			matched++

			count = 0
		}
	}

	if matched < len(gs) && count == gs[matched] {
		matched++
	}

	return matched == len(gs)
}
