package day4

import (
	"bufio"
	"bytes"
	"errors"
	"strconv"
	"strings"
)

type V1 struct{}

func (V1) Solve(input []byte, part int) (int, error) {
	r := bytes.NewReader(input)
	s := bufio.NewScanner(r)

	var totalPoints, scratchcards int

	cardCount := make(map[int]int)

	var ci int
	for s.Scan() {
		ci++

		card := s.Text()[9:]
		win, cur, _ := strings.Cut(card, " | ")

		var points int
		m := in(win, cur)
		for range m {
			points += points
			if points == 0 {
				points++
			}
		}
		totalPoints += points

		// part 2
		cardCount[ci]++
		for j := range m {
			cardCount[ci+j+1] += cardCount[ci]
		}
	}

	if part == 1 {
		return totalPoints, nil
	}

	for _, count := range cardCount {
		scratchcards += count
	}

	if part == 2 {
		return scratchcards, nil
	}

	return 0, errors.New("invalid part")
}

func in(numbers, inNumbers string) (won []int) {
	in := make(map[int]bool)

	for _, n := range strings.Split(numbers, " ") {
		d, err := strconv.Atoi(n)
		if err == nil {
			in[d] = false
		}
	}

	for _, n := range strings.Split(inNumbers, " ") {
		d, err := strconv.Atoi(n)
		if err == nil {
			if _, ok := in[d]; ok {
				won = append(won, d)
			}
		}
	}
	return won
}
