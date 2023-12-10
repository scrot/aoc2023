package day7

import (
	"bufio"
	"bytes"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type V1 struct{}

type hand struct {
	cards [5]int
	bid   int
}

type hands struct {
	part int
	arr  []hand
}

// addHand takes a input line and converts
// it into a hand with cards and a bid
func (hs *hands) addHand(line string) {
	l := strings.Split(line, " ")

	var h hand
	for i, r := range l[0] {
		h.cards[i] = toValue(r, hs.part)
	}

	h.bid, _ = strconv.Atoi(l[1])

	hs.arr = append(hs.arr, h)
}

// implementation of sort.Interface
func (hs *hands) Len() int {
	return len(hs.arr)
}

// implementation of sort.Interface
func (hs *hands) Less(i, j int) bool {
	// hi, hj := highestCard(hs.arr[i]), highestCard(hs.arr[j])
	si, sj := strength(hs.arr[i], hs.part), strength(hs.arr[j], hs.part)

	switch {
	case si < sj:
		return true
	// case si == 0 && sj == 0 && hi != hj:
	// 	return hi < hj
	case si == sj:
		for k := 0; k < 5; k++ {
			ci, cj := hs.arr[i].cards[k], hs.arr[j].cards[k]
			if ci != cj {
				return ci < cj
			}
		}
		return true
	default:
		return false
	}
}

// implementation of sort.Interface
func (hs *hands) Swap(i, j int) {
	hs.arr[i], hs.arr[j] = hs.arr[j], hs.arr[i]
}

func (V1) Solve(input []byte, part int) (int, error) {
	r := bytes.NewReader(input)
	s := bufio.NewScanner(r)

	hs := hands{part: part}
	for s.Scan() {
		hs.addHand(s.Text())
	}

	sort.Sort(&hs)

	var winnings int
	for i, h := range hs.arr {
		winnings += (i + 1) * h.bid
	}

	return winnings, nil
}

// strenght returns a ranking of card combinations
// with the highest is a 5-of-a-kind and weakest a high-card
func strength(h hand, part int) int {
	combo := make(map[int]int)
	for _, c := range h.cards {
		combo[c]++
	}

	var counts []int
	for k, v := range combo {
		if part == 1 {
			counts = append(counts, v)
		}
		if part == 2 && k != 1 {
			counts = append(counts, v)
		}
	}
	slices.Sort(counts)

	if part == 2 {
		if len(counts) == 0 {
			return 6
		}
		counts[len(counts)-1] += combo[1]
	}

	switch hi := counts[len(counts)-1]; hi {
	case 5:
		return 6
	case 4:
		return 5
	case 3:
		if counts[len(counts)-2] == 2 {
			return 4
		}
		return 3
	case 2:
		if counts[len(counts)-2] == 2 {
			return 2
		}
		return 1
	default:
		return 0
	}
}

// toValue turns card symbol into integer
func toValue(card rune, part int) int {
	switch card {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		if part == 2 {
			return 1
		}
		return 11
	case 'T':
		return 10
	default:
		return int(card - '0')
	}
}
