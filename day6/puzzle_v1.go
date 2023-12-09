package day6

import (
	"bufio"
	"bytes"
	"strings"
)

type V1 struct{}

func (V1) Solve(input []byte, part int) (int, error) {
	r := bytes.NewReader(input)
	s := bufio.NewScanner(r)

	answer := 1
	var races [][]int

	for i := 0; s.Scan(); i++ {
		l := s.Text()
		if part == 2 {
			l = strings.ReplaceAll(l, " ", "")
		}
		races = append(races, numbers(l))
	}

	for i := 0; i < len(races[0]); i++ {
		var (
			count        int
			time, record = races[0][i], races[1][i]
		)

		// hold * remaining time = distance
		for t := 1; t < time; t++ {
			distance := t * (time - t)
			if distance > record {
				count++
			}
		}

		answer *= count
	}

	return answer, nil
}

func numbers(line string) []int {
	var (
		number  int
		numbers []int
	)

	for _, r := range line + "\n" {
		if '0' <= r && r <= '9' {
			d := int(r - '0')
			number *= 10
			number += d
			continue
		}

		if number > 0 {
			numbers = append(numbers, number)
			number = 0
		}
	}

	return numbers
}
