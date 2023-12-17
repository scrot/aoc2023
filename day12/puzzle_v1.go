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

		clear(cache)
		sum += configurations(springs, groups, 0, 0, 0)

		fmt.Printf("%v %v > %d\n", springs, groups, sum)
		// fmt.Println(strings.Join(cs, "\n"))
		// fmt.Printf("seq: %s %v\n", springs, groups)
	}
	return sum, nil
}

var cache = make(map[string]int)

func configurations(ss string, gs []int, si, gi, count int) int {
	// fmt.Printf("s: %s g: %v si: %d gi: %d c: %d\n", ss, gs, si, gi, count)

	// check cache if configuration already visited
	key := fmt.Sprintf("%d %d %d", si, gi, count)
	if v, ok := cache[key]; ok {
		return v
	}

	// all groups matched
	if gi == len(gs) {
		// valid if no broken springs left
		if strings.Contains(ss[si:], "#") {
			return 0
		} else {
			return 1
		}
	}

	// no springs left
	if si == len(ss) {
		// process last count
		if count == gs[gi] {
			gi++
		}

		// valid if all groups matched
		if gi == len(gs) {
			return 1
		} else {
			return 0
		}

	}

	var matched int
	switch ss[si] {
	case '.':
		if count > 0 {
			if count == gs[gi] {
				matched = configurations(ss, gs, si+1, gi+1, 0)
			} else {
				matched = 0
			}
		} else {
			matched = configurations(ss, gs, si+1, gi, count)
		}
	case '#':
		matched = configurations(ss, gs, si+1, gi, count+1)
	case '?':
		// handle .
		if count > 0 {
			if count == gs[gi] {
				matched += configurations(ss, gs, si+1, gi+1, 0)
			} else {
				matched = 0
			}
		} else {
			matched += configurations(ss, gs, si+1, gi, count)
		}

		// handle #
		matched += configurations(ss, gs, si+1, gi, count+1)
	}

	// remember result for spring / group combination
	cache[key] = matched

	return matched
}
