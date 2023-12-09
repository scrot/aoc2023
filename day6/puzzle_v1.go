package day6

import (
	"bufio"
	"bytes"
	"fmt"
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
			races = append(races, numbers(l, true))
		} else {
			races = append(races, numbers(l, false))
		}
	}

	for i := 0; i < len(races[0]); i++ {
		var (
			count        int
			time, record = races[0][i], races[1][i]
		)

		// hold * remaining time = distance
		for t := time / 2; t > 0; t-- {
			distance := t * (time - t)
			if distance > record {
				count += 2
			}
		}

		if time%2 == 0 {
			count--
		}

		fmt.Println(count)

		answer *= count
	}

	return answer, nil
}

func numbers(line string, ignoreWs bool) []int {
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

		if ignoreWs && r == ' ' {
			continue
		}

		if number > 0 {
			numbers = append(numbers, number)
			number = 0
		}
	}

	return numbers
}
